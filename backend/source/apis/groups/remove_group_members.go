package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func RemoveGroupMembersHandler(db *gorm.DB) groups_v1.RemoveGroupUsersV1HandlerFunc {
	return func(params groups_v1.RemoveGroupUsersV1Params) middleware.Responder {
		removeGroupUsersResp, errResp := processRemoveGroupUsersRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewRemoveGroupUsersV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewRemoveGroupUsersV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewRemoveGroupUsersV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewRemoveGroupUsersV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewRemoveGroupUsersV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewRemoveGroupUsersV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewRemoveGroupUsersV1OK()
		resp.SetPayload(removeGroupUsersResp)
		return resp
	}
}
func processRemoveGroupUsersRequest(db *gorm.DB, params groups_v1.RemoveGroupUsersV1Params) (resp *models.GroupResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	_, err := checkIfGroupOwner(db, payload.Username, params.GroupID)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusForbidden, err.Error())
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	users := params.Body.Users
	groupid := params.GroupID

	//delete
	tx := db.Begin()
	for _, user := range users {
		if user.Username == payload.Username {
			continue
		}
		//err = tx.Delete(&dbpackages.GroupUser{GroupID: groupid, Username: user.Username}).Error
		err = tx.Table(dbpackages.GroupUserTable).Where("group_id = ? AND username = ?", groupid, user.Username).Delete(&dbpackages.GroupUser{}).Error
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			tx.Rollback()
			return
		}
	}
	if err = tx.Commit().Error; err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}
	resp = &models.GroupResponse{
		GroupID: params.GroupID,
	}

	return
}

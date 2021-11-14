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

func AddGroupMembersHandler(db *gorm.DB) groups_v1.AddGroupUsersV1HandlerFunc {
	return func(params groups_v1.AddGroupUsersV1Params) middleware.Responder {
		addGroupUsersResp, errResp := processAddGroupUsersRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewAddGroupUsersV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewAddGroupUsersV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewAddGroupUsersV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewAddGroupUsersV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewAddGroupUsersV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewAddGroupUsersV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewAddGroupUsersV1OK()
		resp.SetPayload(addGroupUsersResp)
		return resp
	}
}
func processAddGroupUsersRequest(db *gorm.DB, params groups_v1.AddGroupUsersV1Params) (resp *models.GroupResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	_, err := CheckIfGroupUser(db, payload.Username, params.GroupID)
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	users := params.Body.Users
	groupid := params.GroupID

	//insert
	tx := db.Begin()
	for _, user := range users {
		err = tx.Save(&dbpackages.GroupUser{GroupID: groupid, Username: user.Username}).Error
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

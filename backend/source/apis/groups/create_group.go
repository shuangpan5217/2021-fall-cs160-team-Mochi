package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func CreateGroupV1Handler(db *gorm.DB) groups_v1.CreateGroupV1HandlerFunc {
	return func(params groups_v1.CreateGroupV1Params) middleware.Responder {
		newGroupResp, errResp := processCreateGroupRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewCreateGroupV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewCreateGroupV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewCreateGroupV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewCreateGroupV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewCreateGroupV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewCreateGroupV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewCreateGroupV1OK()
		resp.SetPayload(newGroupResp)
		return resp
	}
}

func processCreateGroupRequest(db *gorm.DB, params groups_v1.CreateGroupV1Params) (resp *models.GroupResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username
	body := params.Body
	description := body.Description
	groupname := body.GroupName
	groupID := uuid.NewString()

	var group = dbpackages.Group{
		GroupID:     groupID,
		GroupName:   *groupname,
		Description: description,
		GroupOwner:  payload.Username,
	}
	tx := db.Begin()
	err := tx.Save(&group).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}
	err = tx.Save(&dbpackages.GroupUser{GroupID: groupID, Username: username}).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.GroupResponse{
		GroupID: groupID,
	}

	return
}

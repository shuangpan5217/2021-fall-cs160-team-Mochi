package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func UpdateGroupInfoV1Handler(db *gorm.DB) groups_v1.UpdateGroupInfoV1HandlerFunc {
	return groups_v1.UpdateGroupInfoV1HandlerFunc(func(params groups_v1.UpdateGroupInfoV1Params) middleware.Responder {
		updateGroupResp, errResp := processUpdateGroupInfoRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewUpdateGroupInfoV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewUpdateGroupInfoV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewUpdateGroupInfoV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewUpdateGroupInfoV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewUpdateGroupInfoV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewUpdateGroupInfoV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewUpdateGroupInfoV1OK()
		resp.SetPayload(updateGroupResp)
		return resp
	})
}

func processUpdateGroupInfoRequest(db *gorm.DB, params groups_v1.UpdateGroupInfoV1Params) (resp *models.GroupObj, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	// check if exists
	groupObj, errResp := GetGroupObj(db, params.GroupID)
	if errResp != nil {
		return
	}

	// check if owner
	if groupObj.GroupOwner != payload.Username {
		errResp = commonutils.GenerateErrResp(http.StatusForbidden, " only group owner can update group info ")
		return
	}

	// update
	if errResp = updateGroupInfo(db, params); errResp != nil {
		return
	}

	// get group obj again
	resp, errResp = GetGroupObj(db, params.GroupID)
	return
}

func updateGroupInfo(db *gorm.DB, params groups_v1.UpdateGroupInfoV1Params) (errResp *models.ErrResponse) {
	groupMap := make(map[string]interface{})
	if strings.TrimSpace(params.Body.GroupName) != "" {
		groupMap["group_name"] = params.Body.GroupName
	}
	if strings.TrimSpace(params.Body.Description) != "" {
		groupMap["description"] = params.Body.Description
	}
	err := db.Table(dbpackages.GroupTable).Where("group_id = ?", params.GroupID).Update(groupMap).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
	}
	return
}

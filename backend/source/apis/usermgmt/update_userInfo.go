package usermgmt

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_mgmt_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func UpdateUserInfoV1Handler(db *gorm.DB) user_mgmt_v1.UpdateUserInfoV1HandlerFunc {
	return func(params user_mgmt_v1.UpdateUserInfoV1Params) (responder middleware.Responder) {
		updateUserInfoResp, errResp := processUpdateUserInfoRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_mgmt_v1.NewUpdateUserInfoV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_mgmt_v1.NewUpdateUserInfoV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_mgmt_v1.NewUpdateUserInfoV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_mgmt_v1.NewUpdateUserInfoV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_mgmt_v1.NewUpdateUserInfoV1Conflict().WithPayload(errResp)
			default:
				return user_mgmt_v1.NewUpdateUserInfoV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_mgmt_v1.NewUpdateUserInfoV1OK()
		resp.SetPayload(updateUserInfoResp)
		return resp
	}
}

func processUpdateUserInfoRequest(db *gorm.DB, params user_mgmt_v1.UpdateUserInfoV1Params) (resp *models.UserObj, errResp *models.ErrResponse) {
	return
}

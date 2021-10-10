package usermgmt

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_mgmt_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetUserV1Handler(db *gorm.DB) user_mgmt_v1.GetUserV1HandlerFunc {
	return func(params user_mgmt_v1.GetUserV1Params) (responder middleware.Responder) {
		userResp, errResp := processGetUserRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_mgmt_v1.NewGetUserV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_mgmt_v1.NewGetUserV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_mgmt_v1.NewGetUserV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_mgmt_v1.NewGetUserV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_mgmt_v1.NewGetUserV1Conflict().WithPayload(errResp)
			default:
				return user_mgmt_v1.NewGetUserV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_mgmt_v1.NewGetUserV1OK()
		resp.SetPayload(userResp)
		return resp
	}
}

func processGetUserRequest(db *gorm.DB, params user_mgmt_v1.GetUserV1Params) (resp *models.UserObj, errResp *models.ErrResponse) {
	return
}

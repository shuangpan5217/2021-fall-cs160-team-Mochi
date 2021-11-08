package usermgmt

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_mgmt_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func SearchUserV1Handler(db *gorm.DB) user_mgmt_v1.SearchUserV1HandlerFunc {
	return func(params user_mgmt_v1.SearchUserV1Params) (responder middleware.Responder) {
		searchUserResp, errResp := processSearchUserRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_mgmt_v1.NewSearchUserV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_mgmt_v1.NewSearchUserV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_mgmt_v1.NewSearchUserV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_mgmt_v1.NewSearchUserV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_mgmt_v1.NewSearchUserV1Conflict().WithPayload(errResp)
			default:
				return user_mgmt_v1.NewSearchUserV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_mgmt_v1.NewSearchUserV1OK()
		resp.SetPayload(searchUserResp)
		return resp
	}
}

func processSearchUserRequest(db *gorm.DB, params user_mgmt_v1.SearchUserV1Params) (userObj *models.UserObj, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	searchUser := params.Username
	userObj, errResp = GetUserObj(db, searchUser)
	return
}

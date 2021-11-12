package usermgmt

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_mgmt_v1"
	"net/http"
	"strings"

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
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username

	m := make(map[string]string)
	if strings.TrimSpace(params.Body.Description) != "" {
		m["description"] = params.Body.Description
	}

	if strings.TrimSpace(params.Body.Email) != "" {
		m["email"] = params.Body.Email
	}

	if strings.TrimSpace(params.Body.FirstName) != "" {
		m["first_name"] = params.Body.FirstName
	}

	if strings.TrimSpace(params.Body.LastName) != "" {
		m["last_Name"] = params.Body.LastName
	}

	if strings.TrimSpace(params.Body.MiddleName) != "" {
		m["middle_name"] = params.Body.MiddleName
	}

	err := db.Table(dbpackages.UserTable).Where("username = ?", username).Update(m).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp, errResp = GetUserObj(db, username)

	return
}

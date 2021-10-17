package usermgmt

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_mgmt_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func UpdatePasswordV1Handler(db *gorm.DB) user_mgmt_v1.UpdatePasswordV1HandlerFunc {
	return func(params user_mgmt_v1.UpdatePasswordV1Params) (responder middleware.Responder) {
		updatePasswordResp, errResp := processUpdatePasswordRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_mgmt_v1.NewUpdatePasswordV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_mgmt_v1.NewUpdatePasswordV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_mgmt_v1.NewUpdatePasswordV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_mgmt_v1.NewUpdatePasswordV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_mgmt_v1.NewUpdatePasswordV1Conflict().WithPayload(errResp)
			default:
				return user_mgmt_v1.NewUpdatePasswordV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_mgmt_v1.NewUpdatePasswordV1OK()
		resp.SetPayload(updatePasswordResp)
		return resp
	}
}

func processUpdatePasswordRequest(db *gorm.DB, params user_mgmt_v1.UpdatePasswordV1Params) (userObj *models.UserObj, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username

	m := map[string]string{
		"password": params.Password,
	}
	err := db.Table(dbpackages.UserTable).Where("username = ?", username).Update(m).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	userObj, errResp = GetUserObj(db, username)
	return
}

func GetUserObj(db *gorm.DB, username string) (userObj *models.UserObj, errResp *models.ErrResponse) {
	userObj = &models.UserObj{}
	err := db.Table(dbpackages.UserTable).Where("username = ?", username).First(userObj).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

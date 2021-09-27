package login

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/login_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func LoginV1Handler(db *gorm.DB) login_v1.LoginV1HandlerFunc {
	return func(params login_v1.LoginV1Params) (responder middleware.Responder) {
		storeUserInDB(db, params)

		resp := login_v1.NewLoginV1OK()
		username, errResp := storeUserInDB(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusInternalServerError:
				return login_v1.NewLoginV1InternalServerError().WithPayload(errResp)
			}
		}
		respStruct := models.LoginResponse{
			Username: username,
		}
		resp.SetPayload(&respStruct)
		return resp
	}
}

func storeUserInDB(db *gorm.DB, params login_v1.LoginV1Params) (username string, errResp *models.ErrResponse) {
	username = *params.Body.Username
	password := *params.Body.Password
	if *params.Signup == true {
		user := dbpackages.User{
			Username: username,
			Password: password,
		}
		tx := db.Begin()
		err := tx.Save(&user).Error
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return username, errResp
		}
		if err = tx.Commit().Error; err != nil {
			tx.Rollback()
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
	} else {

	}
	return
}

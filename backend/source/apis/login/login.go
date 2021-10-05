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
		username, errResp := processLoginRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusInternalServerError:
				return login_v1.NewLoginV1InternalServerError().WithPayload(errResp)
			}
		}

		resp := login_v1.NewLoginV1OK()
		respStruct := models.LoginResponse{
			Username: username,
		}
		resp.SetPayload(&respStruct)
		return resp
	}
}

func processLoginRequest(db *gorm.DB, params login_v1.LoginV1Params) (username string, errResp *models.ErrResponse) {
	if *params.Signup {
		return handleSignup(db, params)
	}
	return handleLogin(db, params)
}

// update later
func handleLogin(db *gorm.DB, params login_v1.LoginV1Params) (username string, errResp *models.ErrResponse) {
	username = *params.Body.Username
	password := *params.Body.Password

	_, errResp = checkIfPasswordCorrect(db, username, password)

	// if username and password does not match
	if errResp != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, "Username or password is incorrect")
		return
	}

	return
}

func handleSignup(db *gorm.DB, params login_v1.LoginV1Params) (username string, errResp *models.ErrResponse) {
	username = *params.Body.Username
	password := *params.Body.Password

	// check if user exists
	var user dbpackages.User
	user, errResp = checkIfUserExist(db, username)
	if errResp != nil {
		return
	}
	if user.Username != "" {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, "username already exists")
		return
	}

	// store to db if not exists
	user = dbpackages.User{
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
	return
}

func checkIfUserExist(db *gorm.DB, username string) (user dbpackages.User, errResp *models.ErrResponse) {
	err := db.Table("users").Where("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

func checkIfPasswordCorrect(db *gorm.DB, username string, password string) (user dbpackages.User, errResp *models.ErrResponse) {
	err := db.Table("users").Where("username = ? AND password = ?", username, password).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

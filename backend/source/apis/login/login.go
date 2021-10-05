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
		loginResp, errResp := processLoginRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return login_v1.NewLoginV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return login_v1.NewLoginV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return login_v1.NewLoginV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return login_v1.NewLoginV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return login_v1.NewLoginV1Conflict().WithPayload(errResp)
			default:
				return login_v1.NewLoginV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := login_v1.NewLoginV1OK()
		resp.SetPayload(loginResp)
		return resp
	}
}

func processLoginRequest(db *gorm.DB, params login_v1.LoginV1Params) (resp *models.LoginResponse, errResp *models.ErrResponse) {
	if *params.Signup {
		return handleSignup(db, params)
	}
	return handleLogin(db, params)
}

// update later

func handleLogin(db *gorm.DB, params login_v1.LoginV1Params) (resp *models.LoginResponse, errResp *models.ErrResponse) {
	username := *params.Body.Username
	// useranme exists
	// if not exists, return (password or username is not correct)
	// if exists, check password
	//     if correct, return username
	//     if not correct, (password or username is not correct)
	tokenString, errResp := commonutils.GenerateJWT(username)
	if errResp != nil {
		return
	}
	resp = &models.LoginResponse{
		Token:    tokenString,
		Username: username,
	}

	return
}

func handleSignup(db *gorm.DB, params login_v1.LoginV1Params) (resp *models.LoginResponse, errResp *models.ErrResponse) {
	username := *params.Body.Username
	password := *params.Body.Password

	// check if user exists
	var user dbpackages.User
	user, errResp = checkIfUserExist(db, username)
	if errResp != nil {
		return
	}
	if user.Username != "" {
		errResp = commonutils.GenerateErrResp(http.StatusConflict, "username already exists")
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
		return
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.LoginResponse{
		Username: username,
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



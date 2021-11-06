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

func LoginV1Handler(db *gorm.DB) user_mgmt_v1.LoginV1HandlerFunc {
	return func(params user_mgmt_v1.LoginV1Params) (responder middleware.Responder) {
		loginResp, errResp := processLoginRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_mgmt_v1.NewLoginV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_mgmt_v1.NewLoginV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_mgmt_v1.NewLoginV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_mgmt_v1.NewLoginV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_mgmt_v1.NewLoginV1Conflict().WithPayload(errResp)
			default:
				return user_mgmt_v1.NewLoginV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_mgmt_v1.NewLoginV1OK()
		resp.SetPayload(loginResp)
		return resp
	}
}

func processLoginRequest(db *gorm.DB, params user_mgmt_v1.LoginV1Params) (resp *models.LoginResponse, errResp *models.ErrResponse) {
	if *params.Signup {
		return handleSignup(db, params)
	}
	return handleLogin(db, params)
}

func handleLogin(db *gorm.DB, params user_mgmt_v1.LoginV1Params) (resp *models.LoginResponse, errResp *models.ErrResponse) {
	username := *params.Body.Username
	password := *params.Body.Password

	_, errResp = checkIfPasswordCorrect(db, username, password)

	if errResp != nil {
		return
	}

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

func handleSignup(db *gorm.DB, params user_mgmt_v1.LoginV1Params) (resp *models.LoginResponse, errResp *models.ErrResponse) {
	body := *params.Body
	username := *body.Username
	password := *body.Password

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
		Username:    username,
		Password:    password,
		Email:       body.Email,
		Description: body.Description,
		FirstName:   body.FirstName,
		MiddleName:  body.MiddleName,
		LastName:    body.LastName,
	}
	tx := db.Begin()
	err := tx.Save(&user).Error
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
	resp = &models.LoginResponse{
		Username: username,
	}
	return
}

func checkIfUserExist(db *gorm.DB, username string) (user dbpackages.User, errResp *models.ErrResponse) {
	err := db.Table(dbpackages.UserTable).Where("username = ?", username).First(&user).Error
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
		errResp = commonutils.GenerateErrResp(http.StatusUnauthorized, "Username or password is incorrect") // 401 error
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

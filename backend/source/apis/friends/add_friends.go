package friends

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/friends_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func AddFriendsV1Handler(db *gorm.DB) friends_v1.AddFriendsV1HandlerFunc {
	return func(params friends_v1.AddFriendsV1Params) (responder middleware.Responder) {
		friendResp, errResp := processAddfriendsRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return friends_v1.NewAddFriendsV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return friends_v1.NewAddFriendsV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return friends_v1.NewAddFriendsV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return friends_v1.NewAddFriendsV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return friends_v1.NewAddFriendsV1Conflict().WithPayload(errResp)
			default:
				return friends_v1.NewAddFriendsV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := friends_v1.NewAddFriendsV1OK()
		resp.SetPayload(friendResp)
		return resp
	}
}

func processAddfriendsRequest(db *gorm.DB, params friends_v1.AddFriendsV1Params) (resp *models.FriendResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username

	body := *params.Body
	username2 := body.Username2

	if username == username2 {
		errResp = commonutils.GenerateErrResp(http.StatusConflict, "Friend name and username cannot be the same")
		return
	}

	friend, errResp := checkIfFriendExist(db, username, username2)
	if errResp != nil {
		return
	}
	if friend.Username != "" {
		errResp = commonutils.GenerateErrResp(http.StatusConflict, "friend already exists")
		return
	}

	friends := dbpackages.Friend{
		Username:  username,
		Username2: username2,
	}

	friends2 := dbpackages.Friend{
		Username:  username2,
		Username2: username,
	}
	tx := db.Begin()
	err := tx.Save(&friends).Error
	if err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	err = tx.Save(&friends2).Error
	if err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.FriendResponse{
		Username: username2,
	}

	return
}

func checkIfFriendExist(db *gorm.DB, username string, username2 string) (friend dbpackages.Friend, errResp *models.ErrResponse) {
	err := db.Table("friends").Where("(username = ? AND username2 = ?) OR (username2 = ? AND username = ?)", username, username2, username, username2).First(&friend).Error
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

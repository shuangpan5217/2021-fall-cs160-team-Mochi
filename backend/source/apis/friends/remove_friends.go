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

func RemoveFriendsV1Handler(db *gorm.DB) friends_v1.RemoveFriendsV1HandlerFunc {
	return func(params friends_v1.RemoveFriendsV1Params) (responder middleware.Responder) {
		removeFriendObject, errResp := processRemovefriendsRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return friends_v1.NewRemoveFriendsV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return friends_v1.NewRemoveFriendsV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return friends_v1.NewRemoveFriendsV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return friends_v1.NewRemoveFriendsV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return friends_v1.NewRemoveFriendsV1Conflict().WithPayload(errResp)
			default:
				return friends_v1.NewRemoveFriendsV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := friends_v1.NewRemoveFriendsV1OK()
		resp.SetPayload(removeFriendObject)
		return resp
	}
}

func processRemovefriendsRequest(db *gorm.DB, params friends_v1.RemoveFriendsV1Params) (resp *models.FriendResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username
	username2 := params.Username
	if username == username2 {
		errResp = commonutils.GenerateErrResp(http.StatusConflict, "Friend name and username cannot be the same")
		return
	}
	db = db.Table(dbpackages.FriendTable).Where("(username = ? and username2 = ?) OR (username2 = ? and username = ?)", username, username2, username, username2).Delete(&dbpackages.Friend{})
	if db.Error != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, db.Error.Error())
		return
	} else if db.RowsAffected < 1 {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
		return
	}
	resp = &models.FriendResponse{
		Username: username2,
	}

	return
}

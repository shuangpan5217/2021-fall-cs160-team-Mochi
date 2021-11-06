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

func GetFriendsV1Handler(db *gorm.DB) friends_v1.GetFriendsV1HandlerFunc {
	return func(params friends_v1.GetFriendsV1Params) (responder middleware.Responder) {
		getFriendObject, errResp := processGetfriendsRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return friends_v1.NewGetFriendsV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return friends_v1.NewGetFriendsV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return friends_v1.NewGetFriendsV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return friends_v1.NewGetFriendsV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return friends_v1.NewGetFriendsV1Conflict().WithPayload(errResp)
			default:
				return friends_v1.NewGetFriendsV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := friends_v1.NewGetFriendsV1OK()
		resp.SetPayload(getFriendObject)
		return resp
	}
}

func processGetfriendsRequest(db *gorm.DB, params friends_v1.GetFriendsV1Params) (resp *models.GetFriendObject, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username

	resp = &models.GetFriendObject{
		Friends: []*models.FriendResponse{},
	}

	rawSQL := `SELECT username2, username as username FROM friends WHERE username = ? or username2 = ?`

	rows, err := db.Raw(rawSQL, username, username).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	for rows.Next() {
		friends := dbpackages.Friend{}
		err = db.ScanRows(rows, &friends)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		newfriends := models.FriendResponse{}
		if friends.Username == payload.Username {
			newfriends.Username = friends.Username2
			resp.Friends = append(resp.Friends, &newfriends)
		}
		if friends.Username2 == payload.Username {
			newfriends.Username = friends.Username
			resp.Friends = append(resp.Friends, &newfriends)
		}

	}
	return

}

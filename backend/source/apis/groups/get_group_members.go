package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetGroupMembersHandler(db *gorm.DB) groups_v1.GetGroupUsersV1HandlerFunc {
	return func(params groups_v1.GetGroupUsersV1Params) middleware.Responder {
		getGroupMembersObject, errResp := processGetGroupMembersRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewGetGroupUsersV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewGetGroupUsersV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewGetGroupUsersV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewGetGroupUsersV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewGetGroupUsersV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewGetGroupUsersV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewGetGroupUsersV1OK()
		resp.SetPayload(getGroupMembersObject)
		return resp
	}
}

func processGetGroupMembersRequest(db *gorm.DB, params groups_v1.GetGroupUsersV1Params) (resp *models.GroupMembersObject, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	groupid := params.GroupID

	resp = &models.GroupMembersObject{
		Users: []*models.UserObj{},
	}

	rawSQL := `SELECT distinct g.username, g.last_name, g.middle_name, g.first_name, g.email, g.description
				FROM group_users gu, users g 
				WHERE group_id = ? and gu.username = g.username`

	rows, err := db.Raw(rawSQL, groupid).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	for rows.Next() {
		users := models.UserObj{}
		err = db.ScanRows(rows, &users)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Users = append(resp.Users, &users)

	}
	return

}

package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetGroupsV1Handler(db *gorm.DB) groups_v1.GetGroupsV1HandlerFunc {
	return func(params groups_v1.GetGroupsV1Params) middleware.Responder {
		getAllGroupsObject, errResp := processGetAllGroupsRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewGetGroupsV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewGetGroupsV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewGetGroupsV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewGetGroupsV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewGetGroupsV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewGetGroupsV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewGetGroupsV1OK()
		resp.SetPayload(getAllGroupsObject)
		return resp
	}
}

func processGetAllGroupsRequest(db *gorm.DB, params groups_v1.GetGroupsV1Params) (resp *models.GetAllGroupsObject, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	username := payload.Username
	resp = &models.GetAllGroupsObject{
		AllGroups: []*models.GroupObj{},
	}
	// groupusers := []dbpackages.GroupUser{}
	// err :=db.Table(dbpackages.GroupUserTable).Where("username = ?", username).Find(&groupusers).Error
	// if err != nil {
	// 	errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	rawSQL := `select distinct g.group_id, g.group_name, g.group_owner, g.description
				from group_users gu, groups g
				where gu.username = ? and gu.group_id = g.group_id`

	rows, err := db.Raw(rawSQL, username).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	for rows.Next() {
		groups := models.GroupObj{}
		err = db.ScanRows(rows, &groups)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		resp.AllGroups = append(resp.AllGroups, &groups)
	}

	return
}

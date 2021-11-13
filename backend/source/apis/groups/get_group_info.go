package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetGroupInfoV1Handler(db *gorm.DB) groups_v1.GetGroupInfoV1HandlerFunc {
	return func(params groups_v1.GetGroupInfoV1Params) middleware.Responder {
		groupInfoObj, errResp := processGetGroupInfoRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewGetGroupInfoV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewGetGroupInfoV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewGetGroupInfoV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewGetGroupInfoV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewGetGroupInfoV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewGetGroupInfoV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewGetGroupInfoV1OK()
		resp.SetPayload(groupInfoObj)
		return resp
	}
}

func processGetGroupInfoRequest(db *gorm.DB, params groups_v1.GetGroupInfoV1Params) (resp *models.GroupObj, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username
	groupid := params.GroupID

	err := db.Table(dbpackages.UserTable).Where("username = ?", username).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp, errResp = GetGroupObj(db, groupid)
	return
}
func GetGroupObj(db *gorm.DB, path string) (groupObj *models.GroupObj, errResp *models.ErrResponse) {
	groupObj = &models.GroupObj{}
	err := db.Table(dbpackages.GroupTable).Where("group_id = ?", path).First(groupObj).Error
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return

}

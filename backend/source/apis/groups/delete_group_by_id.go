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

func DeleteGroupByIdV1Handler(db *gorm.DB) groups_v1.DeleteGroupV1HandlerFunc {
	return func(params groups_v1.DeleteGroupV1Params) middleware.Responder {
		deleteGroupByIdResp, errResp := processdeleteGroupByIdRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return groups_v1.NewDeleteGroupV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return groups_v1.NewDeleteGroupV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return groups_v1.NewDeleteGroupV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return groups_v1.NewDeleteGroupV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return groups_v1.NewDeleteGroupV1Conflict().WithPayload(errResp)
			default:
				return groups_v1.NewDeleteGroupV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := groups_v1.NewDeleteGroupV1OK()
		resp.SetPayload(deleteGroupByIdResp)
		return resp
	}
}

func processdeleteGroupByIdRequest(db *gorm.DB, params groups_v1.DeleteGroupV1Params) (resp *models.GroupResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	username := payload.Username
	groupID := params.GroupID

	//Check if group owner
	_, err := checkIfGroupOwner(db, payload.Username, params.GroupID)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusForbidden, " not a group member ")
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// check if group exists
	_, err = checkIfGroupExists(db, params.GroupID)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// delete group by id
	db = db.Table(dbpackages.GroupTable).Where("group_id = ? and group_owner = ?", groupID, username).Delete(&dbpackages.Group{})
	if db.Error != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, db.Error.Error())
		return
	} else if db.RowsAffected < 1 {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
		return
	}
	resp = &models.GroupResponse{
		GroupID: params.GroupID,
	}
	return
}
func checkIfGroupOwner(db *gorm.DB, groupowner, groupID string) (groupOwner dbpackages.Group, err error) {
	err = db.Table(dbpackages.GroupTable).Where("group_owner = ? AND group_id = ?", groupowner, groupID).First(&groupOwner).Error
	return
}

func checkIfGroupExists(db *gorm.DB, groupID string) (group dbpackages.Group, err error) {
	err = db.Table(dbpackages.GroupTable).Where("group_id = ?", groupID).First(&group).Error
	return
}

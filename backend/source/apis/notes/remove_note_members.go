package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func RemoveNoteMembersV1Handler(db *gorm.DB) notes_v1.RemoveNoteMembersV1HandlerFunc {
	return func(params notes_v1.RemoveNoteMembersV1Params) middleware.Responder {
		removeNoteMembersResp, errResp := processRemoveNoteMembersRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewRemoveNoteMembersV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewRemoveNoteMembersV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewRemoveNoteMembersV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewRemoveNoteMembersV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewRemoveNoteMembersV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewRemoveNoteMembersV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewRemoveNoteMembersV1OK()
		resp.SetPayload(removeNoteMembersResp)
		return resp
	}
}

func processRemoveNoteMembersRequest(db *gorm.DB, params notes_v1.RemoveNoteMembersV1Params) (resp *models.NoteResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	// check if user note exists
	_, err := checkIfUserIsNoteOwner(db, params.ID, payload.Username)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusForbidden, " no access to the note ")
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	users := params.Body.Users
	groups := params.Body.Groups

	// delete
	tx := db.Begin()
	for _, user := range users {
		if user.Username == payload.Username {
			errResp = commonutils.GenerateErrResp(http.StatusBadRequest, "not able to remove note owner from the note")
			tx.Rollback()
			return
		}
		err = tx.Table(dbpackages.UserNoteTable).Where("username = ? AND note_id = ?", user.Username, params.ID).Delete(&dbpackages.UserNote{}).Error
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			tx.Rollback()
			return
		}
	}
	for _, group := range groups {
		err = tx.Table(dbpackages.GroupNoteTable).Where("group_id = ? AND note_id = ?", group.GroupID, params.ID).Delete(&dbpackages.UserNote{}).Error
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			tx.Rollback()
			return
		}
	}
	if err = tx.Commit().Error; err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	resp = &models.NoteResponse{
		NoteID: params.ID,
	}
	return
}

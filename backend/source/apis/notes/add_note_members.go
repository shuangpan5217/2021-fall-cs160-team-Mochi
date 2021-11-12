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

func AddNoteMembersV1Handler(db *gorm.DB) notes_v1.AddNoteMembersV1HandlerFunc {
	return func(params notes_v1.AddNoteMembersV1Params) middleware.Responder {
		addNoteMembersResp, errResp := processAddNoteMembersRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewAddNoteMembersV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewAddNoteMembersV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewAddNoteMembersV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewAddNoteMembersV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewAddNoteMembersV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewAddNoteMembersV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewAddNoteMembersV1OK()
		resp.SetPayload(addNoteMembersResp)
		return resp
	}
}

func processAddNoteMembersRequest(db *gorm.DB, params notes_v1.AddNoteMembersV1Params) (resp *models.NoteResponse, errResp *models.ErrResponse) {
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

	// insert
	tx := db.Begin()
	for _, user := range users {
		err = tx.Save(&dbpackages.UserNote{NoteID: params.ID, Username: user.Username}).Error
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			tx.Rollback()
			return
		}
	}
	for _, group := range groups {
		err = tx.Save(&dbpackages.GroupNote{NoteID: params.ID, GroupID: group.GroupID}).Error
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

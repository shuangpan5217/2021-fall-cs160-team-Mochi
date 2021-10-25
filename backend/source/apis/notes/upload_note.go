package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func UploadNoteV1Handler(db *gorm.DB) notes_v1.UploadNoteV1HandlerFunc {
	return func(params notes_v1.UploadNoteV1Params) middleware.Responder {
		uploadNoteResp, errResp := processUploadNoteRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewUploadNoteV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewUploadNoteV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewUploadNoteV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewUploadNoteV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewUploadNoteV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewUploadNoteV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewUploadNoteV1OK()
		resp.SetPayload(uploadNoteResp)
		return resp
	}
}

func processUploadNoteRequest(db *gorm.DB, params notes_v1.UploadNoteV1Params) (resp *models.NoteResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	noteID := uuid.New().String()
	errResp = checkIfNoteExists(db, noteID)
	if errResp != nil {
		return
	}
	// insert notes
	newNote := dbpackages.Note{
		NoteID:        noteID,
		NoteOwner:     payload.Username,
		Description:   params.Body.Description,
		Title:         params.Body.Title,
		Style:         *params.Body.Style,
		NoteReference: *params.Body.NoteReference,
		Type:          *params.Body.Type,
		Tag:           *params.Body.Tag,
	}
	err := db.Table(dbpackages.NoteTable).Save(&newNote).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	// insert user_notes
	userNote := dbpackages.UserNote{
		Username: payload.Username,
		NoteID:   noteID,
	}
	err = db.Table(dbpackages.UserNoteTable).Save(&userNote).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return &models.NoteResponse{NoteID: noteID}, nil
}

func checkIfNoteExists(db *gorm.DB, noteID string) (errResp *models.ErrResponse) {
	var note dbpackages.Note
	err := db.Table(dbpackages.NoteTable).Where("note_id = ?", noteID).First(&note).Error
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	return
}

package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetNotesByUsernameHandler(db *gorm.DB) notes_v1.FindNotesByUsernameHandlerFunc {
	return func(params notes_v1.FindNotesByUsernameParams) middleware.Responder {
		findNotesByUsernameResp, errResp := processGetNotesByUsernameRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewFindNotesByUsernameBadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewFindNotesByUsernameUnauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewFindNotesByUsernameForbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewFindNotesByUsernameNotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewFindNotesByUsernameConflict().WithPayload(errResp)
			default:
				return notes_v1.NewFindNotesByUsernameInternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewFindNotesByUsernameOK()
		resp.SetPayload(findNotesByUsernameResp)
		return resp
	}
}

func processGetNotesByUsernameRequest(db *gorm.DB, params notes_v1.FindNotesByUsernameParams) (resp *models.NotesGetResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username

	resp = &models.NotesGetResponse{
		Notes: []*models.NoteObjectResponse{},
	}

	// get both username's notes and its shared notes
	rawSQL := `SELECT n.note_owner, n.description, n.title, n.note_reference, n.type, n.tag, n.note_id, n.style
				FROM notes n, user_notes un
				WHERE un.username = ? AND n.note_id = un.note_id`

	rows, err := db.Raw(rawSQL, username).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	for rows.Next() {
		note := models.NoteObjectResponse{}
		err = db.ScanRows(rows, &note)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Notes = append(resp.Notes, &note)
	}
	return
}

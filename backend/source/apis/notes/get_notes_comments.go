package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetNoteCommentsHandler(db *gorm.DB) notes_v1.GetNoteCommentsHandlerFunc {
	return func(params notes_v1.GetNoteCommentsParams) middleware.Responder {
		getNotesCommentsResp, errResp := processGetNoteCommentsRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewGetNoteCommentsBadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewGetNoteCommentsUnauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewGetNoteCommentsForbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewGetNoteCommentsNotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewGetNoteCommentsConflict().WithPayload(errResp)
			default:
				return notes_v1.NewGetNoteCommentsInternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewGetNoteCommentsOK()
		resp.SetPayload(getNotesCommentsResp)
		return resp
	}
}

func processGetNoteCommentsRequest(db *gorm.DB, params notes_v1.GetNoteCommentsParams) (resp *models.NoteCommentsResponse, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	noteid := params.NoteID

	resp = &models.NoteCommentsResponse{
		Comments: []*models.CommentObject{},
	}

	rawSQL := `SELECT DISTINCT(comment_id), content, note_ID, username
				FROM comments
				WHERE note_id = ?`

	rows, err := db.Raw(rawSQL, noteid).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	for rows.Next() {
		comments := models.CommentObject{}
		err = db.ScanRows(rows, &comments)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Comments = append(resp.Comments, &comments)
	}
	return
}

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

func GetSingleNoteV1Handler(db *gorm.DB) notes_v1.GetSingleNoteV1HandlerFunc {
	return func(params notes_v1.GetSingleNoteV1Params) middleware.Responder {
		getSingleNoteResp, errResp := processGetSingleNoteRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewGetSingleNoteV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewGetSingleNoteV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewGetSingleNoteV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewGetSingleNoteV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewGetSingleNoteV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewGetSingleNoteV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewGetSingleNoteV1OK()
		resp.SetPayload(getSingleNoteResp)
		return resp
	}
}

func processGetSingleNoteRequest(db *gorm.DB, params notes_v1.GetSingleNoteV1Params) (resp *models.NoteObjectResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	username := payload.Username

	// check if note exists
	note, err := checkIfNoteExists(db, params.ID)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	resp = &models.NoteObjectResponse{
		Description:   note.Description,
		NoteID:        note.NoteID,
		NoteOwner:     note.NoteOwner,
		NoteReference: note.NoteReference,
		Type:          note.Type,
		Title:         note.Title,
		Tag:           note.Tag,
		Style:         note.Style,
	}
	if note.Type == "public" || note.NoteOwner == username {
		return
	}

	// check shared note
	errResp = checkIfUsernameExists(db, username, params.ID)
	if errResp != nil {
		return
	}

	return
}

func checkIfUsernameExists(db *gorm.DB, username, ID string) (errResp *models.ErrResponse) {
	// get all usernames
	rows, err := db.Table(dbpackages.UserNoteTable).Where("note_id = ?", ID).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	// check if username exists in user notes rows
	var exists = false
	for rows.Next() {
		user := models.UserObj{}
		err = db.ScanRows(rows, &user)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		if user.Username == username {
			exists = true
			break
		}
	}
	// shared
	if exists {
		return
	}

	// check if username exists in group notes rows
	rawSQL := `SELECT count(*)
			FROM notes n, 
			(
				SELECT gu.username, gu.group_id, gn.note_id
				FROM group_users gu inner join group_notes gn on gu.group_id = gn.group_id
			) AS gun
			WHERE gun.note_id = n.note_id AND gun.username = ?`
	var totalCount int64 = 0
	err = db.Raw(rawSQL, username).Count(&totalCount).Error
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// not shared
	if totalCount == 0 {
		errResp = commonutils.GenerateErrResp(http.StatusForbidden, "Forbidden: No access to the note")
		return
	}

	return
}

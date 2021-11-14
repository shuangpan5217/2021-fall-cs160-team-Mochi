package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func FindNoteByTagV1Handler(db *gorm.DB) notes_v1.FindNotesByTagsHandlerFunc {
	return func(params notes_v1.FindNotesByTagsParams) middleware.Responder {
		findNotesByTagsResp, errResp := processFindNotesByTagsRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewFindNotesByTagsBadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewFindNotesByTagsUnauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewFindNotesByTagsForbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewFindNotesByTagsNotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewFindNotesByTagsConflict().WithPayload(errResp)
			default:
				return notes_v1.NewFindNotesByTagsInternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewFindNotesByTagsOK()
		resp.SetPayload(findNotesByTagsResp)
		return resp
	}
}

func processFindNotesByTagsRequest(db *gorm.DB, params notes_v1.FindNotesByTagsParams) (resp *models.NotesGetResponse, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	var err error
	notes := []dbpackages.Note{}

	// tags
	tagsDraft := strings.Split(params.Tags, ",")
	var tags []string
	for _, tagDraft := range tagsDraft {
		tag := strings.TrimSpace(tagDraft)
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if len(tags) == 0 {
		errResp = commonutils.GenerateErrResp(http.StatusBadRequest, " correct tags format: [tag1,tag2,tag3...]")
		return
	}

	resp = &models.NotesGetResponse{
		Notes: []*models.NoteObjectResponse{},
	}

	// get where statement and args
	whereStatement := `type = ? AND (`
	var args []interface{}
	args = append(args, "public") // public tag
	for i, tag := range tags {
		if i == len(tags)-1 {
			whereStatement += `tag like ?)`
		} else {
			whereStatement += `tag like ? OR `
		}
		args = append(args, "%"+tag+"%")
	}
	selectColumns := "DISTINCT ON (note_reference) note_reference, note_id, description, note_owner, style, tag, title, type, updated_at"

	// run query
	if *params.UpdatedAt {
		err = db.Table(dbpackages.NoteTable).
			Where(whereStatement, args...).
			Select(selectColumns).
			Order("note_reference, updated_at").
			Offset(*params.Offset).
			Limit(*params.Limit).
			Find(&notes).Error
	} else {
		err = db.Table(dbpackages.NoteTable).
			Where(whereStatement, args...).
			Select(selectColumns).
			Order("note_reference").
			Offset(*params.Offset).
			Limit(*params.Limit).
			Find(&notes).Error
	}

	// err resp
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusNotFound, " record not found ")
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// format
	for _, note := range notes {
		noteResp := models.NoteObjectResponse{
			Description:   note.Description,
			NoteID:        note.NoteID,
			NoteOwner:     note.NoteOwner,
			NoteReference: note.NoteReference,
			Style:         note.Style,
			Tag:           note.Tag,
			Title:         note.Title,
			Type:          note.Type,
		}
		resp.Notes = append(resp.Notes, &noteResp)
	}

	return
}

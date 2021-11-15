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

func GetGroupNotesV1Handler(db *gorm.DB) notes_v1.GetGroupNotesV1HandlerFunc {
	return func(params notes_v1.GetGroupNotesV1Params) middleware.Responder {
		getGroupNotesResp, errResp := processGetGroupNotesRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewGetGroupNotesV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewGetGroupNotesV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewGetGroupNotesV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewGetGroupNotesV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewGetGroupNotesV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewGetGroupNotesV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewGetGroupNotesV1OK()
		resp.SetPayload(getGroupNotesResp)
		return resp
	}
}

func processGetGroupNotesRequest(db *gorm.DB, params notes_v1.GetGroupNotesV1Params) (resp *models.NotesGetResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	// check if it is a group members
	_, err := CheckIfGroupUser(db, payload.Username, params.GroupID)
	if gorm.IsRecordNotFoundError(err) {
		errResp = commonutils.GenerateErrResp(http.StatusUnauthorized, " no access to notes ")
		return
	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// get all group notes
	selectColumns := `DISTINCT on (notes.note_id) notes.note_id, notes.note_owner, notes.description, notes.title,
					notes.type, notes.tag, notes.note_reference, notes.style`
	rows, err := db.Table(dbpackages.NoteTable).
		Joins("inner join group_notes gn on notes.note_id = gn.note_id").
		Select(selectColumns).
		Where("gn.group_id = ?", params.GroupID).Order("notes.note_id").Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// resp
	resp = &models.NotesGetResponse{
		Notes: []*models.NoteObjectResponse{},
	}

	for rows.Next() {
		note := models.NoteObjectResponse{}
		err = db.ScanRows(rows, &note)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		}
		resp.Notes = append(resp.Notes, &note)
	}

	return
}

func CheckIfGroupUser(db *gorm.DB, username, groupID string) (groupUser dbpackages.GroupUser, err error) {
	err = db.Table(dbpackages.GroupUserTable).Where("username = ? AND group_id = ?", username, groupID).First(&groupUser).Error
	return
}

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

func GetNoteMembersV1Handler(db *gorm.DB) notes_v1.GetNoteMembersV1HandlerFunc {
	return func(params notes_v1.GetNoteMembersV1Params) middleware.Responder {
		getNoteMembersResp, errResp := processGetNoteMembersRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewGetNoteMembersV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewGetNoteMembersV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewGetNoteMembersV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewGetNoteMembersV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewGetNoteMembersV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewGetNoteMembersV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewGetNoteMembersV1OK()
		resp.SetPayload(getNoteMembersResp)
		return resp
	}
}

func processGetNoteMembersRequest(db *gorm.DB, params notes_v1.GetNoteMembersV1Params) (resp *models.GetNoteMembersResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	// check if user note exists
	note, err := checkIfUserIsNoteOwner(db, params.ID, payload.Username)
	if gorm.IsRecordNotFoundError(err) {

	} else if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	// only note owner can see all members
	if note.NoteOwner == "" {
		errResp = commonutils.GenerateErrResp(http.StatusForbidden, "no access to members of the note")
		return
	}

	resp = &models.GetNoteMembersResponse{
		Users:  []*models.UserObj{},
		Groups: []*models.GroupObj{},
	}

	rawUsersSQL := `SELECT DISTINCT on (u.username) u.username, u.description, u.email, u.first_name, u.last_name, u.middle_name
					FROM users u, user_notes un
					WHERE un.note_id = ? AND un.username = u.username
					ORDER BY u.username`
	rows, err := db.Raw(rawUsersSQL, params.ID).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	for rows.Next() {
		user := models.UserObj{}
		err = db.ScanRows(rows, &user)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Users = append(resp.Users, &user)
	}

	rawGroupsSQL := `SELECT DISTINCT(g.group_id) group_id, g.group_name, g.group_owner, g.description
					FROM groups g, group_notes gn
					WHERE gn.note_id = ? AND gn.group_id = g.group_id
					ORDER by g.group_id`
	rows, err = db.Raw(rawGroupsSQL, params.ID).Rows()
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	for rows.Next() {
		group := models.GroupObj{}
		err = db.ScanRows(rows, &group)
		if err != nil {
			errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
			return
		}
		resp.Groups = append(resp.Groups, &group)
	}
	return
}

func checkIfUserIsNoteOwner(db *gorm.DB, noteID, username string) (note dbpackages.Note, err error) {
	err = db.Table(dbpackages.NoteTable).Where("note_owner = ? AND note_id = ?", username, noteID).First(&note).Error
	return note, err
}

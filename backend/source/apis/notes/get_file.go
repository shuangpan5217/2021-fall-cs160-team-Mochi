package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetFileV1Handler(db *gorm.DB) notes_v1.GetFileV1HandlerFunc {
	return func(params notes_v1.GetFileV1Params) middleware.Responder {
		postFileResp, errResp := processGetFileRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewGetFileV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewGetFileV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewGetFileV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewGetFileV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewGetFileV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewGetFileV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewGetFileV1OK()
		resp.SetPayload(postFileResp)
		return resp
	}
}

func processGetFileRequest(db *gorm.DB, params notes_v1.GetFileV1Params) (resp *models.GetFileResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	_, errResp = getNoteByFileName(db, params.Path, payload.Username)
	if errResp != nil {
		return
	}

	mochiNoteDir, errResp := commonutils.GetMochiNoteFilesDir()
	if errResp != nil {
		return
	}
	filePath := mochiNoteDir + "/" + params.Path
	pdfFileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.GetFileResponse{
		PdfData: base64.StdEncoding.EncodeToString(pdfFileData),
	}
	return
}

func getNoteByFileName(db *gorm.DB, path string, username string) (note dbpackages.Note, errResp *models.ErrResponse) {
	err := db.Table(dbpackages.NoteTable).Where("note_reference = ?", path).First(&note).Error
	if gorm.IsRecordNotFoundError(err) {
		commonutils.GenerateErrResp(http.StatusNotFound, "record not found")
		return
	} else if err != nil {
		commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}

	if note.Type == "public" || note.NoteOwner == username {

	} else {
		errResp = checkIfUsernameExists(db, username, note.NoteID)
		if errResp != nil {
			return
		}
	}
	return
}

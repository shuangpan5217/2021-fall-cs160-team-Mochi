package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func PostFileV1Handler(db *gorm.DB) notes_v1.PostFileV1HandlerFunc {
	return func(params notes_v1.PostFileV1Params) middleware.Responder {
		postFileResp, errResp := processPostFileRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewPostFileV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewPostFileV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewPostFileV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewPostFileV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewPostFileV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewPostFileV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewPostFileV1OK()
		resp.SetPayload(postFileResp)
		return resp
	}
}

func processPostFileRequest(db *gorm.DB, params notes_v1.PostFileV1Params) (resp *models.PostFileResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	// read file to bytes
	fileBytes, err := ioutil.ReadAll(params.NoteFile)
	if err != nil {
		commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	// check file type
	fileType := http.DetectContentType(fileBytes)
	if !strings.Contains(fileType, "application/pdf") {
		commonutils.GenerateErrResp(http.StatusBadRequest, "Only pdf file is allowed.")
		return
	}

	// file dir
	fileDir, errResp := commonutils.GetMochiNoteFilesDir()
	if errResp != nil {
		return
	}
	fileName := payload.Username + uuid.New().String() + ".pdf"
	exist, err := Exists(fileName)
	if exist || err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, commonutils.InternalServerError)
		return
	}
	// write dir and files
	err = os.MkdirAll(fileDir, 0777)
	if err != nil {
		commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	err = os.WriteFile(fileDir+"/"+fileName, fileBytes, 0666)
	if err != nil {
		commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = &models.PostFileResponse{
		NoteReference: fileName,
	}
	return
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

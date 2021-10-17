package notes

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
)

func GetMultipleFilesV1Handler(db *gorm.DB) notes_v1.GetMultipleFilesV1HandlerFunc {
	return func(params notes_v1.GetMultipleFilesV1Params) middleware.Responder {
		getMultipleFilesResp, errResp := processGetMultipleFilesRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return notes_v1.NewGetMultipleFilesV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return notes_v1.NewGetMultipleFilesV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return notes_v1.NewGetMultipleFilesV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return notes_v1.NewGetMultipleFilesV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return notes_v1.NewGetMultipleFilesV1Conflict().WithPayload(errResp)
			default:
				return notes_v1.NewGetMultipleFilesV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := notes_v1.NewGetMultipleFilesV1OK()
		resp.SetPayload(getMultipleFilesResp)
		return resp
	}
}

func processGetMultipleFilesRequest(db *gorm.DB, params notes_v1.GetMultipleFilesV1Params) (resp *models.GetFilesResponse, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	mochiNoteDir, errResp := commonutils.GetMochiNoteFilesDir()
	if errResp != nil {
		return
	}

	paths := params.Body.FilePaths
	filesData := []*models.GetFileResponse{}
	resp = &models.GetFilesResponse{
		FilesData: filesData,
	}

	for _, path := range paths {
		pdfData, err := ioutil.ReadFile(mochiNoteDir + "/" + path.Path)
		if err == nil {
			fileResp := &models.GetFileResponse{
				PdfData: pdfData,
			}
			resp.FilesData = append(resp.FilesData, fileResp)
		}
	}
	resp.Count = cast.ToInt32(len(resp.FilesData))
	return
}

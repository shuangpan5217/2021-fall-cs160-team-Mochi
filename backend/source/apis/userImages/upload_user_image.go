package userImages

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_images_v1"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func UploadUserImagesHandlerV1(db *gorm.DB) user_images_v1.PostUserImagesV1HandlerFunc {
	return func(params user_images_v1.PostUserImagesV1Params) middleware.Responder {
		postUserImagesResp, errResp := processPostUserImagesRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_images_v1.NewPostUserImagesV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_images_v1.NewPostUserImagesV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_images_v1.NewPostUserImagesV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_images_v1.NewPostUserImagesV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_images_v1.NewPostUserImagesV1Conflict().WithPayload(errResp)
			default:
				return user_images_v1.NewPostUserImagesV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_images_v1.NewPostUserImagesV1OK()
		resp.SetPayload(postUserImagesResp)
		return resp
	}

}

func processPostUserImagesRequest(db *gorm.DB, params user_images_v1.PostUserImagesV1Params) (resp *models.SuccessResponse, errResp *models.ErrResponse) {
	payload, errResp := commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	// read file to bytes
	fileBytes, err := ioutil.ReadAll(params.UserImage)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	// check file type
	fileType := http.DetectContentType(fileBytes)
	if !strings.Contains(fileType, "image/jpeg") {
		errResp = commonutils.GenerateErrResp(http.StatusBadRequest, "Only image/jpeg type is allowed.")
		return
	}

	// file dir
	fileDir, errResp := commonutils.GetUserImagesDir()
	if errResp != nil {
		return
	}
	fileName := payload.Username
	// write dir and files
	err = os.MkdirAll(fileDir, 0777)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	err = os.WriteFile(fileDir+"/"+fileName, fileBytes, 0666)
	if err != nil {
		errResp = commonutils.GenerateErrResp(http.StatusInternalServerError, err.Error())
		return
	}
	resp = commonutils.GenerateSuccessResp(http.StatusOK, "successfully upload profile picture")
	return
}

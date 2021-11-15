package userImages

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_images_v1"
	"encoding/base64"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetMultipleUserImages(db *gorm.DB) user_images_v1.GetMultipleUserImagesV1HandlerFunc {
	return func(params user_images_v1.GetMultipleUserImagesV1Params) middleware.Responder {
		getMultipleUserImagesResp, errResp := processGetMultipleUserImagesRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_images_v1.NewGetMultipleUserImagesV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_images_v1.NewGetMultipleUserImagesV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_images_v1.NewGetMultipleUserImagesV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_images_v1.NewGetMultipleUserImagesV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_images_v1.NewGetMultipleUserImagesV1Conflict().WithPayload(errResp)
			default:
				return user_images_v1.NewGetMultipleUserImagesV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_images_v1.NewGetMultipleUserImagesV1OK()
		resp.SetPayload(getMultipleUserImagesResp)
		return resp
	}
}

func processGetMultipleUserImagesRequest(db *gorm.DB,
	params user_images_v1.GetMultipleUserImagesV1Params) (resp *models.MultipleUserImagesResponse, errResp *models.ErrResponse) {

	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}

	// get file path
	userImagesDir, errResp := commonutils.GetUserImagesDir()
	if errResp != nil {
		return
	}

	// resp
	resp = &models.MultipleUserImagesResponse{
		Images: []*models.UserImagesResponse{},
	}

	// get user images
	users := params.Body.Users
	for _, user := range users {
		var userImageData []byte
		userImageData = readUserImage(userImagesDir, userImageData, user.Username)
		imageData := &models.UserImagesResponse{
			UserImage: base64.StdEncoding.EncodeToString(userImageData),
			Type:      http.DetectContentType(userImageData),
		}
		resp.Images = append(resp.Images, imageData)
	}
	resp.Count = int32(len(resp.Images))

	return
}

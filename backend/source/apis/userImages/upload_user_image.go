package userImages

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_images_v1"
	"net/http"

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
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	return
}

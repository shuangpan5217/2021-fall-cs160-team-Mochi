package userImages

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_images_v1"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetUserImagesHandlerV1(db *gorm.DB) user_images_v1.GetUserImagesV1HandlerFunc {
	return func(params user_images_v1.GetUserImagesV1Params) middleware.Responder {
		getUserImagesResp, errResp := processGetUserImagesRequest(db, params)
		if errResp != nil {
			switch errResp.StatusCode {
			case http.StatusBadRequest:
				return user_images_v1.NewGetUserImagesV1BadRequest().WithPayload(errResp)
			case http.StatusUnauthorized:
				return user_images_v1.NewGetUserImagesV1Unauthorized().WithPayload(errResp)
			case http.StatusForbidden:
				return user_images_v1.NewGetUserImagesV1Forbidden().WithPayload(errResp)
			case http.StatusNotFound:
				return user_images_v1.NewGetUserImagesV1NotFound().WithPayload(errResp)
			case http.StatusConflict:
				return user_images_v1.NewGetUserImagesV1Conflict().WithPayload(errResp)
			default:
				return user_images_v1.NewGetUserImagesV1InternalServerError().WithPayload(errResp)
			}
		}
		resp := user_images_v1.NewGetUserImagesV1OK()
		resp.SetPayload(getUserImagesResp)
		return resp
	}

}

func processGetUserImagesRequest(db *gorm.DB, params user_images_v1.GetUserImagesV1Params) (resp *models.UserImagesResponse, errResp *models.ErrResponse) {
	_, errResp = commonutils.ExtractJWT(params.HTTPRequest)
	if errResp != nil {
		return
	}
	return
}

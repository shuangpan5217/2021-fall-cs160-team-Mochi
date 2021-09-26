package login

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/login_v1"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func LoginV1Handler(db *gorm.DB) login_v1.LoginV1HandlerFunc {
	return func(params login_v1.LoginV1Params) (responder middleware.Responder) {
		username := *params.Body.Username
		// password := *params.Body.Password

		resp := login_v1.NewLoginV1OK()

		respStruct := models.LoginResponse{
			Username: username,
		}
		resp.SetPayload(&respStruct)
		return resp
	}
}

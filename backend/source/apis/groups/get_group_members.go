package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func GetGroupMembersHandler(db *gorm.DB) groups_v1.GetGroupUsersV1HandlerFunc {
	return func(params groups_v1.GetGroupUsersV1Params) middleware.Responder {
		resp := groups_v1.NewGetGroupUsersV1OK()
		resp.SetPayload(nil)
		return resp
	}
}

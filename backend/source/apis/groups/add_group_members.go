package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func AddGroupMembersHandler(db *gorm.DB) groups_v1.AddGroupUsersV1HandlerFunc {
	return func(params groups_v1.AddGroupUsersV1Params) middleware.Responder {
		resp := groups_v1.NewAddGroupUsersV1OK()
		resp.SetPayload(nil)
		return resp
	}
}

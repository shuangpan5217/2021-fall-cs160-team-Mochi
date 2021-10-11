package groups

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

func RemoveGroupMembersHandler(db *gorm.DB) groups_v1.RemoveGroupUsersV1HandlerFunc {
	return func(params groups_v1.RemoveGroupUsersV1Params) middleware.Responder {
		resp := groups_v1.NewRemoveGroupUsersV1OK()
		resp.SetPayload(nil)
		return resp
	}
}

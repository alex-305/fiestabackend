package auth

import (
	"github.com/alex-305/fiestabackend/db"
	"github.com/alex-305/fiestabackend/models"
)

func GetPermissions(username, fiestaid string, db *db.DB) models.UserPermissions {
	var perms models.UserPermissions

	perms.IsOwner = db.IsOwner(username, fiestaid)
	perms.CanPost = db.HasPermission(username, fiestaid)

	return perms
}

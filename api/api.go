package api

import (
	"github.com/berkansasmaz/kartaca-internship-application.git/models"
)

// API -
type API struct {
	users *models.UserManager
	posts *models.PostManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)
	postmgr, _ := models.NewPostManager(db)

	return &API{
		users: usermgr,
		posts: postmgr,
	}
}

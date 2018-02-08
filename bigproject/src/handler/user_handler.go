package handler

import (
	"fmt"
	"net/http"

	"github.com/tokopedia/gosample/bigproject/src/database"
	"github.com/tokopedia/gosample/bigproject/src/utils"
)

var limit int

func init() {
	limit = 10
}

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	user, err := database.GetUser(name, limit)
	if err != nil {
		fmt.Printf("%s", err)
		utils.SendAjaxResponseError(w, err.Error())
		return
	}

	utils.SendAjaxResponseSuccess(w, user)
}

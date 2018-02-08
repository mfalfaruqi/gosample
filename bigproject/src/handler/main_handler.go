package handler

import (
	"html/template"
	"net/http"

	"github.com/tokopedia/gosample/bigproject/src/database"
	"github.com/tokopedia/gosample/bigproject/src/utils"
)

// Home for home url handler
func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("bigproject/files/templates/index.html")
	t.Execute(w, "")
}

// GetVisitorCount to get visitor count
func GetVisitorCount(w http.ResponseWriter, r *http.Request) {
	visitor, err := database.IncVisitorCount()

	if err != nil {
		utils.SendAjaxResponseError(w, err.Error())
	}

	utils.SendAjaxResponseSuccess(w, visitor)
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scriptnsam/url-shortner/pkg/config"
	"github.com/scriptnsam/url-shortner/pkg/models"
	"github.com/scriptnsam/url-shortner/pkg/utils"
)

func GetUrls(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	u := models.GetAll(db)
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	db := config.GetDB()

	urlDetails, _ := models.GetUrlById(db, uid)

	models.UpdateUrl(db, uid, urlDetails)

	w.Header().Set("Content-Type", "application/json")
	http.Redirect(w, r, urlDetails.Redirect_Url, http.StatusFound)
	return
}

func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	CreateUrl := &models.UrlDetails{}
	utils.ParseBody(r, CreateUrl)
	db := config.GetDB()
	u := CreateUrl.CreateUrl(db)
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

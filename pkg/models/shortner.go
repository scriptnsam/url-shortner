package models

import (
	"github.com/jinzhu/gorm"
	"github.com/scriptnsam/url-shortner/pkg/config"
)

type UrlDetails struct {
	gorm.Model
	U_id         string `json:"u_id"`
	Redirect_Url string `json:"redirect_url"`
	View_Count   int64  `json:"view_count"`
}

func init() {
	db := config.Connect()
	db.AutoMigrate(&UrlDetails{})
}

func GetAll(db *gorm.DB) []UrlDetails {
	var urls []UrlDetails
	db.Find(&urls)
	return urls
}

func GetUrlById(db *gorm.DB, Id string) (*UrlDetails, *gorm.DB) {
	if db == nil {
		return nil, db
	}
	var getUrl UrlDetails
	d := db.Where("U_id=?", Id).Find(&getUrl)

	return &getUrl, d
}

func (u *UrlDetails) CreateUrl(db *gorm.DB) *UrlDetails {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func UpdateUrl(db *gorm.DB, UID string, u *UrlDetails) *gorm.DB {
	var updateUrl UrlDetails
	d := db.Where("U_id=?", UID).Find(&updateUrl)
	updateUrl.Redirect_Url = u.Redirect_Url
	updateUrl.View_Count++
	d.Save(&updateUrl)
	return d
}

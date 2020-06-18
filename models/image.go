package models

import "github.com/jinzhu/gorm"


// Image model
type Image struct {
	URL string `json:"url" gorm:"type:varchar(50)"`
	UserID string `json:"url" gorm:"type:varchar(10)"`
	DeleteHash string `json:"delete_has"`
	ImgurID string `json:"imgur_id" gorm:"type:varchar(10)"`
}

// IMGURResp ...
type IMGURResp struct {
	data ImgData
}

// ImgData ...
type ImgData struct {
	id string
	deletehash string
	link string
}
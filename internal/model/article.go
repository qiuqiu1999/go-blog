package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_tag"
}

func (a *Article) Create(db *gorm.DB) error {
	return db.Create(a).Error
}

func (a *Article) Update(db *gorm.DB) error {
	return db.Model(&Article{}).Where("id = ? AND is_del = ?", a.ID, 0).Update(a).Error
}

func (a *Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(a).Error
}
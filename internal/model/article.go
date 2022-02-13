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

func (a *Article) Get(db *gorm.DB) (*Article, error) {
	var article *Article
	var err error

	db = db.Where("id = ?", a.ID)
	if err = db.Where("state = ?", a.State).Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (a *Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("name = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
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

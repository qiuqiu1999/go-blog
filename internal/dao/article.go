package dao

import (
	"github.com/qiuqiu1999/go-blog/internal/model"
	"github.com/qiuqiu1999/go-blog/pkg/app"
)

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetArticle(id uint32) (*model.Article, error) {
	article := model.Article{}
	return article.Get(d.engine)
}

func (d *Dao) CreateArticle(title string, desc string, content string, coverImageUrl string, state uint8, CreatedBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{CreatedBy: CreatedBy},
	}

	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, desc string, content string, coverImageUrl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	return article.Update(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

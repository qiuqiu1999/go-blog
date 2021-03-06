package service

import (
	"github.com/qiuqiu1999/go-blog/internal/model"
	"github.com/qiuqiu1999/go-blog/pkg/app"
)

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"max:100"`
	Content       string `form:"content" binding:"max:100"`
	CoverImageUrl string `form:"cover_image_url" binding:"max:100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"max:100"`
	Content       string `form:"content" binding:"max:100"`
	CoverImageUrl string `form:"cover_image_url" binding:"max:100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) Get(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}
func (svc *Service) List(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}
func (svc *Service) Create(param CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.CreatedBy)
}
func (svc *Service) Update(param UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.ModifiedBy)
}
func (svc *Service) Delete(param DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}

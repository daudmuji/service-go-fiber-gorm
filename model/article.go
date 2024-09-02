package model

import (
	"github.com/gofiber/fiber/v2"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"gorm.io/plugin/soft_delete"
)

type Article struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"<-:create"`
	Content   *string
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

type ArticleRepository interface {
	CreateArticle(article Article) (error, int64)
	GetByTitle(title string) (error, []Article)
	UpdateArticleByTitle(title string, article Article) (error, int64)
	DeleteArticleByTitle(title string) (error, int64)
}

type ArticleService interface {
	CreateArticle(ctx *fiber.Ctx, articleRequest request.CreateArticle) (error, response.SuccessResponse)
	GetByTitle(ctx *fiber.Ctx, title string) (error, response.SuccessResponse)
	UpdateArticleByTitle(ctx *fiber.Ctx, title string, articleUpdate request.UpdateArticle) (error, response.SuccessResponse)
	DeleteArticleByTitle(ctx *fiber.Ctx, title string) (error, response.SuccessResponse)
}

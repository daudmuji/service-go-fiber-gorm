package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-template-service/model"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"log"
	"time"
)

type ArticleUsecase struct {
	ArticleRepo model.ArticleRepository
}

func NewArticleUsecase(ArticleRepo model.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{ArticleRepo: ArticleRepo}
}

func (a *ArticleUsecase) CreateArticle(ctx *fiber.Ctx, articleRequest request.CreateArticle) (err error, res response.SuccessResponse) {

	jsonByte, err := json.Marshal(articleRequest.Content)
	if err != nil {
		e := fmt.Sprintf("Error Json : %s, content: %s | Error: %s", articleRequest.Title, articleRequest.Content, err.Error())
		log.Println(e)
		return err, res
	}

	stringContent := string(jsonByte)

	article := model.Article{
		Title:   articleRequest.Title,
		Content: &stringContent,
	}

	err, _ = a.ArticleRepo.CreateArticle(article)
	if err != nil {
		e := fmt.Sprintf("Error inserting article with title: %s, content: %s | Error: %s", articleRequest.Title, articleRequest.Content, err.Error())
		log.Println(e)
		return err, res
	}

	res = response.SuccessResponse{
		Message:   "Success Store",
		TimeStamp: time.Now(),
		Data:      articleRequest,
	}

	return nil, res
}

func (a *ArticleUsecase) GetByTitle(ctx *fiber.Ctx, title string) (err error, res response.SuccessResponse) {

	err, articles := a.ArticleRepo.GetByTitle(title)
	if err != nil {
		e := fmt.Sprintf("Error getting articles with title: %s, Error: %s", title, err.Error())
		log.Println(e)
		return err, res
	}

	return nil, response.SuccessResponse{
		Message:   "Success Get",
		TimeStamp: time.Now(),
		Data:      articles,
	}
}

func (a *ArticleUsecase) UpdateArticleByTitle(ctx *fiber.Ctx, title string, articleUpdate request.UpdateArticle) (err error, res response.SuccessResponse) {

	jsonByte, err := json.Marshal(articleUpdate.Content)
	if err != nil {
		e := fmt.Sprintf("Error Marshal Title : %s | Error: %s", title, err.Error())
		log.Println(e)
		return err, res
	}

	stringContent := string(jsonByte)

	article := model.Article{
		Content: &stringContent,
	}

	res = response.SuccessResponse{
		Message:   "Success Update",
		TimeStamp: time.Now(),
		Data:      articleUpdate,
	}

	err, _ = a.ArticleRepo.UpdateArticleByTitle(title, article)
	if err != nil {
		e := fmt.Sprintf("Error inserting article with title: %s, content: %s | Error: %s", title, articleUpdate.Content, err.Error())
		log.Println(e)
		return err, res
	}

	return nil, res
}

func (a *ArticleUsecase) DeleteArticleByTitle(ctx *fiber.Ctx, title string) (err error, res response.SuccessResponse) {
	err, _ = a.ArticleRepo.DeleteArticleByTitle(title)
	if err != nil {
		e := fmt.Sprintf("Error deleting article with title: %s, Error: %s", title, err.Error())
		log.Println(e)
		return err, res
	}

	res = response.SuccessResponse{
		Message:   "Success Delete",
		TimeStamp: time.Now(),
		Data:      nil,
	}

	return nil, res
}

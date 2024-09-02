package repository

import (
	"golang-template-service/model"
	"gorm.io/gorm"
)

type ArticleRepositoryPostgres struct {
	Conn *gorm.DB
}

func NewArticleRepositoryPostgres(Conn *gorm.DB) *ArticleRepositoryPostgres {
	return &ArticleRepositoryPostgres{Conn}
}

func (a *ArticleRepositoryPostgres) CreateArticle(article model.Article) (err error, result int64) {
	tx := a.Conn.Create(&article)
	err = tx.Error

	if err != nil {
		return err, result
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return err, result
	}

	return nil, rowsAffected
}

func (a *ArticleRepositoryPostgres) GetByTitle(title string) (err error, articles []model.Article) {

	rows, err := a.Conn.Table("article").Where("title = ?", title).Where("is_deleted", 0).Rows()
	if err != nil {
		return err, nil
	}

	defer rows.Close()

	listArticles := make([]model.Article, 0)
	for rows.Next() {
		var article model.Article
		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.IsDeleted,
		)
		if err != nil {
			return err, nil
		}
		listArticles = append(listArticles, article)
	}

	return nil, listArticles
}

func (a *ArticleRepositoryPostgres) UpdateArticleByTitle(title string, article model.Article) (err error, result int64) {

	tx := a.Conn.Model(&article).Where("title = ?", title).Updates(article)
	err = tx.Error
	if err != nil {
		return err, result
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return err, result
	}

	return nil, result
}

func (a ArticleRepositoryPostgres) DeleteArticleByTitle(title string) (err error, result int64) {
	tx := a.Conn.Table("article").Where("title = ?", title).Update("is_deleted", 1)

	err = tx.Error
	if err != nil {
		return err, result
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return err, result
	}

	return nil, result
}

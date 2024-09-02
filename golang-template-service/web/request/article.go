package request

type CreateArticle struct {
	Title   string `json:"title" validate:"required,min=5"`
	Content any    `json:"content" validate:"required"`
}

type UpdateArticle struct {
	Content any `json:"content" validate:"required"`
}

type GetArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

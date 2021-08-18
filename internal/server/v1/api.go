package v1

import (
	"net/http"

	"go-restful/internal/data"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	articles := &ArticleRouter{
		Repository: &data.ArticleRepository{
			Data: data.New(),
		},
	}

	r.Mount("/articles", articles.Routes())

	return r
}

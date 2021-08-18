package v1

import (
	"encoding/json"
	"fmt"
	"go-restful/models/article"
	"net/http"
	"strconv"

	"go-restful/models/response"

	"github.com/go-chi/chi"
)

type ArticleRouter struct {
	Repository article.Repository
}

func (art *ArticleRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", art.GetAllHandler)
	r.Get("/create", art.CreateHandler)
	r.Get("/{id}", art.GetOneHandler)
	r.Get("/{id}/update", art.UpdateHandler)
	r.Get("/{id}/delete", art.DeleteHandler)
	return r
}

func (art *ArticleRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var newArticle article.Article
	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = art.Repository.Create(ctx, &newArticle)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), newArticle.Id))
	response.JSON(w, r, http.StatusCreated, response.Map{"article": newArticle})
}

func (art *ArticleRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	articles, err := art.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"articles": articles})
}

func (art *ArticleRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	u, err := art.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"article": u})
}

func (art *ArticleRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var updateArticle article.Article
	err = json.NewDecoder(r.Body).Decode(&updateArticle)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = art.Repository.Update(ctx, uint(id), updateArticle)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (art *ArticleRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = art.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

package data

import (
	"context"
	"go-restful/models/article"
)

type ArticleRepository struct {
	Data *Data
}

func (arts *ArticleRepository) GetAll(ctx context.Context) ([]article.Article, error) {
	q := `SELECT id, title, description, content
        	FROM articles; 
	`

	rows, err := arts.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []article.Article
	for rows.Next() {
		var article article.Article
		rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		articles = append(articles, article)
	}

	return articles, nil
}

func (art *ArticleRepository) GetOne(ctx context.Context, id uint) (article.Article, error) {
	q := `
    SELECT id, title, description, content
        FROM articles WHERE id = $1;
    `

	row := art.Data.DB.QueryRowContext(ctx, q, id)

	var newArt article.Article
	err := row.Scan(&newArt.Id, &newArt.Title, &newArt.Desc, &newArt.Content)
	if err != nil {
		return article.Article{}, err
	}

	return newArt, nil
}

func (art *ArticleRepository) Create(ctx context.Context, article *article.Article) error {
	q := `
    INSERT INTO articles (title, description, content)
        VALUES ($1, $2, $3)
        RETURNING id;
    `
	row := art.Data.DB.QueryRowContext(
		ctx, q, article.Title, article.Desc, article.Content,
	)

	err := row.Scan(&article.Id)
	if err != nil {
		return err
	}

	return nil
}

func (art *ArticleRepository) Update(ctx context.Context, id uint, article article.Article) error {
	q := `
    UPDATE articles set title=$1, description=$2, content=$3
        WHERE id=$4;
    `

	stmt, err := art.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, article.Title, article.Desc, article.Content, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (art *ArticleRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM articles WHERE id=$1;`

	stmt, err := art.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

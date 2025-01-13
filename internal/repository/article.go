package repository

import (
	"database/sql"
	"github.com/nickloujia/mycrm/internal/models"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAll() ([]models.Article, error) {
	articles := []models.Article{}
	rows, err := r.db.Query("SELECT id, title, content, created_at, updated_at FROM articles ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

// GetByID 通过ID获取文章
func (r *ArticleRepository) GetByID(id int) (*models.Article, error) {
	var article models.Article
	err := r.db.QueryRow("SELECT id, title, content, created_at, updated_at FROM articles WHERE id = ?", id).
		Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// Update 更新文章
func (r *ArticleRepository) Update(article models.Article) error {
	_, err := r.db.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ?",
		article.Title, article.Content, article.ID)
	return err
}

// Delete 删除文章
func (r *ArticleRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}

// 其他 CRUD 方法... 
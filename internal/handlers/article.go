package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/nickloujia/mycrm/internal/models"
    "github.com/nickloujia/mycrm/internal/repository"
    "github.com/gorilla/mux"
)

type ArticleHandler struct {
    repo *repository.ArticleRepository
}

func NewArticleHandler(repo *repository.ArticleRepository) *ArticleHandler {
    return &ArticleHandler{repo: repo}
}

func (h *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
    articles, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(articles)
}

// GetArticle 获取单个文章
func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "无效的文章ID", http.StatusBadRequest)
        return
    }

    article, err := h.repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(article)
}

// UpdateArticle 更新文章
func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "无效的文章ID", http.StatusBadRequest)
        return
    }

    var article models.Article
    if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    article.ID = id

    if err := h.repo.Update(article); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(article)
}

// DeleteArticle 删除文章
func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "无效的文章ID", http.StatusBadRequest)
        return
    }

    if err := h.repo.Delete(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// 其他处理方法...
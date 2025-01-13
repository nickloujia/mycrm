package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/nickloujia/mycrm/internal/repository"
)

type CustomerHandler struct {
    repo *repository.CustomerRepository
}

func NewCustomerHandler(repo *repository.CustomerRepository) *CustomerHandler {
    return &CustomerHandler{repo: repo}
}

func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
    customers, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(customers)
}

// 其他处理方法... 
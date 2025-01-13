package main

import (
    "log"
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/nickloujia/mycrm/internal/database"
    "github.com/nickloujia/mycrm/internal/handlers"
    "github.com/nickloujia/mycrm/internal/repository"
)

func main() {
    // 初始化数据库连接
    db, err := database.NewDB("user:password@tcp(127.0.0.1:3306)/crm_db?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 初始化存储层
    articleRepo := repository.NewArticleRepository(db)

    // 初始化处理器
    articleHandler := handlers.NewArticleHandler(articleRepo)

    // 创建路由
    router := mux.NewRouter()

    // API 路由
    api := router.PathPrefix("/api").Subrouter()
    
    // 文章相关路由
    api.HandleFunc("/articles", articleHandler.GetArticles).Methods("GET")
    api.HandleFunc("/articles/{id}", articleHandler.GetArticle).Methods("GET")
    api.HandleFunc("/articles/{id}", articleHandler.UpdateArticle).Methods("PUT")
    api.HandleFunc("/articles/{id}", articleHandler.DeleteArticle).Methods("DELETE")

    // 静态文件服务
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("templates")))

    // 启动服务器
    log.Println("服务器运行在 :8080 端口...")
    log.Fatal(http.ListenAndServe(":8080", router))
} 
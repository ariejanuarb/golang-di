package main

import (
	"github.com/go-playground/validator/v10"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	// memanggil constructor service category
	categoryRepository := repository.NewCategoryRepository()

	// dependency : categoryReposotory, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// membuat object categoryConytoller, dengan memanggil constructor NewCategoryController
	// dependency : categoryService
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

package api

import (
	"flag"
	"fmt"
	"hardware-store/api/controllers"
	"hardware-store/api/database"
	"hardware-store/api/models"
	"hardware-store/api/repository"
	"hardware-store/api/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	port        = flag.Int("p", 5000, "set port")
	resetTables = flag.Bool("rt", false, "reset tables")
)

func Run() {

	flag.Parse()

	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	fmt.Println("Database connected...")

	if *port != 5000 && *resetTables {
		createSuperTestTables()
	}

	categoriesRepository := repository.NewCategoriesRepository(db)
	productsRepository := repository.NewProductsRepository(db)

	categoriesController := controllers.NewCategoriesRepository(categoriesRepository)
	productsController := controllers.NewProductsController(productsRepository)

	categoryRoutes := routes.NewCategoryRoutes(categoriesController)
	productRoutes := routes.NewProductRoutes(productsController)

	router := mux.NewRouter().StrictSlash(true)

	routes.Install(router, categoryRoutes, productRoutes)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location", "Entity", "Accept"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("API Listening on", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}

func createSuperTestTables() {
	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	tx := db.Begin()
	err := tx.Debug().DropTableIfExists(&models.Product{}, &models.Category{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().CreateTable(&models.Category{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().CreateTable(&models.Product{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().Model(&models.Product{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}

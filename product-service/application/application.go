package application

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/bernardn38/ebuy/product-service/handler"
	"github.com/bernardn38/ebuy/product-service/service"
	"github.com/bernardn38/ebuy/product-service/sql/products"
	"github.com/bernardn38/ebuy/product-service/token"
	"github.com/cristalhq/jwt/v4"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type config struct {
	jwtSecretKey     string
	jwtSigningMethod jwt.Algorithm
	postgresDsn      string
	port             string
}
type server struct {
	router *chi.Mux
	port   string
}
type App struct {
	srv *server
}

func New() *App {
	app := App{}
	config, err := getEnvConfig()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	app.runAppSetup(config)
	return &app
}

func (app *App) Run() {
	//start server
	log.Printf("listening on port %s", app.srv.port)
	log.Fatal(http.ListenAndServe(app.srv.port, app.srv.router))
}

func (app *App) runAppSetup(c config) {
	//open connection to postgres
	db, err := sql.Open("postgres", c.postgresDsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	// init sqlc user queries
	queries := products.New(db)

	//create jwt manager
	tm := token.NewManager([]byte(c.jwtSecretKey), c.jwtSigningMethod)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "password",
		DB:       0,
	})
	productService := service.NewProductService(queries, db, rdb)
	go productService.RunAsync()
	// init request handler
	h := handler.NewHandler(productService, tm)

	app.srv = &server{
		router: setupRouter(h),
		port:   c.port,
	}
}

func setupRouter(h *handler.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/api/v1/products/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Server is up and running"))
	})
	router.Post("/api/v1/products", h.CreateProduct)
	router.Get("/api/v1/products", h.GetAllProducts)
	router.Get("/api/v1/products/recent", h.GetRecentUploadedProducts)
	return router
}

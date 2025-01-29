package app

import (
	"awesomeProjectSamb/internal/database"
	"awesomeProjectSamb/internal/handler"
	"awesomeProjectSamb/pkg/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.elastic.co/apm/module/apmchi/v2"
)

func newRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(chiMiddleware.RequestID)

	router.Use(middleware.Recover())

	router.Use(apmchi.Middleware())

	return router

}

func router() *chi.Mux {
	router := newRouter()

	router.Get("/ping", handler.Ping)

	dbConnection := database.DBConnection

	mysqlDB := dbConnection[database.MysqlDB]

	handlerR := handler.NewHandler(mysqlDB)

	router.Get("/health-check", handlerR.HealthCheck)

	router.Group(func(s chi.Router) {
		s.Route("/api/v1", func(rApi chi.Router) {

			rApi.Post("/incoming", handlerR.IncomingGoods)
			rApi.Post("/outgoing", handlerR.OutgoingGoods)
			rApi.Get("/stock", handlerR.StockReport)
		})
	})

	return router

}

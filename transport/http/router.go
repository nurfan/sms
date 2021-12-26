package http

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nurfan/sms/repository"
	adp "github.com/nurfan/sms/transport/http/adapter"
)

func Serve(conn *repository.RepositoryPsql) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	adapter := adp.NewAdapter(conn)

	// auth group
	u := e.Group("/api/v1/auth")
	{
		u.GET("/token", adapter.GetToken)
	}

	a := e.Group("/api/v1")
	{
		config := middleware.JWTConfig{
			SigningKey: []byte("secret"),
		}
		a.Use(middleware.JWTWithConfig(config))

		a.GET("/findings", adapter.GetFindings)
		a.POST("/findings", adapter.CreateFindings)
		a.POST("/upload", adapter.Upload)
	}

	var appPort = ":8000"

	if os.Getenv("APP_PORT") != "" {
		appPort = ":" + os.Getenv("APP_PORT")
	}

	// Start server
	e.Logger.Fatal(e.Start(appPort))
}

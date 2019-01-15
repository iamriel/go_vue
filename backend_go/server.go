package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Email string `json:"email"`
}

type StatsCard struct {
	Type    string    `json:"type"`
	Icon    string    `json:"icon"`
	Title   string    `json:"title"`
	Value   int       `json:"value"`
	Unit    string    `json:"unit"`
	Updated time.Time `json:"updated"`
}

func login(c echo.Context) error {
	const email = "test@email.com"
	const password = "password"

	if email != strings.ToLower(c.FormValue("email")) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Email does not exist.")
	}
	if password != c.FormValue("password") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password incorrect.")
	}
	u := &User{
		Email: email,
	}
	return c.JSON(http.StatusOK, u)
}

func cards(c echo.Context) error {
	currentDateTime := time.Now()

	data := []StatsCard{
		StatsCard{
			Type:    "warning",
			Icon:    "ti-server",
			Title:   "Capacity",
			Value:   rand.Intn(500),
			Unit:    "GB",
			Updated: currentDateTime,
		},
		StatsCard{
			Type:    "success",
			Icon:    "ti-wallet",
			Title:   "Revenue",
			Value:   rand.Intn(500),
			Unit:    "USD",
			Updated: currentDateTime,
		},
		StatsCard{
			Type:    "danger",
			Icon:    "ti-pulse",
			Title:   "Errors",
			Value:   rand.Intn(100),
			Updated: currentDateTime,
		},
		StatsCard{
			Type:    "info",
			Icon:    "ti-twitter-alt",
			Title:   "Followers",
			Value:   rand.Intn(200),
			Updated: currentDateTime,
		},
	}
	return c.JSON(http.StatusOK, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// Allows requests from `http://localhost:8080`
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/login", login)
	e.GET("/statistics/cards", cards)
	e.Logger.Fatal(e.Start(":1323"))
}

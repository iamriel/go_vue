package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/leekchan/timeutil"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StatsCard struct {
	Type    string    `json:"type"`
	Icon    string    `json:"icon"`
	Title   string    `json:"title"`
	Value   int       `json:"value"`
	Unit    string    `json:"unit"`
	Updated time.Time `json:"updated"`
}

func randomNumber(min int, max int) int {
	max = max - min
	return rand.Intn(max) + min
}

func login(c echo.Context) (err error) {
	const email = "test@email.com"
	const password = "password"

	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	data := []Error{}

	if email != strings.ToLower(u.Email) {
		data = append(data, Error{
			Field:   "email",
			Message: "Email does not exist",
		})
		return echo.NewHTTPError(http.StatusUnauthorized, data)
	}
	if password != u.Password {
		data = append(data, Error{
			Field:   "password",
			Message: "Password incorrect.",
		})
		return echo.NewHTTPError(http.StatusUnauthorized, data)
	}
	user := &User{
		Email: email,
	}
	return c.JSON(http.StatusOK, user)
}

func randomUpdatedValue() time.Time {
	base := time.Now()
	td := timeutil.Timedelta{Minutes: time.Duration(-rand.Intn(60))}
	return base.Add(td.Duration())
}

func cards(c echo.Context) error {
	data := []StatsCard{
		StatsCard{
			Type:    "warning",
			Icon:    "ti-server",
			Title:   "Capacity",
			Value:   rand.Intn(500),
			Unit:    "GB",
			Updated: randomUpdatedValue(),
		},
		StatsCard{
			Type:    "success",
			Icon:    "ti-wallet",
			Title:   "Revenue",
			Value:   rand.Intn(500),
			Unit:    "USD",
			Updated: randomUpdatedValue(),
		},
		StatsCard{
			Type:    "danger",
			Icon:    "ti-pulse",
			Title:   "Errors",
			Value:   rand.Intn(100),
			Updated: randomUpdatedValue(),
		},
		StatsCard{
			Type:    "info",
			Icon:    "ti-twitter-alt",
			Title:   "Followers",
			Value:   rand.Intn(200),
			Updated: randomUpdatedValue(),
		},
	}
	return c.JSON(http.StatusOK, data)
}

func usersBehavior(c echo.Context) error {
	type UsersBehavior struct {
		Labels []string `json:"labels"`
		Series [][]int  `json:"series"`
	}
	data := UsersBehavior{
		Labels: []string{
			"9:00AM",
			"12:00AM",
			"3:00PM",
			"6:00PM",
			"9:00PM",
			"12:00PM",
			"3:00AM",
			"6:00AM",
		},
		Series: [][]int{},
	}
	user1 := []int{}
	user2 := []int{}
	user3 := []int{}
	for range [8]int{} {
		user1 = append(user1, rand.Intn(900))
		user2 = append(user2, rand.Intn(900))
		user3 = append(user3, rand.Intn(900))
	}
	data.Series = append(data.Series, user1)
	data.Series = append(data.Series, user2)
	data.Series = append(data.Series, user3)
	return c.JSON(http.StatusOK, data)
}

func emailStatistics(c echo.Context) error {
	type EmailStats struct {
		Labels []string  `json:"labels"`
		Series []float32 `json:"series"`
	}
	open := float32(randomNumber(100, 200))
	bounce := float32(randomNumber(100, 200))
	unsubscribe := float32(randomNumber(100, 200))
	total := open + bounce + unsubscribe

	openPercent := open / total * 100
	bouncePercent := bounce / total * 100
	unsubscribePercent := unsubscribe / total * 100

	data := EmailStats{
		Labels: []string{
			fmt.Sprintf("%f%%", openPercent),
			fmt.Sprintf("%f%%", bouncePercent),
			fmt.Sprintf("%f%%", unsubscribePercent),
		},
		Series: []float32{
			openPercent,
			bouncePercent,
			unsubscribePercent,
		},
	}
	return c.JSON(http.StatusOK, data)
}

func salesData(c echo.Context) error {
	type SalesData struct {
		Labels []string `json:"labels"`
		Series [][]int  `json:"series"`
	}
	data := SalesData{
		Labels: []string{
			"Jan",
			"Feb",
			"Mar",
			"Apr",
			"Mai",
			"Jun",
			"Jul",
			"Aug",
			"Sep",
			"Oct",
			"Nov",
			"Dec",
		},
		Series: [][]int{},
	}
	product1 := []int{}
	product2 := []int{}
	for range [12]int{} {
		product1 = append(product1, randomNumber(200, 900))
		product2 = append(product2, randomNumber(200, 900))
	}
	data.Series = append(data.Series, product1)
	data.Series = append(data.Series, product2)
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
	e.GET("/statistics/users-behavior", usersBehavior)
	e.GET("/statistics/email", emailStatistics)
	e.GET("/statistics/sales", salesData)
	e.Logger.Fatal(e.Start(":1323"))
}

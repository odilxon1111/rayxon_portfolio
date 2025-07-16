package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	ID        int64
	FirstName string
	LastName  string
	Phone     uint64
}

func main() {

	dsn := "host=localhost user=admin password=admin dbname=test0 port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dbErr := db.AutoMigrate(&Student{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.GET("/add", func(c echo.Context) error {
		return c.File("add.html")
	})

	e.POST("/save", func(c echo.Context) error {

		var student Student
		student.FirstName = c.FormValue("FirstName")
		student.LastName = c.FormValue("LastName")
		var phone string = c.FormValue("Phone")
		student.Phone, _ = strconv.ParseUint(phone, 10, 64)

		db.Create(&student)

		return c.Redirect(http.StatusMovedPermanently, "/")
	})

	_ = e.Start(":8186")
}

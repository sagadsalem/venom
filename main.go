package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"projects/starter/router"
)

func init() {
	// load the .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
func main() {
	// init the router
	r := router.New()
	// running the server
	port := os.Getenv("APP_PORT")
	logrus.Info("server is now running http://localhost:" + port)
	panic(http.ListenAndServe(":"+port, r))
}

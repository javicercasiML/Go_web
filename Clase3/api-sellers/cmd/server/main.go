package main

import (
	"database/sql"

	"grabaciones/functional_test/cmd/server/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	db, _ := sql.Open("mysql", "root:root@/melisprint")
	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/config"
	"github.com/LucasSaraiva019/devbook-arango/src/database"
	"github.com/LucasSaraiva019/devbook-arango/src/models"
	"github.com/LucasSaraiva019/devbook-arango/src/router"
	arango "github.com/joselitofilho/gorm-arango/pkg"
	"gorm.io/gorm"
)

var err error

func main() {
	config.Initialize()

	database.DB, err = gorm.Open(arango.Open(database.BuildDBConfig()), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	database.DB.AutoMigrate(&models.User{})

	r := router.Generate()
	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

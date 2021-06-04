package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/config"
	"github.com/LucasSaraiva019/devbook-arango/src/migrations"
	"github.com/LucasSaraiva019/devbook-arango/src/router"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	migrations.AutoMigrate()
	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

package banco

import (
	"os"

	arango "github.com/joselitofilho/gorm-arango/pkg"
	"gorm.io/gorm"
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*gorm.DB, error) {
	arangodbConfig := &arango.Config{
		URI:                  "http://localhost:8529",
		User:                 os.Getenv("DB_USUARIO"),
		Password:             os.Getenv("DB_SENHA"),
		Database:             os.Getenv("DB_NOME"),
		Timeout:              120,
		MaxConnectionRetries: 10}
	db, err := gorm.Open(arango.Open(arangodbConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

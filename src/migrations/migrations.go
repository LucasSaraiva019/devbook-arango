package migrations

import (
	"github.com/LucasSaraiva019/devbook-arango/src/banco"
	"github.com/LucasSaraiva019/devbook-arango/src/modelos"
)

func AutoMigrate() (interface{}, error) {
	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}
	return db.AutoMigrate(&modelos.Usuario{}), nil
}

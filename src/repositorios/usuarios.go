package repositorios

import (
	"errors"

	"github.com/LucasSaraiva019/devbook-arango/src/modelos"
	"gorm.io/gorm"
)

type RepositorioDeUsuarios struct {
	db *gorm.DB
}

func NovoRepositorioDeUsuarios(db *gorm.DB) *RepositorioDeUsuarios {
	return &RepositorioDeUsuarios{db}
}

// Criar insere um usuário no banco de dados
func (repo *RepositorioDeUsuarios) Criar(usuario *modelos.Usuario) error {
	tx := repo.db.Create(usuario)
	if tx == nil {
		return errors.New("Erro ao criar o Usuario")
	}

	return tx.Error
}

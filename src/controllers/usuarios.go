package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/banco"
	"github.com/LucasSaraiva019/devbook-arango/src/modelos"
	"github.com/LucasSaraiva019/devbook-arango/src/repositorios"
	"github.com/LucasSaraiva019/devbook-arango/src/respostas"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	// TODO: Criar conexão na mais e passar onde necessário
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.Criar(&usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

package rotas

import (
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
	},
}

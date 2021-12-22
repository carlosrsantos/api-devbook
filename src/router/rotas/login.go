package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rotas{

	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}

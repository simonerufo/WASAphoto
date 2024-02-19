package api

import (
	"net/http"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*
		TODO:
			-get dei parametri della req e cast
			-controllare se l'user vuole bannarsi da solo
			-controllare se l'user che deve essere bannato esiste
			-auth
			-add to db
			-messaggio di successo

	*/
}

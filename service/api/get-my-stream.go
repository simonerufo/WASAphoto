package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//get dai params per l'auth + check
	user_id, err := strconv.Atoi(ps.ByName("Authorization"))
	if err != nil {
		http.Error(w, "Invalid auth token", http.StatusBadGateway)
	}

	isAuth(w, r, user_id)
	profile, err := rt.db.GetStream(user_id)
	if err != nil {
		http.Error(w, "Profile not found", http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
	fmt.Printf("Posts")
	//get da db dei post degli user che seguo e che non mi hanno bannato(? se mi hanno bannato mi si toglie il follow??)
	//faccio poi l'encode in json del file
}

package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*
		-get dei parametri di chi vuole uppare foto
		-auth
		-uppare su db foto (con caption)
		-messaggio di successo
	*/

	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user id from parameters", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)
	//POSSO UTILIZZARE MULTIPART FORM PER RACCOGLIERE FILE(IMMAGINE) E CAPTION
	err = r.ParseMultipartForm(16 << 20) //max 16 mb
	if err != nil {
		http.Error(w, "Error while parsing multipart form", http.StatusBadRequest)
		return
	}

	caption := r.FormValue("caption")

	image, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error while parsing the image", http.StatusBadRequest)
		return
	}
	imageData, err := io.ReadAll(image)
	if err != nil {
		http.Error(w, "Error while reading the image", http.StatusBadRequest)
		return
	}

	fileType := http.DetectContentType(imageData)
	if fileType != "image/jpeg" {
		http.Error(w, "wrong file type", http.StatusBadRequest)
		return
	}

	defer image.Close()

	err = rt.db.InsertPhoto(user_id, caption, imageData)
	if err != nil {
		http.Error(w, "Error while trying to insert post in database", http.StatusBadRequest)
	}

	fmt.Printf("photo successfully uploaded!\n")

}

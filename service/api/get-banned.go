package api

import(
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter")

func (rt *_router) GetBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Extract user_id from request parameters
		userID, err := strconv.Atoi(ps.ByName("user_id"))
		if err != nil {
			http.Error(w, "Error while parsing user id", http.StatusBadRequest)
			return
		}
	
		// Retrieve the ban list for the user from the database
		bannedUserIDs, err := rt.db.GetBanList(userID)
		if err != nil {
			http.Error(w, "Error while retrieving ban list from database", http.StatusInternalServerError)
			return
		}
	
		// Prepare the response data
		response := map[string][]int{
			"banned_user_ids": bannedUserIDs,
		}
	
		// Set the content type and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
}
	
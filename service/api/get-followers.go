package api
import(
	"net/http"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Retrieve user ID from parameters
    userID, err := strconv.Atoi(ps.ByName("user_id"))
    if err != nil {
        http.Error(w, "Error while fetching user ID from parameters", http.StatusBadRequest)
        return
    }

    // Retrieve the list of followers for the specified user
    followers, err := rt.db.GetFollowersForUser(userID)
    if err != nil {
        if strings.Contains(err.Error(), "does not exist") {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error while retrieving followers for user", http.StatusInternalServerError)
        }
        return
    }

    // Set response headers and write JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(followers)
}

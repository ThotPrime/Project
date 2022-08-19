package handlers 

import(
	"encoding/json"
	"net/http"

	//"github.com/tutorials/go/crud/pkg/mocks"
)

func GeAllbooks(w http.ResponseWriter, r *http.Request)
{
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOk)
	json.NewEncoder(w).Encode(mocks.Books)
}
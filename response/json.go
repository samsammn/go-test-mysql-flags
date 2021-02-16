package response

import (
	"encoding/json"
	"net/http"
)

type Map map[string]interface{}

func Json(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}

package server

import (
	"boston/learner"
	"encoding/json"
	"fmt"
	"net/http"
)

func handleListRequest(lmap *learner.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("List Learner Keys")
		keys := lmap.Keys()
		keysBytes, err := json.Marshal(keys)
		if err != nil {
			apiError := NewAPIError("List Learner Keys API Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		w.Write(keysBytes)
	}
}

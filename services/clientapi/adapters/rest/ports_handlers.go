package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"ndanamedtt/services/clientapi/application"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func HandlePortByID(find application.FindPortsProxyFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logErr := func(err error) {
			log.Println(fmt.Errorf("HandlePortByID: %w", err))
		}
		ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			//we certainly don't have that resource
			http.NotFound(w, r)
			return
		}
		p, err := find(r.Context(), ID)
		if err != nil {
			logErr(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if errors.Is(err, application.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		if err := json.NewEncoder(w).Encode(portToRest(p)); err != nil {
			//nothing to write into response here, json encoder probably already wrote
			// something and that will result in error of "already written response"
			logErr(err)
		}
	}
}

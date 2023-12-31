package history

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ChaosIsFramecode/horinezumi/data"
	"github.com/ChaosIsFramecode/horinezumi/jsonresp"
	"github.com/go-chi/chi/v5"
)

func SetupHistoryRoute(rt *chi.Mux, db data.Datastore) {
	// History subroute handler
	rt.Route("/h", func(histRouter chi.Router) {
		// Retrieve page history
		histRouter.Get("/{title}", func(w http.ResponseWriter, r *http.Request) {
			// Expect json response
			w.Header().Set("Content-Type", "application/json")

			titleParam := chi.URLParam(r, "title")
			if titleParam == "Main_Page" {
				// Main page
				w.Write([]byte("Nothing to see"))
			} else {
				// Redirect if not lowercase
				if strings.ToLower(titleParam) != titleParam {
					http.Redirect(w, r, strings.ToLower(titleParam), http.StatusSeeOther)
					return
				}

				// Fetch page history
				pH, err := db.FetchPageHistory(titleParam)
				if err != nil {
					jsonresp.JsonERR(w, 400, "Error with fetching page history: %s", err)
				}

				// Convert list to json
				json.NewEncoder(w).Encode(pH)
			}
		})
	})
}

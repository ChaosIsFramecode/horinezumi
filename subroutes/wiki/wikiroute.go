package wiki

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ChaosIsFramecode/horinezumi/data"
	"github.com/ChaosIsFramecode/horinezumi/subroutes/edit"
)

func SetupWikiroute(rt *chi.Mux, db data.Datastore) {
	// Add view subroute
	rt.Route("/wiki", func(wikirouter chi.Router) {
		// Redirect root to main page
		wikirouter.Get("/", http.RedirectHandler("/wiki/Main_Page", http.StatusSeeOther).ServeHTTP)

		// Page view handler
		wikirouter.Route("/{title}", func(pagerouter chi.Router) {
			// Retrieve page content
			pagerouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
				titleParam := chi.URLParam(r, "title")
				if titleParam == "Main_Page" {
					// Main page
					w.Write([]byte("Welcome to 堀ネズミ!"))
				} else {
					w.Write([]byte(titleParam))
				}
			})
		})
	})

	// Call edit subroute
	edit.SetupEditRoute(rt, db)
}
package routes

import (
	"net/http"

	"github.com/berkansasmaz/kartaca-internship-application.git/api"
	"github.com/berkansasmaz/kartaca-internship-application.git/auth"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	))

	// posts
	z := a.PathPrefix("/post").Subrouter()
	z.HandleFunc("/newpost", api.NewPost).Methods("POST")
	z.HandleFunc("/getposts", api.GetPosts).Methods("GET")

	return mux
}

package server

import (
	"net/http"
	readDir "trade-solar"
	"trade-solar/cmd/web/views"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	/* Static dirs routes */
	fs := readDir.ReadDirJs()
	mux.Handle("/js/", http.StripPrefix("/js/", fs))
	fs = readDir.ReadDirCss()
	mux.Handle("/css/", http.StripPrefix("/css/", fs))
	fs = readDir.ReadDirAssets()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	/* views routes */
	mux.Handle("/", templ.Handler(views.Home()))

	return mux
}

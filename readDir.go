package readDir

import "net/http"

func ReadDirJs() http.Handler {
	fs := http.FileServer(http.Dir("cmd/web/js"))
	return fs
}

func ReadDirCss() http.Handler {
	fs := http.FileServer(http.Dir("cmd/web/css"))
	return fs
}

func ReadDirAssets() http.Handler {
	fs := http.FileServer(http.Dir("cmd/web/assets"))
	return fs
}

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//go:embed ui
var embeddedFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("ui"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embeddedFiles, "ui")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()

	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	assetHandler := http.FileServer(getFileSystem(useOS))

	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	e.Logger.Fatal(e.Start(":1323"))
}

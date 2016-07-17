package main

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini"
)
import "github.com/ilucin/go-mdm/db"

type Projects struct {
	Items []db.Project
}

type Libs struct {
	Items []db.Lib
}

func main() {
	db.Init()

	m := martini.Classic()
	static := martini.Static("public/", martini.StaticOptions{Fallback: "/index.html"})
	m.NotFound(static, http.NotFound)

	m.Get("/projects", func(w http.ResponseWriter, r *http.Request) string {
		w.Header().Set("Content-Type", "application/json")
		projects := db.GetProjects()
		str, _ := json.Marshal(&Projects{Items: projects})
		return string(str)
	})

	m.Get("/libs", func(w http.ResponseWriter, r *http.Request) string {
		w.Header().Set("Content-Type", "application/json")
		libs := db.GetLibs()
		str, _ := json.Marshal(&Libs{Items: libs})
		return string(str)
	})

	m.Run()
}

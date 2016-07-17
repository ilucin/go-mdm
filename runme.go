package main

import (
	"github.com/ilucin/go-mdm/db"
	"github.com/ilucin/go-mdm/lib-manager"
	"github.com/ilucin/go-mdm/parser"
	"github.com/ilucin/go-mdm/repo"
)

func main() {
	db.Init()
	// db.TestPrepare()

	projectRepos := []string{"github.com/ilucin/plutus-app", "github.com/angular/angular", "github.com/facebook/react", "github.com/twbs/bootstrap", "github.com/rackt/redux", "github.com/postcss/postcss", "github.com/infinumjs/hektor-gulp", "github.com/infinumjs/hektor-grunt", "github.com/ilucin/plasticknives-web"}

	scanProjects(projectRepos)
	scanLibs()
}

func scanLibs() {
	libs := db.GetLibs()

	for i := 0; i < len(libs); i++ {
		scanLib(libs[i])
	}
}

func scanLib(lib db.Lib) {
	version, err := libManager.Scan(lib.Name, lib.Type)
	if err == nil {
		db.UpdateLibVersion(lib, version)
	}
}

func scanProjects(projectRepos []string) {
	if len(projectRepos) < 0 {
		return
	}

	for i := 0; i < len(projectRepos); i++ {
		scanProject(projectRepos[i])
	}
}

func scanProject(projectRepo string) {
	project := db.CreateOrGetProject(projectRepo)
	file, libType, _ := repo.Scan(projectRepo)
	libMap := parser.ParseStruct{}.Parse(libType, file)
	db.UpdateProjectLibs(project, libMap, libType)
}

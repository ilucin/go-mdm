package db

import (
	"fmt"

	"github.com/ilucin/go-mdm/enum"
)

// CreateOrGetProject creates new or returns existing project model
func CreateOrGetProject(projectRepo string) (project Project) {
	db.FirstOrCreate(&project, Project{Repo: projectRepo})
	return project
}

// UpdateOrCreateProjectLib updates or creates a project lib
func UpdateOrCreateProjectLib(project Project, libName string, libVersion string, libType enum.LibType) (projectLib ProjectLib) {
	var lib Lib
	db.FirstOrCreate(&lib, Lib{Name: libName, Type: libType})
	db.FirstOrCreate(&projectLib, ProjectLib{LibID: lib.ID, ProjectID: project.ID})

	if projectLib.Version != libVersion {
		projectLib.Version = libVersion
		db.Save(&projectLib)
	}

	return projectLib
}

// UpdateOrCreateLib updates or creates a lib
func UpdateOrCreateLib(libName string, libVersion string) (lib Lib) {
	db.FirstOrCreate(&lib, Lib{Name: libName})
	if lib.Version != libVersion {
		lib.Version = libVersion
		db.Save(&lib)
	}

	return lib
}

// UpdateProjectLibs updates a project lib data
func UpdateProjectLibs(project Project, libMap map[string]string, libType enum.LibType) {
	for libName := range libMap {
		fmt.Println(libName)
		UpdateOrCreateProjectLib(project, libName, libMap[libName], libType)
	}
}

// UpdateLibVersion does everything
func UpdateLibVersion(lib Lib, libVersion string) {
	lib.Version = libVersion
	db.Save(&lib)
}

// GetProjects gets all projects
func GetProjects() (projects []Project) {
	db.Find(&projects)
	for index := 0; index < len(projects); index++ {
		db.Model(&projects[index]).Related(&projects[index].Libs, "Libs")
		for j := 0; j < len(projects[index].Libs); j++ {
			db.Model(&projects[index].Libs[j]).Related(&projects[index].Libs[j].Lib, "Lib")
		}
	}

	fmt.Println(projects[0].Libs)
	return projects
}

// GetLibs gets all libs
func GetLibs() (libs []Lib) {
	db.Find(&libs)
	return libs
}

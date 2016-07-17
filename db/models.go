package db

import "github.com/ilucin/go-mdm/enum"

// Project struct
type Project struct {
	ID   int
	Repo string `sql:"not null;unique"`

	Libs []ProjectLib
}

func (p Project) String() string {
	return "string(p.ID)"
}

// ProjectLib struct
type ProjectLib struct {
	ID        int
	Version   string `sql:"not null"`
	Lib       Lib
	LibID     int
	Project   Project
	ProjectID int
}

// Lib is a library used in a project
type Lib struct {
	ID          int
	Name        string `sql:"not null;unique"`
	Version     string `sql:"not null"`
	Repository  string
	Type        enum.LibType
	ProjectLibs []ProjectLib
}

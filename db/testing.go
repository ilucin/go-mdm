package db

import "fmt"

// TestPrepare prepares
func TestPrepare() {
	Truncate()
}

// TestInsert tests if insert works properly
func TestInsert() {
	project := Project{Repo: "Test project"}
	lib := Lib{Name: "Test lib"}
	projectLib := ProjectLib{Version: "0.0.1", Project: project, Lib: lib}

	db.Create(&projectLib)

	fmt.Printf("New Project ID = %s\n", project.ID)
	fmt.Printf("New Lib ID = %s\n", lib.ID)
	fmt.Printf("New Project %d has %d Libs\n", project.ID, len(project.Libs))
}

// TestQuery tests querying
func TestQuery() {
	var projects []Project

	db.Find(&projects)

	fmt.Printf("Found %d projects\n", len(projects))
	fmt.Println(projects)

	var project Project

	db.First(&project)
	db.Model(&project).Related(&project.Libs, "Libs")

	fmt.Printf("Project %d\n", project.ID)
	fmt.Println(project)
	fmt.Printf("Project %d has %d Libs\n", project.ID, len(project.Libs))
}

package repo

import (
	"errors"
	"strings"

	"github.com/ilucin/go-mdm/enum"
)

// Scanner interface
type Scanner interface {
	Scan() ([]byte, enum.LibType)
}

// Repo struct
type Repo struct {
	repoType enum.RepoType
	url      string
}

// Scan scannes for a specific repo type defined on struct
func (r Repo) Scan() (file []byte, libType enum.LibType, err error) {
	if r.repoType == enum.GITHUB {
		file, libType, err = ScanGithub(r.url)
	} else {
		err = errors.New("Invalid input type")
	}

	return file, libType, err
}

// Scan scannes a repo and returns lib file content and a type
func Scan(url string) (file []byte, fileType enum.LibType, err error) {
	repoType, trimmedURL := parseURL(url)
	repo := Repo{repoType: repoType, url: trimmedURL}
	return repo.Scan()
}

func parseURL(url string) (repoType enum.RepoType, trimmedURL string) {
	trimmedURL = strings.Replace(strings.Replace(strings.Replace(url, "https://", "", -1), "http://", "", -1), "www", "", -1)

	if strings.Index(trimmedURL, GithubURL) == 0 {
		repoType = enum.GITHUB
	}

	return repoType, trimmedURL
}

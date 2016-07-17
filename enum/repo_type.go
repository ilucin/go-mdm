package enum

// RepoType is enum for remote repository types
type RepoType int

const (
	// GITHUB is enum for github
	GITHUB RepoType = 1 + iota
	// BITBUCKET is enum for bitbucket
	BITBUCKET
)

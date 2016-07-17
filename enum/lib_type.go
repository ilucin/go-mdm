package enum

// LibType is enum for dependency package types
type LibType int

const (
	// NPM is npm
	NPM LibType = 1 + iota
	// GEM is a gem
	GEM
)

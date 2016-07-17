package parser

import "github.com/ilucin/go-mdm/enum"

// Parser interface
type Parser interface {
	Parse() (enum.LibType, []byte)
}

// ParseStruct is struct that holds parsing data
type ParseStruct struct{}

//Parse parses lib type package file
func (p ParseStruct) Parse(libType enum.LibType, fileContent []byte) (versions map[string]string) {

	if libType == enum.NPM {
		versions = ParsePackageJSON(fileContent)
	}

	return versions
}

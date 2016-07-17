package parser

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ilucin/go-mdm/enum"
)

//Test tests
func Test() {
	filenames := [...]string{"parser/sample/package.json"}
	for fileI := 0; fileI < len(filenames); fileI++ {
		dat, err := ioutil.ReadFile(filenames[fileI])
		if err != nil {
			log.Fatal(err)
		}
		versions := ParseStruct{}.Parse(enum.NPM, dat)

		fmt.Print(versions)
	}
}

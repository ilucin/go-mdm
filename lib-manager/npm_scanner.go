package libManager

import (
	"bytes"
	"fmt"
	"os/exec"
)

// NpmScan scans npm
func NpmScan(libName string) (outputContent string, err error) {
	cmd := exec.Command("npm", "view", libName, "version")

	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut

	fmt.Printf("Starting command \"%s\" ... ", "npm info "+libName)
	err = cmd.Run()

	// if err != nil {
	// 	fmt.Printf("Command error")
	// 	return "{}"
	// }

	outputContent = cmdOut.String()
	fmt.Printf("%s", outputContent)

	return outputContent, err
}

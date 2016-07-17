package libManager

import "github.com/ilucin/go-mdm/enum"

// Scan scans packages to get current versions
func Scan(libName string, libType enum.LibType) (version string, err error) {
	if libType == enum.NPM {
		version, err = NpmScan(libName)
	}
	return version, err
}

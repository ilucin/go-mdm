package parser

import "github.com/antonholmquist/jason"

const removeSpecialFirstChars = true

// ParsePackageJSON parses package.json file
func ParsePackageJSON(rawFile []byte) map[string]string {
	v, _ := jason.NewObjectFromBytes(rawFile)

	devDependecies, _ := v.GetObject("devDependencies")
	deps, _ := v.GetObject("dependencies")

	versions := map[string]string{}
	assignVersionsFromDeps(devDependecies, versions)
	assignVersionsFromDeps(deps, versions)

	return versions
}

func assignVersionsFromDeps(deps *jason.Object, versions map[string]string) {
	if deps != nil {
		for key, value := range deps.Map() {
			s, _ := value.String()
			versions[key] = parseVersion(s)
		}
	}
}

func parseVersion(version string) string {
	if removeSpecialFirstChars {
		switch version[0:1] {
		case "^", "*", "~":
			return version[1:len(version)]
		}
	}
	return version
}

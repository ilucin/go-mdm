package repo

import (
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
	"github.com/ilucin/go-mdm/enum"
)

const mockGithub = false
const githubToken = "cb835b1dc214c1e5c949cd4d4253e0545dd38619"

// GithubURL is a prefix for github repos
const GithubURL = "github.com/"

func getFileFromGithub(user string, repo string) (file []byte, err error) {
	fmt.Printf("Scanning github repo %s/%s", user, repo)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	b64fileContent, _, resp, err := client.Repositories.GetContents(user, repo, "package.json", nil)

	fmt.Printf("\n\nGITHUB STATUS\n------\n%s\n------\n\n", resp)

	file, err = b64fileContent.Decode()

	return file, err
}

func getMockPackageJSON() ([]byte, error) {
	mockBase64 := "ewogICJuYW1lIjogInBsdXR1cyIsCiAgInZlcnNpb24iOiAiMC4wLjEiLAogICJkZXNjcmlwdGlvbiI6ICJNb25leSBtYW5hZ2VtZW50IGFwcCIsCiAgInJlcG9zaXRvcnkiOiAiIiwKICAiZGVwZW5kZW5jaWVzIjoge30sCiAgImRldkRlcGVuZGVuY2llcyI6IHsKICAgICJjb25uZWN0LWxpdmVyZWxvYWQiOiAifjAuNS4zIiwKICAgICJncnVudCI6ICJ+MC40LjUiLAogICAgImdydW50LWF1dG9wcmVmaXhlciI6ICJeMy4wLjAiLAogICAgImdydW50LWJvd2VyLXJlcXVpcmVqcyI6ICJ+Mi4wLjAiLAogICAgImdydW50LWNvbnRyaWItY2xlYW4iOiAifjAuNi4wIiwKICAgICJncnVudC1jb250cmliLWNvbmNhdCI6ICJ+MC41LjEiLAogICAgImdydW50LWNvbnRyaWItY29ubmVjdCI6ICJ+MC4xMC4xIiwKICAgICJncnVudC1jb250cmliLWNvcHkiOiAifjAuOC4wIiwKICAgICJncnVudC1jb250cmliLWNzc21pbiI6ICJ+MC4xMi4yIiwKICAgICJncnVudC1jb250cmliLWhhbmRsZWJhcnMiOiAiXjAuMTAuMiIsCiAgICAiZ3J1bnQtY29udHJpYi1odG1sbWluIjogIn4wLjQuMCIsCiAgICAiZ3J1bnQtY29udHJpYi1pbWFnZW1pbiI6ICJ+MC45LjQiLAogICAgImdydW50LWNvbnRyaWItanNoaW50IjogIn4wLjExLjIiLAogICAgImdydW50LWNvbnRyaWItanN0IjogIn4wLjYuMCIsCiAgICAiZ3J1bnQtY29udHJpYi11Z2xpZnkiOiAifjAuOS4xIiwKICAgICJncnVudC1jb250cmliLXdhdGNoIjogIl4wLjYuMSIsCiAgICAiZ3J1bnQtbm90aWZ5IjogIl4wLjQuMSIsCiAgICAiZ3J1bnQtb3BlbiI6ICJ+MC4yLjMiLAogICAgImdydW50LXByb21wdCI6ICJeMS4zLjAiLAogICAgImdydW50LXJlcXVpcmVqcyI6ICJ+MC40LjIiLAogICAgImdydW50LXJldiI6ICJ+MC4xLjAiLAogICAgImdydW50LXNhc3MiOiAiXjEuMC4wIiwKICAgICJncnVudC1zaGVsbCI6ICJeMS4xLjIiLAogICAgImdydW50LXVzZW1pbiI6ICIwLjEuMTMiLAogICAgImpzaGludC1zdHlsaXNoIjogIn4xLjAuMSIsCiAgICAibG9hZC1ncnVudC10YXNrcyI6ICJ+My4xLjAiLAogICAgIm53IjogIl4wLjEyLjIiLAogICAgInRpbWUtZ3J1bnQiOiAifjEuMS4xIgogIH0sCiAgImVuZ2luZXMiOiB7CiAgICAibm9kZSI6ICI+PTAuMTAuMCIKICB9LAogICJjdXN0b20iOiB7CiAgICAiZ3J1bnQtdXNlbWluIjogIjAuMS4xMiIKICB9LAogICJtYWluIjogImluZGV4Lmh0bWwiLAogICJzY3JpcHRzIjogewogICAgInN0YXJ0IjogIm53IG53IgogIH0sCiAgIndpbmRvdyI6IHsKICAgICJ0b29sYmFyIjogdHJ1ZQogIH0KfQo="

	return base64.StdEncoding.DecodeString(mockBase64)
}

// ScanGithub scans github
func ScanGithub(url string) (file []byte, libType enum.LibType, err error) {
	// fmt.Printf("Parsing github %s", url)

	split := strings.Split(url, "/")
	user := split[1]
	repo := split[2]

	if mockGithub {
		file, err = getMockPackageJSON()
	} else {
		file, err = getFileFromGithub(user, repo)
	}

	libType = enum.NPM

	// fileContentString := string(file[:])
	// fmt.Println(fileContentString)

	return file, libType, err
}

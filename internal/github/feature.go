package github

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var branches []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "*") {
			branches = append(branches, strings.TrimPrefix(line, "* "))
		}
	}

	return branches, nil
}

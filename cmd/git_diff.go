package cmd

import (
	"os/exec"
	"strings"
)

type DiffList struct {
	Added    []string
	Modified []string
	Deleted  []string
}

func NewDiffList(capacity int) *DiffList {
	return &DiffList{
		Added:    make([]string, 0, capacity),
		Modified: make([]string, 0, capacity),
		Deleted:  make([]string, 0, capacity),
	}
}

func GetGitDiff(crrentRev string, nextRev string) (diffList *DiffList, err error) {
	out, e := exec.Command("git", "diff", "--name-status", crrentRev, nextRev).Output()
	if e != nil {
		err = e
	}
	tmp := strings.Replace(string(out), "\n", "\t", -1)
	str := strings.Split(tmp, "\t")

	diffList = NewDiffList(len(str))
	for i := 0; i < len(str); i += 2 {
		switch {
		case str[i] == "M":
			diffList.Modified = append(diffList.Modified, str[i+1])
		case str[i] == "A":
			diffList.Added = append(diffList.Added, str[i+1])
		case str[i] == "D":
			diffList.Deleted = append(diffList.Deleted, str[i+1])
		}
	}
	return
}

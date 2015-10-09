package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	DifupLogBranch = "difup-log"
)

func IsExistDifupLogBranch() bool {
	out, e := exec.Command("git", "branch").Output()
	if e != nil {
		return false
	}
	tmp := string(out)
	tmp = strings.Replace(tmp, "* ", "", -1)
	tmp = strings.Replace(tmp, "  ", "", -1)
	strs := strings.Split(tmp, "\n")
	for _, v := range strs {
		if v == DifupLogBranch {
			return true
		}
	}
	return false
}

func CreateDifupLogBranch() error {
	var err error
	_, err = exec.Command("git", "checkout", "--orphan", DifupLogBranch).Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("git", "rm", "-rf", "*").Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("git", "commit", "--allow-empty", "-m", "\"Create difup-log branch\"").Output()

	return err
}

func SaveRevision(name string, revision string) (err error) {
	if !IsExistDifupLogBranch() {
		err = CreateDifupLogBranch()
		if err != nil {
			return
		}
	}

	var buf bytes.Buffer
	buf.WriteString("./")
	buf.WriteString(name)
	logFilePath, _ := filepath.Abs(buf.String())

	e2 := ioutil.WriteFile(logFilePath, []byte(buf.String()), os.ModePerm)
	if e2 != nil {
	}

	return
}

func GetCurrentLog(name string) (crrentRev string, err error) {
	var (
		buf bytes.Buffer
	)
	buf.WriteString(DifupLogBranch)
	buf.WriteString(":")
	buf.WriteString(name)
	optStr := buf.String()

	out, e := exec.Command("git", "show", optStr).Output()
	if e != nil {
		err = e
		return
	}

	crrentRev = strings.Replace(string(out), "\n", "", -1)
	fmt.Printf("-----\nhash:%#v\n-----\n", crrentRev)

	return
}

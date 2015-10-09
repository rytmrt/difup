package cmd

import (
	"testing"
)

func TestGitDiff(t *testing.T) {
	crrentRev, nextRev := "3cf24108a6f4d3a290185c33c582ca592782af39", ""
	res, err := GetGitDiff(crrentRev, nextRev)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	modList := []string{"", ""}
	addList := []string{"", ""}
	delList := []string{""}

	if len(modList) == len(res.Modified) {
		for i, v := range res.Modified {
			if modList[i] != v {
				t.Errorf("ERROR: mod response error. %v : %v", v, modList[i])
			}
		}
	} else {
		t.Errorf("ERROR: modRes length is not matched. %v : %v", len(modList), len(res.Modified))
	}
	if len(addList) == len(res.Added) {
		for i, v := range res.Added {
			if addList[i] != v {
				t.Errorf("ERROR: add response error. %v : %v", v, addList[i])
			}
		}
	} else {
		t.Errorf("ERROR: addRes length is not matched. %v : %v", len(addList), len(res.Added))
	}
	if len(delList) == len(res.Deleted) {
		for i, v := range res.Deleted {
			if delList[i] != v {
				t.Errorf("ERROR: del response error. %v : %v", v, delList[i])
			}
		}
	} else {
		t.Errorf("ERROR: delRes length is not matched. %v : %v", len(delList), len(res.Deleted))
	}
}

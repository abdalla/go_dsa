package realworld_test

import (
	"testing"

	"github.com/abdalla/go_dsa/challenges/realworld"
)

func getFolders() []realworld.Folder {
	folders := make([]realworld.Folder, 0)
	folders = append(folders, realworld.Folder{
		Id:         0,
		Subfolders: []int{7, 3},
		Name:       "abc",
	})

	folders = append(folders, realworld.Folder{
		Id:         0,
		Subfolders: []int{},
		Name:       "xyz",
	})

	folders = append(folders, realworld.Folder{
		Id:         8,
		Subfolders: []int{},
		Name:       "def",
	})

	folders = append(folders, realworld.Folder{
		Id:         7,
		Subfolders: []int{9},
		Name:       "ijk",
	})

	folders = append(folders, realworld.Folder{
		Id:         9,
		Subfolders: []int{},
		Name:       "lmn",
	})

	return folders
}

func TestPath9(t *testing.T) {
	got := realworld.PrintPath(9, getFolders())
	want := "abc -> ijk -> lmn"

	if got != want {
		t.Errorf("want: %+v got: %+v", want, got)
	}
}

func TestPath8(t *testing.T) {
	got := realworld.PrintPath(8, getFolders())
	want := ""

	if got != want {
		t.Errorf("want: %+v got: %+v", want, got)
	}
}

func TestPath7(t *testing.T) {
	got := realworld.PrintPath(7, getFolders())
	want := ""

	if got != want {
		t.Errorf("want: %+v got: %+v", want, got)
	}
}

func TestPath0(t *testing.T) {
	got := realworld.PrintPath(0, getFolders())
	want := ""

	if got != want {
		t.Errorf("want: %+v got: %+v", want, got)
	}
}

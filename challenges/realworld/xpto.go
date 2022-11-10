package realworld

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Folder struct {
	Name       string
	Subfolders []int
	Id         int
}

var (
	results = make(map[int]string)
	folders = make([]Folder, 0)
	nodes   = make(map[int]Folder)
)

func getFolder(id int) *Folder {
	if folder, has := nodes[id]; has {
		return &folder
	}

	for _, folder := range folders {
		if folder.Id == id {
			nodes[id] = folder
			return &folder
		}
	}

	return nil
}

func getPath(folderId int, folder Folder) string {
	if v, exist := results[folderId]; exist {
		return v
	}

	if slices.Contains(folder.Subfolders, folderId) {
		return folder.Name
	}

	for _, subFolderId := range folder.Subfolders {
		node := getFolder(subFolderId)
		if node != nil {
			p := getPath(folderId, *node)
			if p != "" {
				results[node.Id] = fmt.Sprintf("%s -> %s", folder.Name, p)
				return results[node.Id]
			}
			return ""
		}
	}

	return ""
}

func PrintPath(folderId int, nodes []Folder) string {
	folders = append(folders, nodes...)

	if folderId <= 0 {
		return ""
	}

	if v, exist := results[folderId]; exist {
		return v
	}

	desiredFolder := getFolder(folderId)
	if desiredFolder == nil {
		return ""
	}

	for _, folder := range folders {
		if folder.Id == 0 {
			p := getPath(folderId, folder)

			if p != "" {
				results[folderId] = p
				break
			}
		}
	}

	if results[folderId] != "" {
		results[folderId] = fmt.Sprintf("%s -> %s", results[folderId], desiredFolder.Name)
	}

	fmt.Printf("%+v \n", results)

	return results[folderId]
}

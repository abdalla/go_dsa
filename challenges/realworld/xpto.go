/*
*
Given list of folders, print the path of a given folder from root. If there is no path to the root folder then return an empty string. Root level folders will have 0 as id. The structure of Folder is like this:

class Folder {

int id;

List<int> subfolders;

String name;

}

Ex:

list = [Folder(0, [7, 3], “abc”), Folder(0, [], “xyz”), Folder(8, [], “def”), Folder(7, [9], “ijk), Folder(9, [], “lmn”)]

printPath(9) = “abc” -> “ijk” -> “lmn”
printPath(8) = ""

Clarification: There can be multiple root level folders. All other folders have unique ids.

Note: ç method may be called multiple times. Design your solution taking that into account

//QUESTION:
// - HOW MANY FOLDERS ?
// -
*
*/
//    abc .             xyz
// 	/  \
// 7 .   3 .   8
// 9
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
)

func getNode(id int) *Folder {
	for _, folder := range folders {
		if folder.Id == id {
			return &folder
		}
	}

	return nil
}

func getPath(path int, folder Folder) string {
	if v, exist := results[path]; exist {
		return v
	}

	if slices.Contains(folder.Subfolders, path) {
		if results[path] != "" {
			return fmt.Sprintf("%s -> %s", results[path], folder.Name)
		}

		if folder.Id != 0 {
			return folder.Name
		}

		return ""
	}

	for _, subFolderId := range folder.Subfolders {
		node := getNode(subFolderId)
		if node != nil {
			return getPath(path, *node)
		}
	}

	return ""
}

// input = data: [Folder(0, [7, 3], “abc”), Folder(0, [], “xyz”), Folder(8, [], “def”), Folder(7, [9], “ijk), Folder(9, [], “lmn”)]
func PrintPath(path int, nodes []Folder) string {
	folders = append(folders, nodes...)

	if path <= 0 {
		return ""
	}

	if v, exist := results[path]; exist {
		return v
	}

	lastNode := getNode(path)
	if lastNode == nil {
		return ""
	}

	for _, folder := range folders {
		if folder.Id == 0 {
			p := getPath(path, folder)

			if p != "" {
				results[path] = fmt.Sprintf("%s -> %s", folder.Name, p)
				break
			}
		}
	}

	if results[path] != "" {
		results[path] = fmt.Sprintf("%s -> %s", results[path], lastNode.Name)
	}

	return results[path]
}

// package filetree provides utility function to create and manipulate the filetree
package filetree

import (
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
	"wikinow/internal/utils"

	log "github.com/sirupsen/logrus"
)

type TreeNodeType string

const (
	Root TreeNodeType = "Root"
	Dir  TreeNodeType = "Dir"
	File TreeNodeType = "File"
)

type TreeNode struct {
	Title    string       `json:"title"`
	IsActive bool         `json:"isActive"`
	Order    int64        `json:"order"`
	Type     TreeNodeType `json:"type"`
	Path     string       `json:"path"`
	Children []TreeNode   `json:"children,omitempty"`
}

// This function orchestrates the creation of file tree structures
func GetFileTree(root string, currentPath string) (*TreeNode, error) {
	output, err := walkDir(root)
	if err != nil {
		return nil, err
	}

	keys := getSortedKeys(output)
	err = buildNodes(output, keys, root, currentPath)
	if err != nil {
		return nil, err
	}
	entrypoint := linkNodes(output, keys)

	return output[entrypoint].(*TreeNode), nil
}

// This function adds child nodes as children
func linkNodes(input map[string]interface{}, keys []string) string {
	entrypoint := keys[0]

	for _, k := range keys {
		v, ok := input[k].(*TreeNode)
		if !ok {
			continue
		}
		i := strings.LastIndex(k, "/")
		if i == -1 {
			continue
		}
		parentK := k[:i]
		if parentNode, exists := input[parentK]; exists {
			if parent, ok := parentNode.(*TreeNode); ok {
				parent.Children = append(parent.Children, *v)
			}
		}
		entrypoint = k
	}
	return entrypoint
}

// This function converts the paths into TreeNodes
func buildNodes(input map[string]interface{}, keys []string, root string, currentPath string) error {
	for _, k := range keys {
		parent := &TreeNode{
			Path: k + "/" + "main.md",
			Type: Dir,
		}
		title, order, err := utils.GetFileTitleAndOrder(parent.Path)
		if err != nil {
			return err
		}
		parent.Title = title
		parent.Order = order
		parent.Path = NormalizePath(parent.Path, root)
		parent.IsActive = isActive(parent, currentPath)

		if files, ok := input[k].([]string); ok {
			for _, file := range files {
				if file == "main.md" {
					continue
				}

				child := &TreeNode{
					Path: k + "/" + file,
					Type: File,
				}
				title, order, err := utils.GetFileTitleAndOrder(child.Path)
				if err != nil {
					return err
				}
				child.Title = title
				child.Order = order
				child.Path = NormalizePath(child.Path, root)
				child.IsActive = isActive(child, currentPath)

				parent.Children = append(parent.Children, *child)
			}
		}

		input[k] = parent
	}

	return nil
}

func isActive(node *TreeNode, currentPath string) bool {
	nodePath := node.Path
	if node.Type == Dir {
		if strings.HasSuffix(nodePath, "main") {
			nodePath = strings.TrimRight(nodePath, "main")
		}

		if strings.HasSuffix(currentPath, "main") {
			currentPath = strings.TrimRight(currentPath, "main")
		}
	}

	return nodePath == currentPath
}

// Correct the path removing the .md extension and the root directory
func NormalizePath(input string, root string) string {
	newPath := strings.ReplaceAll(input, root, "")
	newPath = strings.ReplaceAll(newPath, ".md", "")
	return "/wiki" + newPath
}

func getSortedKeys(input map[string]interface{}) []string {
	keys := []string{}
	for k := range input {
		keys = append(keys, k)
	}
	slices.SortStableFunc(keys, func(a, b string) int {
		return len(b) - len(a)
	})

	return keys
}

// walkDir traverses the root directory and builds a map of paths to files
func walkDir(root string) (map[string]interface{}, error) {
	output := make(map[string]interface{})

	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			output[path] = []string{}
			return nil
		}

		i := strings.LastIndex(path, entry.Name())
		key := path[:i-1]
		if _, exists := output[key]; !exists {
			output[key] = []string{}
		}
		output[key] = append(output[key].([]string), entry.Name())
		return nil
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetRelativePath(url string, path string) string {
	relativePath, err := filepath.Rel(url, path)
	if err != nil {
		log.Fatal("Error while getting relative path")
	}

	return relativePath
}

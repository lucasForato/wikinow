package filetree

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func Tree(projectRoot string, currentPath string) (TreeNode, error) {
	paths, err := GetFilePaths(projectRoot)
	if err != nil {
		return TreeNode{}, err
	}

	rootNode := TreeNode{Title: "Root", Type: Root, Path: "/wiki", Children: []TreeNode{}}

	for _, path := range paths {
		// Clean the path
		path = strings.ReplaceAll(path, "//", "/")
		path = strings.TrimSuffix(path, ".md")
		// Ensure that path starts with "/wiki"
		if !strings.HasPrefix(path, "/wiki") {
			path = "/wiki" + path
		}
		pathParts := strings.Split(path[5:], "/") // Remove "/wiki" and split
		addNode(&rootNode, pathParts, path == currentPath)
	}

	return rootNode, nil
}

func addNode(parent *TreeNode, pathParts []string, isActive bool) {
	if len(pathParts) == 0 {
		return
	}

	// Find or create the node for the current part
	var childNode *TreeNode
	for i := range parent.Children {
		if parent.Children[i].Title == pathParts[0] {
			childNode = &parent.Children[i]
			break
		}
	}

	if childNode == nil {
		nodeType := File
		if len(pathParts) > 1 {
			nodeType = Dir
		}
		childNode = &TreeNode{
			Title:    pathParts[0],
			Type:     nodeType,
			Path:     "/wiki/" + strings.Join(pathParts, "/"),
			IsActive: isActive,
		}
		parent.Children = append(parent.Children, *childNode)
	}

	// Recurse to add remaining parts
	if len(pathParts) > 1 {
		addNode(childNode, pathParts[1:], isActive)
	}
}

func GetFilePaths(root string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			// Format the path with leading slash and ensure no double slashes
			formattedPath := "/wiki/" + strings.ReplaceAll(relPath, "\\", "/")
			filePaths = append(filePaths, formattedPath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func PrintTree(node TreeNode, indent string) {
	fmt.Println(indent + node.Title)
	for _, child := range node.Children {
		PrintTree(child, indent+"--")
	}
}

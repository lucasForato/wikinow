package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type TreeNodeType string

const (
	Root TreeNodeType = "Root"
	Dir  TreeNodeType = "Dir"
	File TreeNodeType = "File"
)

type TreeNode struct {
	Title        string       `json:"title"`
	IsActive     bool         `json:"isActive"`
	Order        int64        `json:"order"`
	Type         TreeNodeType `json:"type"`
	Path         string       `json:"path"`
	RelativePath string       `json:"relativePath"`
	Children     []TreeNode   `json:"children,omitempty"`
}

func GetFileTree(rootPath string, currentPath string) *TreeNode {
	rootNode := &TreeNode{
		Title:        "root",
		Type:         Root,
		Path:         rootPath,
		RelativePath: "./",
		Children:     []TreeNode{},
	}

	if currentPath == "/" {
		rootNode.IsActive = true
	}

	dir := filepath.Dir(rootPath + currentPath)
	err := buildTree(rootNode, rootPath, dir, rootPath+currentPath)
	if err != nil {
		log.Fatal("Error while building the file tree:", err)
	}

	return rootNode
}

func buildTree(parentNode *TreeNode, currentPath string, directory string, accessedPath string) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

  if parentNode.Type == Root {
    relativePath, err := filepath.Rel(directory, currentPath)
    if err != nil {
			return errors.New("Error while getting relative path")
		}
    if strings.Contains(relativePath, ".md") {
			parentNode.RelativePath = strings.TrimRight(relativePath, ".md")
		} else {
			parentNode.RelativePath = strings.Join([]string{relativePath, "/"}, "")
		}
  }

	for _, entry := range entries {
		path := filepath.Join(currentPath, entry.Name())

		relativePath, err := filepath.Rel(directory, path)

		if err != nil {
			return errors.New("Error while getting relative path")
		}

		node := TreeNode{
			Path:     path,
			Children: []TreeNode{},
		}

    if accessedPath[len(accessedPath)-1] == '/' {
      alteredPath := accessedPath + "main.md"
      if alteredPath == path {
        parentNode.IsActive = true
      }
    } else {
      alteredPath := accessedPath + ".md"
      if alteredPath == path {
        node.IsActive = true
      }
    }


		if strings.Contains(relativePath, ".md") {
			node.RelativePath = strings.TrimRight(relativePath, ".md")
		} else {
			node.RelativePath = strings.Join([]string{relativePath, "/"}, "")
		}

		if entry.Name() == "main.md" {
			title, order, err := GetFileTitleAndOrder(path)
			if err != nil {
				return err
			}
			(*parentNode).Title = title
			(*parentNode).Order = order

		} else if entry.IsDir() {
			node.Type = Dir
			err = buildTree(&node, path, directory, accessedPath)
			if err != nil {
				return err
			}
		} else {
			title, order, err := GetFileTitleAndOrder(node.Path)
			if err != nil {
				return err
			}
			node.Type = File
			node.Title = title
			node.Order = order
		}

		// Append the node to the parent
		sort.Slice(parentNode.Children, func(i, j int) bool {
			return parentNode.Children[i].Order < parentNode.Children[j].Order
		})

		parentNode.Children = append(parentNode.Children, node)
	}

	return nil
}

// PrintTreeAsJSON prints the TreeNode as JSON
func (n *TreeNode) PrintTreeAsJSON() {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		log.Fatal("Error while marshalling tree to JSON:", err)
	}
	log.Println(string(jsonData))
}

func ParseAnchor(node TreeNode) template.HTML {
	return template.HTML(fmt.Sprintf(`<a href="%s" class="text-amber-600">%s</a>`, node.RelativePath, node.Title))
}

func GetIconClasses(level int) string {
	str := fmt.Sprintf("%dpx", (level * 8))
	return fmt.Sprintf("ml-[%s]", str)
}

package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TreeNodeType string

const (
	Dir  TreeNodeType = "Dir"
	File TreeNodeType = "File"
)

type TreeNode struct {
	Title        string       `json:"title"`
	Type         TreeNodeType `json:"type"`
	Path         string       `json:"path"`
	RelativePath string       `json:"relativePath"`
	Children     []TreeNode   `json:"children,omitempty"`
}

func GetFileTree(rootPath string) *TreeNode {
	rootNode := &TreeNode{
		Title:    "root",
		Type:     Dir,
		Path:     rootPath,
		Children: []TreeNode{},
	}

	err := buildTree(rootPath, rootNode, rootPath)
	if err != nil {
		log.Fatal("Error while building the file tree:", err)
	}

	return rootNode
}

func buildTree(rootPath string, parentNode *TreeNode, currentPath string) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := filepath.Join(currentPath, entry.Name())
		relativePath, err := filepath.Rel(rootPath, path)
		node := TreeNode{
			Path:         path,
			RelativePath: strings.TrimRight(relativePath, ".md"),
			Children:     []TreeNode{},
		}

		if entry.IsDir() {
			node.Type = Dir
			err = buildTree(rootPath, &node, path)
			if err != nil {
				return err
			}
		} else {
			title, err := GetFileTitle(node.Path)
			if err != nil {
				return err
			}
			node.Type = File
			node.Title = title
		}

		// Append the node to the parent
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

func GetLevelClasses(level int) string {
  str := fmt.Sprintf("%dpx", (level * 15))
  return fmt.Sprintf("ml-[%s] border-l-2 gap-0.5 border-neutral-700 flex flex-col", str)
}

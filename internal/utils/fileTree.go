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
	Order        int64        `json:"order"`
	Type         TreeNodeType `json:"type"`
	Path         string       `json:"path"`
	RelativePath string       `json:"relativePath"`
	Children     []TreeNode   `json:"children,omitempty"`
}

func GetFileTree(rootPath string, currentPath string) *TreeNode {
	rootNode := &TreeNode{
		Title:    "root",
		Type:     Root,
		Path:     rootPath,
    RelativePath: "./",
		Children: []TreeNode{},
	}

  dir := filepath.Dir(rootPath+currentPath)
	err := buildTree(rootNode, rootPath, dir)
	if err != nil {
		log.Fatal("Error while building the file tree:", err)
	}

	return rootNode
}

func buildTree(parentNode *TreeNode, currentPath string, directory string) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
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
			err = buildTree(&node, path, directory)
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

func GetLevelClasses(level int) string {
	str := fmt.Sprintf("%dpx", (level * 15))
	return fmt.Sprintf("ml-[%s] border-l-2 gap-0.5 border-neutral-700 flex flex-col", str)
}

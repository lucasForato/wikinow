package cmd

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"wikinow/infra/logger"

	"github.com/spf13/cobra"
)

//go:embed files
var files embed.FS

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes wikinow by adding a /wiki directory and configuration file",
	Long: `Initialization command:

  This command will create a config file at the current directory and a /wiki
  directory with a main.md file in it.

  You should run this command before using any other command. 
  `,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		root, err := os.Getwd()
		if err != nil {
			logger.Error("Error retrieving working directory.", err)
		}

		wikidir := createWikidir(root)
		imgdir := createImgdir(root)
		copyMainFile(wikidir)
    copyImgFile(imgdir)

		logger.Info("Initialization completed ★")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createWikidir(root string) string {
	wikidir := filepath.Join(root, "wiki")
	if err := os.MkdirAll(wikidir, fs.ModePerm); err != nil {
		logger.Error("Error creating directory.", err)
	}
	return wikidir
}

func createImgdir(root string) string {
	imgdir := filepath.Join(root, "images")
	if err := os.MkdirAll(imgdir, fs.ModePerm); err != nil {
		logger.Error("Error creating directory.", err)
	}
	return imgdir
}

func copyMainFile(wikidir string) {
	mainFileOrigin, err := files.ReadFile("files/main.md")
	if err != nil {
		logger.Error("Error retrieving main documentation file.", err)
	}
	mainFileDest := filepath.Join(wikidir, "main.md")
	err = os.WriteFile(mainFileDest, mainFileOrigin, 0644)
	if err != nil {
		logger.Error("Error creating main documentation file.", err)
	}
}

func copyImgFile(imgdir string) {
	imgOrigin, err := files.ReadFile("files/images/example.jpeg")
	if err != nil {
		logger.Error("Error retrieving main documentation file.", err)
	}
	imgDest := filepath.Join(imgdir, "example.jpeg")
	err = os.WriteFile(imgDest, imgOrigin, 0644)
	if err != nil {
		logger.Error("Error creating main documentation file.", err)
	}
}

func copyConfig(root string) {
	configFileOrigin, err := files.ReadFile("files/wikinow.yml")
	if err != nil {
		logger.Error("Error reading config file.", err)
	}
	configFileDest := filepath.Join(root, "wikinow.yml")
	err = os.WriteFile(configFileDest, configFileOrigin, 0644)
	if err != nil {
		logger.Error("Error creating config file.", err)
	}
}

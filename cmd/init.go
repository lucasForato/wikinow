package cmd

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
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
		workingDir, err := os.Getwd()
		if err != nil {
			log.WithError(err).Fatal("Error retrieving working directory.")
		}

		wikiDir := filepath.Join(workingDir, "wiki")
		if err := os.MkdirAll(wikiDir, fs.ModePerm); err != nil {
			log.WithError(err).Fatal("Error creating directory.")
		}

		mainFileOrigin, err := files.ReadFile("files/main.md")
		if err != nil {
			log.WithError(err).Fatal("Error retrieving main documentation file.")
		}
		mainFileDest := filepath.Join(wikiDir, "main.md")
		err = os.WriteFile(mainFileDest, mainFileOrigin, 0644)
		if err != nil {
			log.WithError(err).Fatal("Error creating main documentation file.")
		}

		configFileOrigin, err := files.ReadFile("files/config.yml")
		if err != nil {
			log.WithError(err).Fatal("Error reading config file.")
		}
		configFileDest := filepath.Join(workingDir, "config.yml")
		err = os.WriteFile(configFileDest, configFileOrigin, 0644)
		if err != nil {
			log.WithError(err).Fatal("Error creating config file.")
		}

		log.Info("Initialization complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

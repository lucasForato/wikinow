package cmd

import (
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

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
		log.Info("Initializing wikinow...")

		path, err := os.Getwd()
		if err != nil {
			log.Fatal("Error while retrieving the current directory.")
		}

		wikiDir := filepath.Join(path, "wiki")
		if err := os.MkdirAll(wikiDir, fs.ModePerm); err != nil {
			log.WithFields(log.Fields{
				"directory": wikiDir,
			}).Fatal("Error creating directory.")
		}

    mainFileSource := "./files/main.md"
		mainFileDest := filepath.Join(wikiDir, "main.md")
		mainInput, err := os.ReadFile(mainFileSource)
		if err != nil {
			log.WithFields(log.Fields{
				"source": mainFileSource,
			}).Fatal("Error reading main documentation file.")
		}

		err = os.WriteFile(mainFileDest, mainInput, 0644)
		if err != nil {
			log.WithFields(log.Fields{
				"destination": mainFileDest,
			}).Fatal("Error creating main documentation file.")
		}

		configFileSource := "./files/config.yml"
		configFileDest := filepath.Join(path, "config.yml")
		configInput, err := os.ReadFile(configFileSource)
		if err != nil {
			log.WithFields(log.Fields{
				"source": configFileSource,
			}).Fatal("Error reading config file.")
		}

		err = os.WriteFile(configFileDest, configInput, 0644)
		if err != nil {
			log.WithFields(log.Fields{
				"destination": configFileDest,
			}).Fatal("Error creating config file.")
		}

		log.Info("Initialization complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

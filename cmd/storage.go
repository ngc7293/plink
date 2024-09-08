package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ngc7293/plink/internal/api"
	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Storage commands",
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files on remote",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ListFiles(args)
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete files on remote",
	RunE: func(cmd *cobra.Command, args []string) error {
		return DeleteFile(args)
	},
}

func ListFiles(args []string) error {
	config, err := ReadConfig()

	if err != nil {
		slog.Error("could not read config", "error", err)
		return err
	}

	if len(args) == 0 {
		storage, err := api.GetStorage(config.Host, config.Username, config.Password)

		if err != nil {
			slog.Error("could not get storage", "error", err)
			return err
		}

		for _, info := range storage.StorageList {
			fmt.Println(strings.TrimLeft(info.Path, "/"))
		}
	} else if len(args) == 1 {
		files, err := api.GetFiles(args[0], config.Host, config.Username, config.Password)

		if err != nil {
			slog.Error("could not get storage", "error", err)
			return err
		}

		for _, info := range files.Children {
			line := fmt.Sprintf("%s\t%s", info.Name, info.DisplayName)

			if info.Type == "FOLDER" {
				line += "/"
			}

			fmt.Println(line)
		}
	} else {
		return errors.New("invalid usage")
	}

	return nil
}

func DeleteFile(args []string) error {
	config, err := ReadConfig()

	if err != nil {
		slog.Error("could not read config", "error", err)
		return err
	}

	if len(args) != 1 {
		return errors.New("invalid usage")
	}

	if err := api.DeleteFile(args[0], config.Host, config.Username, config.Password); err != nil {
		slog.Error("could not remove file", "error", err)
		return err
	}

	return nil
}

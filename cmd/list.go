package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/smashedr/bup/internal/styles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func listCmd(cmd *cobra.Command, args []string) {
	log.Debug("listCmd:", "args", args)

	destination := viper.GetString("destination")
	if len(args) == 0 {
		listDir(destination, "All Backups")
	} else {
		for i := 0; i < len(args); i++ {
			path := filepath.Join(destination, args[i])
			listDir(path, args[i])
		}
	}
}

func listDir(path, header string) {
	log.Infof("Destination: %s", path)
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
	log.Debugf("entries: %v", entries)
	if len(entries) == 0 {
		log.Warnf("No backups for: %v", header)
		return
	}
	var rows [][]string
	for _, e := range entries {
		rows = append(rows, []string{e.Name()})
	}
	log.Debugf("rows: %v", rows)
	styles.RenderTable(rows, header)
}

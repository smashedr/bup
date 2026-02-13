package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/smashedr/bup/internal/styles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func infoCmd(cmd *cobra.Command, args []string) {
	log.Debug("infoCmd:", "args", args)

	styles.PrintKV("Config Flag", fmt.Sprintf("%q", cfgFile))
	styles.PrintKV("Config Used", viper.ConfigFileUsed())

	destination := viper.GetString("destination")
	styles.PrintKV("Destination", destination)

	styles.PrintKV("Excludes", fmt.Sprintf("%v", viper.GetStringSlice("excludes")))
}

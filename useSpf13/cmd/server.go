package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run in server mode",
	Long:  "run in server mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transcode server " + strings.Join(args, " "))
	},
}

func init() {
	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
	serverCmd.PersistentFlags().IntP("port", "", 0, "listening port")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))
}

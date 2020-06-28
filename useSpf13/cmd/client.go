package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "run in client mode",
	Long:  "run in client mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transcode client " + strings.Join(args, " "))
	},
}

func init() {
	clientCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
	clientCmd.PersistentFlags().IntP("poolsize", "", 0, "maximum concurrent number")
	viper.BindPFlag("poolsize", clientCmd.PersistentFlags().Lookup("poolsize"))
}

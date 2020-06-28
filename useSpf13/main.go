package main

import (
	"fmt"
	"wheel/useSpf13/cmd"

	"github.com/spf13/viper"
)

func main() {
	cmd.Execute()
	fmt.Println(viper.GetInt("port"))
	fmt.Println(viper.GetString("server"))
}

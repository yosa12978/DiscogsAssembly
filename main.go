package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yosa12978/DiscogsAssembly/cmd"
)

func init() {
	cobra.OnInitialize(initConfig)
}

const config_file = "./configs/config.json"
const basic_config = "{\"discogs\": {\"token\": \"\"}}"

func initConfig() {
	_, err := os.Stat(config_file)
	if err != nil {
		os.Mkdir("configs", 0666)
		file, err := os.Create("./configs/config.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()
		_, err = file.Write([]byte(basic_config))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	viper.SetConfigFile(config_file)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Can't read config: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	cmd.Execute()
}

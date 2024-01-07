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

const config_file string = "/configs/config.json"

const basic_config = "{\"discogs\": {\"token\": \"\"}}"

var home_dir string

func initConfig() {
	home_dir = os.Getenv("DISCASM_HOME")
	if home_dir == "" {
		fmt.Println("Please setup a DISCASM_HOME env variable")
		return
	}
	config_path := home_dir + config_file
	_, err := os.Stat(config_path)
	if err != nil {
		os.Mkdir(home_dir+"/configs", 0666)
		file, err := os.Create(config_path)
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
	viper.SetConfigFile(config_path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Can't read config: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	cmd.Execute()
}

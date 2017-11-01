package main

import (
	"flag"
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/train-cat/sniffer-transilien/cache"
	"github.com/train-cat/sniffer-transilien/metadata"
	"github.com/train-cat/sniffer-transilien/utils"
	"github.com/train-cat/sniffer-transilien/sncf"
	"github.com/train-cat/client-train-go"
)

func init() {
	cfgFile := flag.String("config", "", "config file")
	flag.Parse()

	if *cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(*cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".rc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Init function should be call after viper initialisation
	utils.Init()
	sncf.Init()
	metadata.Init()
	cache.Init()

	traincat.SetConfig(traincat.Config{
		Host: viper.GetString("api-train.host"),
		Username: viper.GetString("api-train.username"),
		Password: viper.GetString("api-train.password"),
	})
}

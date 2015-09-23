package main

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

// Config File stuff
const confFile string = "config.toml"
const confCommonEnv string = "common"

type tomlConfig struct {
	Environments map[string]environmentCfg
}

type environmentCfg struct {
	APIURL    string `toml:"api_url"`
	APIUser   string `toml:"api_user"`
	APIPasswd string `toml:"api_passwd"`
}

var env environmentCfg

func initializeConfig() {
	//
	// read config from config file
	//
	realConfFile := confFile
	confPath, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(path.Join(confPath, fmt.Sprintf(".%s", confFile))); os.IsNotExist(err) {
		log.Info("%s/.%s does not exist, trying ./%s", confPath, confFile, confFile)

		if _, err := os.Stat(confFile); os.IsNotExist(err) {
			log.Critical("%s does not exist./%s", confFile)
			os.Exit(1)
		}
	} else {
		realConfFile = path.Join(confPath, fmt.Sprintf(".%s", confFile))
	}
	var config tomlConfig
	if _, err := toml.DecodeFile(realConfFile, &config); err != nil {
		log.Critical("%s", err)
		os.Exit(1)
	}

	var found bool
	envs := make([]string, len(config.Environments))
	i := 0
	for e := range config.Environments {
		envs[i] = e
		if e == *environment {
			found = true
		}
		i++
	}
	if !found {
		log.Critical("Unknown environment. Did you mean one of %s", envs)
		os.Exit(1)
	}
	env = config.Environments[*environment]
	c := config.Environments[confCommonEnv]
	log.Info("Requested Environment: %s, URL: %s, User: %s, Password: %s\n", *environment, env.APIURL, env.APIUser, env.APIPasswd)
	log.Info("Common Environment: %s, URL: %s, User: %s, Password: %s\n", confCommonEnv, c.APIURL, c.APIUser, c.APIPasswd)
	if env.APIURL == "" {
		env.APIURL = c.APIURL
	}
	if env.APIUser == "" {
		env.APIUser = c.APIUser
	}
	if env.APIPasswd == "" {
		env.APIPasswd = c.APIPasswd
	}
}

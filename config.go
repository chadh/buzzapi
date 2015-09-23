package main

import (
	"os"

	"github.com/BurntSushi/toml"
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
	var config tomlConfig
	if _, err := toml.DecodeFile(confFile, &config); err != nil {
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

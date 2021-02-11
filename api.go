package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/claudiocandio/gemini-api"
	"github.com/claudiocandio/gemini-api/logger"
	"gopkg.in/yaml.v2"
)

type gemini_yml struct {
	Gemini_api_credentials struct {
		Gemini_api_key        string `yaml:"gemini_api_key"`
		Gemini_api_secret     string `yaml:"gemini_api_secret"`
		Gemini_api_production string `yaml:"gemini_api_production"`
	} `yaml:"gemini_api_credentials"`
}

func start_api(gemini_config_yml string) (*gemini.Api, error) {

	var gc gemini_yml
	var gemini_api_production bool
	var err error

	if gemini_config_yml != "" {
		fp, err := os.Open(gemini_config_yml)
		if err != nil {
			logger.Debug("Open configuration file error",
				fmt.Sprintf("gemini_config_yml:%s", gemini_config_yml),
				fmt.Sprintf("error:%s", err))
			return nil, fmt.Errorf("Cannot open configuration file: %s", gemini_config_yml)
		}
		decoder := yaml.NewDecoder(fp)
		err = decoder.Decode(&gc)
		if err != nil {
			return nil, err
		}
		fp.Close()
	}

	if gc.Gemini_api_credentials.Gemini_api_key == "" {
		gc.Gemini_api_credentials.Gemini_api_key = os.Getenv("GEMINI_API_KEY")
		if gc.Gemini_api_credentials.Gemini_api_key == "" {
			errlog.Println("Missing GEMINI_API_KEY")
		}
	}

	if gc.Gemini_api_credentials.Gemini_api_secret == "" {
		gc.Gemini_api_credentials.Gemini_api_secret = os.Getenv("GEMINI_API_SECRET")
		if gc.Gemini_api_credentials.Gemini_api_secret == "" {
			errlog.Println("Missing GEMINI_API_SECRET")
		}
	}

	if gc.Gemini_api_credentials.Gemini_api_production != "" {
		gemini_api_production, err = strconv.ParseBool(gc.Gemini_api_credentials.Gemini_api_production)
		if gc.Gemini_api_credentials.Gemini_api_production != "" && err != nil {
			errlog.Println("GEMINI_API_PRODUCTION environment variable must be set as true or false")
			gc.Gemini_api_credentials.Gemini_api_production = ""
		}
	} else {
		gc.Gemini_api_credentials.Gemini_api_production = os.Getenv("GEMINI_API_PRODUCTION")
		if gc.Gemini_api_credentials.Gemini_api_production == "" {
			errlog.Println("Missing GEMINI_API_PRODUCTION")
		}
		gemini_api_production, err = strconv.ParseBool(gc.Gemini_api_credentials.Gemini_api_production)
		if gc.Gemini_api_credentials.Gemini_api_production != "" && err != nil {
			errlog.Println("GEMINI_API_PRODUCTION environment variable must be set as true or false")
			gc.Gemini_api_credentials.Gemini_api_production = ""
		}
	}

	if gc.Gemini_api_credentials.Gemini_api_key == "" ||
		gc.Gemini_api_credentials.Gemini_api_secret == "" ||
		gc.Gemini_api_credentials.Gemini_api_production == "" {
		return nil, fmt.Errorf("Set Gemini API credentials.")
	}

	if gemini_api_production {
		logger.Debug("Connecting to Gemini Production site.")
	} else {
		logger.Debug("Connecting to Gemini Sandbox site.")
	}

	api := gemini.New(gemini_api_production,
		gc.Gemini_api_credentials.Gemini_api_key,
		gc.Gemini_api_credentials.Gemini_api_secret)

	//will show gemini api key & secret !!!
	logger.Trace("Gemini",
		fmt.Sprintf("api:%v", api),
		fmt.Sprintf("gc:%v", gc))

	return api, nil
}

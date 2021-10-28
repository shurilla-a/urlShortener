package ymalConfigParsing

import (
"io/ioutil"
"gopkg.in/yaml.v2"
"urlShortener/logger"
)
type ConfigYaml struct {
	HostDb string `yaml:"hostdb"`
	PortDb string `yaml:"portdb"`
	DbName string `yaml:"dbname"`
	CoreCount string `yaml:"corecount"`
	}

func confiParsinsg(fileNme string){

}

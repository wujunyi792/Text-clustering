package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"os"
)

// 解析yml文件
type ClusteringConfig struct {
	BadWords               []string   `yaml:"badWords"`
	ReBadWords             []Restruct `yaml:"reBadWords"`
	LegitimateMatchingRate float64    `yaml:"legitimateMatchingRate"`
}

type Restruct struct {
	Start     string `yaml:"start"`
	End       string `yaml:"end"`
	MaxLength int    `yaml:"maxLength"`
}

func (c *ClusteringConfig) GetConf() *ClusteringConfig {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

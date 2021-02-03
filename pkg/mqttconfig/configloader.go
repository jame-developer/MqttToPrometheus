package mqttconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func LoadMqttConfig() (*MqttConfig, error) {
	filename, _ := filepath.Abs("config/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config = &MqttConfig{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

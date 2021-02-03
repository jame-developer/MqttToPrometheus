package mqttconfig

type MqttConfig struct {
	Server struct {
		Host     string      `yaml:"host"`
		Port     int         `yaml:"port"`
		Username interface{} `yaml:"username"`
		Password interface{} `yaml:"password"`
		Topics   []string    `yaml:"topics"`
	} `yaml:"server"`
}

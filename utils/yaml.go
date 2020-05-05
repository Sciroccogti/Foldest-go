package utils

// Yaml : a struct for conf.yml
type Yaml struct {
	Verbose   bool   `yaml:"verbose"`
	Targetdir string `yaml:"targetdir"`
	Tmpbin    struct {
		Enable bool     `yaml:"enable"`
		Name   string   `yaml:"name"`
		Thresh int      `yaml:"treshday"`
		Delete int      `yaml:"deleteday"`
		Ignore []string `yaml:"ignore"`
	}
}

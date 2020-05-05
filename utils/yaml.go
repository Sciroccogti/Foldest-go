package utils

// Conf : a struct for conf.yml
type Conf struct {
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

// Rule : a template struct for a rule
type Rule struct {
	Enable bool     `yaml:"enable"`
	Name   string   `yaml:"name"`
	Regex  []string `yaml:"regex"`
}

// Rules :
type Rules struct {
	Rule1  Rule `yaml:"rule1"`
	Rule2  Rule `yaml:"rule2"`
	Rule3  Rule `yaml:"rule3"`
	Rule4  Rule `yaml:"rule4"`
	Rule5  Rule `yaml:"rule5"`
	Rule6  Rule `yaml:"rule6"`
	Rule7  Rule `yaml:"rule7"`
	Rule8  Rule `yaml:"rule8"`
	Rule9  Rule `yaml:"rule9"`
	Rule10 Rule `yaml:"rule10"`
}

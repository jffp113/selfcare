package config

import (
	"strings"
	"time"
)

type Config struct {
	Context  string             `yaml:"current_context"`
	Contexts map[string]Context `yaml:"contexts"`

	Projects map[string]Project `yaml:"projects"`
}

type Context struct {
	Timezone string `yaml:"timezone"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Project struct {
	Client       string `yaml:"client"`
	Package      string `yaml:"package"`
	Project      string `yaml:"project"`
	ServiceLevel string `yaml:"service_level"`
	Option       string `yaml:"option"`
}

type Timesheet struct {
	Entries []Entry `yaml:"entries"`
}

type Entry struct {
	Project string   `yaml:"project"`
	Time    float64  `yaml:"time"`
	Day     Date     `yaml:"day"`
	Tasks   []string `yaml:"tasks"`
}

type Date struct {
	Time time.Time
}

func (t *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return nil
	}

	tt, err := time.Parse("02-01-2006", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t Date) MarshalYAML() (interface{}, error) {
	return t.Time.Format("2006-01-02"), nil
}

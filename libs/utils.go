package libs

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type KeycloakInstance struct {
	Address  string
	User     string
	Password string
}

type Db struct {
	User     string
	Password string
}

type KeyclaokDb struct {
	Name     string
	User     string
	Password string
}

type Git struct {
	Protcol  string
	Url      string
	Name     string
	Server   string
	Password string
	User     string
	Branch   string
}

func GetYaml(path string) error {
	s, _ := ioutil.ReadFile(path)
	var yamldata map[string]interface{}
	yaml.Unmarshal([]byte(s), &yamldata)
	fmt.Println(yamldata["yuuvis"])

	return nil
}

func UpdateEnvValue(key, value string) {
	os.Setenv(key, value)
}

package conf

import (
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var ymlAliaser *ymlAlias

type ymlAlias struct {
	path    string
	aliases map[string]string
}

func (yml *ymlAlias) Path() string {
	return yml.path
}

func (yml *ymlAlias) Get(key string) string {
	for mapKey, val := range yml.aliases {
		if strings.Compare(key, mapKey) == 0 {
			return val
		}
	}

	return ""
}

func (yml *ymlAlias) All() map[string]string {
	return yml.aliases
}

func NewAlias(path string) *ymlAlias {

	aliaser := ymlAlias{
		path: path,
	}

	file, err := ioutil.ReadFile(aliaser.Path())

	if err != nil {
		log.Fatal(err)
	}

	errParse := yaml.Unmarshal(file, &aliaser.aliases)

	if errParse != nil {
		log.Fatal(errParse)
	}

	return &aliaser
}

func init() {
	ymlAliaser = NewAlias("aliases.yml")
}

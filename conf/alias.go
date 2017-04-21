package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const aliasesFile = "aliases.json"

//Aliases components aliases
var Aliases map[string]string

//Component receive an alias or complete components
//and return the component path
func Component(name string) string {
	for key, val := range Aliases {
		if strings.Compare(name, key) == 0 {
			return val
		}
	}

	return name
}

func loadAliases() {
	file, err := ioutil.ReadFile(aliasesFile)

	if err != nil {
		fmt.Println(err)
	}

	parseAlias(file, &Aliases)
}

func parseAlias(blob []byte, alias *map[string]string) {
	err := json.Unmarshal(blob, &alias)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	loadAliases()
}

package conf

import (
	"strings"
)

// Alias is the...
type Alias interface {
	Path() string
	Get(key string) string
	All() map[string]string
}

//Component receive an alias or complete components
//and return the component path
func Component(name string) string {
	res := ymlAliaser.Get(name)

	var component string

	if res == "" {
		component = name
	} else {
		component = res
	}

	if !strings.HasSuffix(component, "/") {
		component = component + "/"
	}

	return component
}

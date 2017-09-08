package conf

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

	if res == "" {
		return name
	}

	return res
}

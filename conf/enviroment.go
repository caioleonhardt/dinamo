package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const serverFile = "servers.json"

var conf Conf

//Conf configuration of service
type Conf struct {
	Enviroments []Env
	Aliases     map[string]string
	Default     string
}

//Env is the main structure
type Env struct {
	Name     string
	User     string
	Password string
	Servers  []Server
}

//Server is each host with dynamo
type Server struct {
	Name string
	Host string
	Port string
}

//URL build the url
func (s *Server) URL() string {
	return "http://" + s.Host + ":" + s.Port
}

func loadConf() {
	file, err := ioutil.ReadFile(serverFile)

	if err != nil {
		fmt.Println(err)
	}

	parseConf(file, &conf)
}

//Context get the context enviroment
func Context(env string) *Env {
	for _, val := range conf.Enviroments {
		if strings.Compare(val.Name, env) == 0 {
			return &val
		}
	}
	return nil
}

func parseConf(blob []byte, conf *Conf) {
	err := json.Unmarshal(blob, &conf)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

//DefaultEnviroment is the default enviroment
//to execute
func DefaultEnviroment() string {
	return conf.Default
}

func init() {
	loadConf()
}

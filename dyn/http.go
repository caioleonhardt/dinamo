package http

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/caioleonhardt/dinamo/conf"
)

// ChangeValue change value on dynamo
func ChangeValue(completeURL string, property string, newValue string) (*http.Response, error) {

	form := url.Values{}
	form.Add("propertyName", property)
	form.Add("newValue", newValue)
	form.Add("change", "Change Value")

	return basic(completeURL, form)
}

// InvokeMethod invokes method on dynamo
func InvokeMethod(completeURL string, methodName string) (*http.Response, error) {
	form := url.Values{}
	form.Add("invokeMethod", methodName)
	form.Add("submit", "Invoke Method")

	return basic(completeURL, form)
}

func basic(urlServer string, form url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", urlServer, strings.NewReader(form.Encode()))

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(conf.Enviroment.User, conf.Enviroment.Password)

	client := http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

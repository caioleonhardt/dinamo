package http

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
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
		fmt.Println(err.Error())
		os.Exit(1)
	}

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Lenght", "60")
	req.SetBasicAuth("admin", "admin")

	client := http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

package cmd

import (
	"fmt"
	"os"

	"sync"

	"github.com/caioleonhardt/dinamo/conf"
	"github.com/caioleonhardt/dinamo/dyn"
	"github.com/spf13/cobra"
)

var enviroment string
var invoke bool
var dynCtx = "/dyn/admin/nucleus"

//RootCmd default command
var RootCmd = &cobra.Command{
	Use:   "dinamo <component> <property> <value>",
	Short: "Dinamo is a fast Dynamo command executor",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Use dinamo <component> <property> <value>")
			os.Exit(0)
		}

		if len(enviroment) == 0 {
			if len(conf.DefaultEnviroment()) == 0 {
				fmt.Println("you should inform the enviroment to execute")
				os.Exit(0)
			}

			enviroment = conf.DefaultEnviroment()
		}

		ctx := conf.Context(enviroment)

		if ctx == nil {
			fmt.Println("Enviroment not found")
			os.Exit(0)
		}

		if invoke {
			invokeMethod(ctx.Servers, args)
		} else {
			changeProperty(ctx.Servers, args)
		}
	},
}

func changeProperty(servers []conf.Server, args []string) {
	if len(args) < 3 {
		fmt.Println("Use dinamo <component> <property> <value>")
		os.Exit(0)
	}
	var component = conf.Component(args[0])
	var property = args[1]
	var newValue = args[2]

	var wg sync.WaitGroup

	for _, server := range servers {
		for _, url := range server.URL() {
			wg.Add(1)

			var fullURL = url + dynCtx + component

			go func(fullURL string) {
				defer wg.Done()

				res, err := http.ChangeValue(fullURL, property, newValue)

				if err != nil {
					fmt.Printf("Execution %s %s\n", fullURL, "NOT OK")
				} else {
					fmt.Printf("Execution %s %s\n", fullURL, res.Status)
				}
			}(fullURL)
		}
	}
	wg.Wait()
}

func invokeMethod(servers []conf.Server, args []string) {
	if len(args) < 2 {
		fmt.Println("Use dinamo <component> <methodName>")
		os.Exit(0)
	}

	var component = conf.Component(args[0])
	var methodName = args[1]

	var wg sync.WaitGroup

	for _, server := range servers {
		for _, url := range server.URL() {
			wg.Add(1)

			var fullURL = url + dynCtx + component

			go func(fullURL string) {
				defer wg.Done()

				res, err := http.InvokeMethod(fullURL, methodName)

				if err != nil {
					fmt.Printf("Execution %s %s\n", fullURL, "NOT OK")
				} else {
					fmt.Printf("Execution %s %s\n", fullURL, res.Status)
				}
			}(fullURL)
		}
	}
	wg.Wait()
}

func init() {
	RootCmd.Flags().StringVarP(&enviroment, "env", "e", "", "enviroment to execute")
	RootCmd.Flags().BoolVarP(&invoke, "invoke", "i", false, "invoke method on dynamo")
}

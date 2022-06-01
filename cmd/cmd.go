package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains`,

	Run: func(cmd *cobra.Command, args []string) {
		httpMethod := ""
		prompt := &survey.Select{
			Message: "Choose a HTTP Method:",
			Options: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		}
		survey.AskOne(prompt, &httpMethod)

		client := resty.New()
		switch httpMethod {
		case "GET":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)

			resp, err := client.R().EnableTrace().Get(url)

			if err != nil {
				log.Println(err)
			}
			fmt.Println(resp)

		case "POST":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)
			body := ""
			prompt = &survey.Input{
				Message: "Enter Body:",
			}
			survey.AskOne(prompt, &body)

			resp, err := client.R().SetBody(body).Post(url)

			if err != nil {
				log.Println(err)
			}

			fmt.Println(resp)

		case "PATCH":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)
			body := ""
			prompt = &survey.Input{
				Message: "Enter Body:",
			}
			survey.AskOne(prompt, &body)

			resp, err := client.R().SetBody(body).Patch(url)

			if err != nil {
				log.Println(err)
			}

			fmt.Println(resp)

		case "PUT":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)
			body := ""
			prompt = &survey.Input{
				Message: "Enter Body:",
			}
			survey.AskOne(prompt, &body)

			resp, err := client.R().SetBody(body).Put(url)

			if err != nil {
				log.Println(err)
			}

			fmt.Println(resp)

		case "DELETE":
			url := ""
			prompt := &survey.Input{
				Message: "Enter URL:",
			}
			survey.AskOne(prompt, &url)

			resp, err := client.R().Delete(url)

			if err != nil {
				log.Println(err)
			}
			fmt.Println(resp)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

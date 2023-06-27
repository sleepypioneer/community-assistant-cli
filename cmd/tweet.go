/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sleepypioneer/community-assistant-cli/utils"
)

func ShortenText(inputText string) string {
	// return inputText[0:240]
	tweet, err := utils.GenerateTweet(inputText)
	if err != nil {
		fmt.Println("Error generating tweet: ", err)
		os.Exit(1)
	}
	return tweet
}

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Take input file with text and creates a tweet",
	Long: `Takes an input file with text and creates a short form version of it suitable for posting to Twitter. For example:
			mycli tweet -i input.txt -o output.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := cmd.Flag("input").Value.String()
		outputFile := cmd.Flag("output").Value.String()
		fmt.Println("input file: ", inputFile)
		fmt.Println("output file: ", outputFile)

		if inputFile == "" {
			fmt.Println("Input file is required")
			os.Exit(1)
		}

		if outputFile == "" {
			fmt.Println("Output file is required")
			os.Exit(1)
		}

		inputText, err := utils.ReadFileText(inputFile)
		if err != nil {
			fmt.Println("Error reading input file: ", err)
			os.Exit(1)
		}
		outputText := ShortenText(inputText)
		utils.WriteTextToFile(outputFile, outputText)
		fmt.Println("Tweet created successfully, saved to: ", outputFile)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tweetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tweetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.AddCommand(tweetCmd)

	tweetCmd.Flags().StringP("input", "i", "", "Input file name")
	tweetCmd.Flags().StringP("output", "o", "", "Output file name")

}

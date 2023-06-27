package cmd

import (
	"testing"

	"github.com/sleepypioneer/community-assistant-cli/utils"
)

func TestConvert(t *testing.T) {
	t.Run("Test successful tweet created", func(t *testing.T) {
		RootCmd.SetArgs([]string{"tweet", "-i", "../test_input.txt", "-o", "../output.txt"})
		err := RootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		outPuttedText, err := utils.ReadFileText("../output.txt")
		if err != nil {
			t.Fatal(err)
		}
		// check length of outputted text]
		expectedOutputLength := 240
		lengthOutputtedText := len(outPuttedText)
		if lengthOutputtedText > expectedOutputLength {
			t.Fatalf("expected output length to be lower than %d, got a length of %d", expectedOutputLength, lengthOutputtedText)
		}
	})

	t.Run("Test no input file", func(t *testing.T) {
		RootCmd.SetArgs([]string{"tweet", "-o", "../output.txt"})
		err := RootCmd.Execute()
		if err == nil {
			t.Fatalf("expected err (no input file) but got nil")
		}
	})

	t.Run("Test no output file", func(t *testing.T) {
		RootCmd.SetArgs([]string{"tweet", "-i", "../test_input.txt"})
		err := RootCmd.Execute()
		if err == nil {
			t.Fatalf("expected err (no output file) but got nil")
		}
	})

}

package world

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// promptSelection asks the user to select from a slice and returns the selected item.
func promptSelection(label string, items []string) (selection string, err error) {
	fmt.Println() // not sure why the prompt swallows the previous line
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, selection, err = prompt.Run()

	return selection, err
}

// promptString asks the user to enter a string and returns it.
func promptString(label string, validation func(i string) error) (input string, err error) {
	fmt.Println() // not sure why the prompt swallows the previous line
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			return nil
		},
	}

	return prompt.Run()
}

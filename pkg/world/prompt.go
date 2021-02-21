package world

import (
	"github.com/manifoldco/promptui"
)

// promptSelection asks the user to select from a slice and returns the selected item.
func promptSelection(label string, items []string) (selection string, err error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, selection, err = prompt.Run()
	return selection, err
}

// promptString asks the user to enter a string and returns it.
func promptString(label string, validation func(i string) error) (input string, err error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			return nil
		},
	}
	return prompt.Run()
}

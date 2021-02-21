package world

import (
	"github.com/chzyer/readline"
	"github.com/tigergraph/promptui"
)

// promptSelection asks the user to select from a slice and returns the selected item.
func promptSelection(label string, items []string) (selection string, err error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
		Keys: &promptui.SelectKeys{
			Next: promptui.Key{
				Code:    's',
				Display: "s",
			},
			Prev: promptui.Key{
				Code:    'w',
				Display: "w",
			},
			PageUp: promptui.Key{
				Code:    readline.CharBackward,
				Display: "←",
			},
			PageDown: promptui.Key{
				Code:    readline.CharForward,
				Display: "→",
			},
			Search: promptui.Key{Code: '/', Display: "/"},
		},
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

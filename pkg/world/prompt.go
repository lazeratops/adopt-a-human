package world

import (
	"github.com/AlecAivazis/survey/v2"
)

func promptSelection(label string, items []string) (string, error) {
	prompt := &survey.Select{
		Message: label,
		Options: items,
	}

	var res string
	if err := survey.AskOne(prompt, &res); err != nil {
		return "", err
	}
	return res, nil
}

// promptString asks the user to enter a string and returns it.
func promptString(label string, validation func(i string) error) (input string, err error) {
	prompt := survey.Input{
		Message: label,
	}
	var res string
	if err := survey.AskOne(&prompt, &res); err != nil {
		return "", err
	}
	return res, nil
}

package util

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"io"
	"os"
)

func PromptForMissingString(field *string, prompt string, help string, secret bool) error {
	if *field == "" {
		var err error
		*field, err = PickValue(prompt, "", help, secret)
		if err != nil {
			return err
		}
	}
	return nil
}
// TODO add validator
func PromptForMissingInt(field *int, prompt string, help string, secret bool) error {
	var err error
	*field, err = PickValue(prompt, "", help, secret)
	if err != nil {
		return err
	}
	return nil
}

func Pick(message string, names []string, defaultChoice string) (string, error) {
	if len(names) == 0 {
		return "", nil
	}
	if len(names) == 1 {
		return names[0], nil
	}
	name := ""
	prompt := &survey.Select{
		Message: message,
		Options: names,
		Default: defaultChoice,
	}

	surveyOpts := survey.WithStdio(os.Stdin, os.Stdout, os.Stderr)
	err := survey.AskOne(prompt, &name, nil, surveyOpts)
	return name, err
}

func PickValue(message string, defaultChoice string, help string, secret bool) (string, error) {
	if secret {
		return PromptValuePassword(message, help)
	} else {
		return PromptValue(message, defaultChoice, help)
	}
}

func PromptValue(message string, defaultChoice string, help string) (string, error) {
	name := ""
	prompt := &survey.Input{
		Message: message,
		Default: defaultChoice,
		Help:    help,
	}

	surveyOpts := survey.WithStdio(os.Stdin, os.Stdout, os.Stderr)
	err := survey.AskOne(prompt, &name, nil, surveyOpts)
	return name, err
}
func PromptValuePassword(message string, help string) (string, error) {
	name := ""
	prompt := &survey.Password{
		Message: message,
		Help:    help,
	}

	surveyOpts := survey.WithStdio(os.Stdin, os.Stdout, os.Stderr)
	err := survey.AskOne(prompt, &name, nil, surveyOpts)
	return name, err
}

// PickValue gets an answer to a prompt from a user's free-form input
func PickValueFromPath(message string, defaultValue string, required bool, help string, in terminal.FileReader, out terminal.FileWriter, outErr io.Writer) (string, error) {
	answer := ""
	prompt := &survey.Input{
		Message: message,
		Default: defaultValue,
		Help:    help,
	}
	validator := survey.Required
	if !required {
		validator = nil
	}
	surveyOpts := survey.WithStdio(in, out, outErr)

	err := survey.AskOne(prompt, &answer, validator, surveyOpts)
	if err != nil {
		return "", err
	}
	return answer, nil
}

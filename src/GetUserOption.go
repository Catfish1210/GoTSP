package Gotsp

import (
	"github.com/AlecAivazis/survey/v2"
)

func GetUserOption(header string, options []string) string {
	Uinput := ""
	prompt := &survey.Select{
		Message:       header,
		Options:       options,
		FilterMessage: "",
		VimMode:       true,
		Filter:        nil,
		Description:   nil,
	}

	survey.AskOne(prompt, &Uinput, nil, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = "!"
		icons.Question.Format = "red+b"
	}))

	return Uinput
}

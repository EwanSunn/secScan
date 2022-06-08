package cmd

import (
	"fmt"
	cmdCrack "github.com/EwanSunn/secScan/cmd/crack"
	cmdScan "github.com/EwanSunn/secScan/cmd/scan"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var version = 1

var App = grumble.New(&grumble.Config{
	Name:                  "secScan",
	Description:           fmt.Sprintf("An interactive penetration tool written in GO (version: %v)", version),
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

func init() {
	App.AddCommand(cmdScan.PortScan)
	App.AddCommand(cmdCrack.Crack)
}

func Run() {
	grumble.Main(App)
}

package main

import (
	"github.com/conplementag/cops-hq/v2/pkg/hq"
	"github.com/spf13/viper"
	"tado/internal/cli_flags"
	"tado/internal/server"
)

func createCommands(hq hq.HQ) {
	openWindowCommand := hq.GetCli().AddBaseCommand("open-window", "open window commands", "checks for open windows and deactivates heating", func() {
		server.Start(viper.GetViper())
	})
	openWindowCommand.AddParameterString(cli_flags.Homename, "", true, "n", "Name of the tado home to enable")
	openWindowCommand.AddParameterString(cli_flags.Username, "", true, "u", "Username of the tado home account")
	openWindowCommand.AddParameterString(cli_flags.Password, "", true, "p", "Password of the tado home account")
}

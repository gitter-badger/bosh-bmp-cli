package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"runtime/debug"

	flags "github.com/jessevdk/go-flags"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var options cmds.Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	args, err := parser.ParseArgs(os.Args)
	if err != nil {
		handlePanic()
		os.Exit(1)
	}

	command, err := createCommand(args, options)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rc, err := command.Execute(args)
	if err != nil {
		handlePanic()
		os.Exit(rc)
	}

	os.Exit(rc)
}

func createCommand(args []string, options cmds.Options) (cmds.Command, error) {
	if len(args) < 2 {
		return nil, errors.New("No bmp command specified")
	}

	var command cmds.Command

	cmdName := args[1]
	switch cmdName {
	case "bms":
		command = cmds.NewBmsCommand(args, options)
		break

	case "create-baremetals":
		command = cmds.NewCreateBaremetalsCommand(args, options)
		break

	case "login":
		command = cmds.NewLoginCommand(args, options)
		break

	case "sl-package-options":
		command = cmds.NewSlPackageOptionsCommand(args, options)
		break

	case "sl-packages":
		command = cmds.NewSlPackagesCommand(args, options)
		break

	case "status":
		command = cmds.NewStatusCommand(args, options)
		break

	case "stemcells":
		command = cmds.NewStemcellsCommand(args, options)
		break

	case "target":
		command = cmds.NewTargetCommand(args, options)
		break

	case "task":
		command = cmds.NewTaskCommand(args, options)
		break

	case "tasks":
		command = cmds.NewTasksCommand(args, options)
		break

	case "update-state":
		command = cmds.NewUpdateStateCommand(args, options)
		break

	default:
		return nil, errors.New(fmt.Sprintf("Unknown bmp command: %s", cmdName))
	}

	return command, nil
}

func handlePanic() {
	err := recover()

	if err != nil {
		switch err := err.(type) {
		case error:
			displayCrashDialog(err.Error())
		case string:
			displayCrashDialog(err)
		default:
			displayCrashDialog("An unexpected type of error")
		}
	}

	if err != nil {
		os.Exit(1)
	}
}

func displayCrashDialog(errorMessage string) {
	formattedString := `
Something completely unexpected happened. This is a bug in %s.
Please file this bug : https://github.com/maximilien/bosh-bmp-cli/issues
Tell us that you ran this command:

	%s

this error occurred:

	%s

and this stack trace:

%s
	`

	stackTrace := "\t" + strings.Replace(string(debug.Stack()), "\n", "\n\t", -1)
	println(fmt.Sprintf(formattedString, "bosh-bmp-cli", strings.Join(os.Args, " "), errorMessage, stackTrace))
}

package cmdsApp

import "fmt"

func RunCliCmdApp(input []string) error {
	subCmd := input[1]
	args := input[1:]

	if subCmd == "new" {
		return RunCliCmdAppNew(args)
	}
	if subCmd == "list" {
		return RunCliCmdAppList(args)
	}
	if subCmd == "create" {
		return RunCliCmdAppCreate(args)
	}
	if subCmd == "delete" {
		return RunCliCmdAppDelete(args)
	}
	return fmt.Errorf("unknown sub command [%s]", subCmd)
}

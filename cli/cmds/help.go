package cmds

import "fmt"

func RunCliCmdHelp(args []string) {
	fmt.Print(helpText)
}

const helpText = `
OC-APM is a helpful tool to manage applications on an Openshift-Cluster.

Usage:
	oc-apm (namespace) [command] [arguments]

Commands:
	list			Lists all items from openshift namespace.
	list <kinds>		All kinds to list. 
	list <kinds> <selector>	Sets LabelSelector to list.
	
	app		Manage your applications.
	app list	List all items found on disk optional filtered by name: list <name>
	app create	Creates all items from app on cluster: create <appName>
	app delete	Deletes all items from app on cluster: delete <appName>
	app new		Create new template app or load from cluster: new <name>

	start		Starts local web server UI(default :8082): start <port>	

	install		Self install APM on cluster.
	remove		Self remove APM from cluster.
	help		Print help message.

Global:
	namespace (ns)
			Required Namespace to operate: namespace=default | ns=default
`

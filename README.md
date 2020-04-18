# OC-APM - Openshift Application Manager

Bring your application to the Openshift cloud. 
Describe what your app needs and deploy it to the cluster with only one click. Perfect for 
testing new PoC's and also managing large applications with many Openshift items.


## Install
tbd
```

```

## Help

Start UI with:
```bash
$ oc-apm ns=default start
```

Command-Line usage:

```
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
```



## Types

The following Openshift-Item-Types are supported:
 - DeploymentConfig
 - StatefulSet
 - Service
 - Route
 - ConfigMap
 - PersistentVolumeClaim
 - ServiceAccount
 - Role
 - RoleBinding
 - Pod*

(* read only)


## License

Openshift Application Manager(oc-apm) is licensed under the Apache License 2.0.


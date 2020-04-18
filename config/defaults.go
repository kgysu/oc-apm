package config

// NamespaceEnvVar
const defaultNamespaceEnvName = "NAMESPACE"

var namespaceEnvName = defaultNamespaceEnvName

func GetNamespaceEnvNameOrDefault() string {
	return namespaceEnvName
}
func SetNamespaceEnvName(name string) {
	namespaceEnvName = name
}

// Namespace
const defaultNamespace = "default"

var namespace = defaultNamespace

func GetNamespaceOrDefault() string {
	return namespace
}
func SetNamespace(ns string) {
	namespace = ns
}

// Port
const defaultPort = ":8080"

var port = defaultPort

func GetPortOrDefault() string {
	return port
}
func SetPort(name string) {
	port = name
}

// DefaultRootFolder
const defaultRootFolder = "/apps"

var rootFolder = defaultRootFolder

func GetRootFolderOrDefault() string {
	return rootFolder
}

func SetRootFolder(folder string) {
	rootFolder = folder
}

// InClusterEnvName
const defaultInClusterEnvName = "RUN_IN_CLUSTER"

var inClusterEnvName = defaultInClusterEnvName

func GetInClusterEnvNameOrDefault() string {
	return inClusterEnvName
}

func SetInClusterEnvName(name string) {
	inClusterEnvName = name
}

// InClusterMode
const defaultInClusterMode = false

var inClusterMode = defaultInClusterMode

func IsInClusterModeOrDefault() bool {
	return inClusterMode
}

func SetInClusterMode(im bool) {
	inClusterMode = im
}

// DebugMode
const defaultDebugMode = false

var debugMode = defaultDebugMode

func IsInDebugModeOrDefault() bool {
	return debugMode
}

func SetDebugMode(dm bool) {
	debugMode = dm
}

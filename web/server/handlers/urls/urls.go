package urls

const GitPageUrl = "/git"
const LoadFromGitUrl = "/git/load"

const CmdAppStartUrl = "/cmd/app/start/"
const CmdAppStopUrl = "/cmd/app/stop/"
const CmdAppCreateUrl = "/cmd/app/create/"
const CmdAppUpdateUrl = "/cmd/app/update/"
const CmdAppNewUrl = "/cmd/app/new/"
const CmdAppDeleteUrl = "/cmd/app/delete/"

const AppListPageUrl = "/apps"
const AppNewPageUrl = "/new/app/"
const AppDetailsPageUrl = "/app/"

const OcCmdList = "/oc/cmd/list/"

const CmdAppItemStartUrl = "/cmd/app/item/start/"
const CmdAppItemStopUrl = "/cmd/app/item/stop/"
const CmdAppItemCreateUrl = "/cmd/app/item/create/"
const CmdAppItemUpdateUrl = "/cmd/app/item/update/"
const CmdAppItemDeleteUrl = "/cmd/app/item/delete/"
const AppItemDetailsPageUrl = "/appitem/"

func CreateCmdAppItemCreateUrl(appName, itemKind, itemName string) string {
	return CmdAppItemCreateUrl + appName + "/" + itemKind + "/" + itemName
}

func CreateCmdAppItemUpdateUrl(appName, itemKind, itemName string) string {
	return CmdAppItemUpdateUrl + appName + "/" + itemKind + "/" + itemName
}

func CreateCmdAppItemDeleteUrl(appName, itemKind, itemName string) string {
	return CmdAppItemDeleteUrl + appName + "/" + itemKind + "/" + itemName
}

func CreateAppItemDetailsPageUrl(appName, itemKind, itemName string) string {
	return AppItemDetailsPageUrl + appName + "/" + itemKind + "/" + itemName
}

func CreateCmdAppItemStartUrl(appName, itemKind, itemName string) string {
	return CmdAppItemStartUrl + appName + "/" + itemKind + "/" + itemName
}

func CreateCmdAppItemStopUrl(appName, itemKind, itemName string) string {
	return CmdAppItemStopUrl + appName + "/" + itemKind + "/" + itemName
}

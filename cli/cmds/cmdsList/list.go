package cmdsList

import (
	"fmt"
	"github.com/kgysu/oc-apm/client/util"
)

func RunCliCmdList(input []string) error {
	kinds := ""
	if len(input) > 1 {
		kinds = input[1]
	}
	labelSelector := ""
	if len(input) > 2 {
		labelSelector = input[2]
	}
	fmt.Printf("List: kinds=[%s] labelSelector=[%s]\n", kinds, labelSelector)
	appFromNamespace, err := util.GetAppFromNamespace("list", labelSelector)
	if err != nil {
		return err
	}

	for i, item := range appFromNamespace.GetItemsByKinds(kinds) {
		fmt.Printf("%d: %s \n", i, item.String())
	}
	return nil
}

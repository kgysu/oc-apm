package htmlTemplates

import "fmt"

func CreateBadge(class string, text string) string {
	return fmt.Sprintf(BadgeTemplate, class, text)
}

const BadgeTemplate = `<span class="%[1]s">%[2]s</span>`

func GetLabelsAsBadges(labels map[string]string) string {
	result := ""
	for labelKey, labelVal := range labels {
		result = result + CreateBadge("badge badge-secondary", labelKey+"="+labelVal) + " "
	}
	return result
}

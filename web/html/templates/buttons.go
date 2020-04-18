package htmlTemplates

import "fmt"

func CreateLinkButton(class, href, label string) string {
	return fmt.Sprintf(LinkButtonTemplate, class, href, label)
}

const LinkButtonTemplate = `<a class="%s" href="%s" role="button">%s</a>`

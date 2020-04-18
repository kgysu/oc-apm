package htmlTemplates

import "fmt"

// Headings
func CreateHeader(size int, title string) string {
	return fmt.Sprintf(HeadingTemplate, size, title)
}

const HeadingTemplate = `<h%[1]d>%[2]s</h%[1]d>`

// Breaks
func CreateBreaks(number int) string {
	var result string
	for i := 0; i < number; i++ {
		result = result + "<br />"
	}
	return result
}

func CreateDiv(class, content string) string {
	return fmt.Sprintf(DivTemplate, class, content)
}

const DivTemplate = `<div class="%s">%s<div>`

func CreateCodeBlock(content string) string {
	return fmt.Sprintf(CodeBlockTemplate, content)
}

const CodeBlockTemplate = `<pre><code>%s</code></pre>`

func CreateSampleOutput(content string) string {
	return fmt.Sprintf(SampleOutputTemplate, content)
}

const SampleOutputTemplate = `<samp>%s</samp>`

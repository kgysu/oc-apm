package htmlTemplates

import "fmt"

func CreateForm(action string, method string, button string, content string) string {
	return fmt.Sprintf(FormTemplate, action, method, button, content)
}

const FormTemplate = `<form action="%[1]s" method="%[2]s">
%[4]s
<input type="submit" class="btn btn-primary float-right" value="%[3]s">
</form>
`

func CreateInputFormGroup(class, id, label, content string) string {
	return fmt.Sprintf(FormGroupTemplate, class, id, label, content)
}

const FormGroupTemplate = `<div class="%[1]s">
<label for="%[2]s">%[3]s</label>
%[4]s
</div>
`

func CreateInputFormGroupRow(class, id, label, content string) string {
	return fmt.Sprintf(FormGroupRowTemplate, class, id, label, content)
}

const FormGroupRowTemplate = `<div class="%[1]s">
<label for="%[2]s" class="col-sm-2 col-form-label">%[3]s</label>
<div class="col-sm-10">
%[4]s
</div>
</div>
`

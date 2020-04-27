package htmlTemplates

import "fmt"

func CreateTextArea(class, id, name, rows, readOnly, value string) string {
	return fmt.Sprintf(TextAreaTemplate, class, id, name, rows, readOnly, value)
}

const TextAreaTemplate = `<textarea class="%[1]s" id="%[2]s" name="%[3]s" rows="%[4]s" %[5]s>
%[6]s
</textarea>
`

func CreateInput(inputType string, class string, id string, value string) string {
	return fmt.Sprintf(InputTemplate, inputType, class, id, value)
}

const InputTemplate = `<input type="%[1]s" class="%[2]s" id="%[3]s" name="%[3]s" value="%[4]s">`

func CreateInputDescription(id string, value string) string {
	return fmt.Sprintf(InputDescriptionTemplate, id, value)
}

const InputDescriptionTemplate = `<small id="%[1]s" class="form-text text-muted">%[2]s</small>`

func CreateCheckBox(id, label string) string {
	return fmt.Sprintf(CheckBoxTemplate, id, label)
}

const CheckBoxTemplate = `<div class="custom-control custom-checkbox">
  <input type="checkbox" class="custom-control-input" id="%[1]s" name="%[1]s">
  <label class="custom-control-label" for="%[1]s">%[2]s</label>
</div>`

const SelectTemplate = `<select class="form-control" id="exampleFormControlSelect1">
      <option>1</option>
      <option>2</option>
      <option>3</option>
      <option>4</option>
      <option>5</option>
    </select>
`

func CreateItemKindSelectTemplate() string {
	return ItemKindSelectTemplate
}

// TODO: add more kinds
const ItemKindSelectTemplate = `<select class="form-control" id="kind">
      <option>All</option>
      <option>DeploymentConfig</option>
      <option>Service</option>
      <option>Route</option>
      <option>Pod</option>
    </select>
`

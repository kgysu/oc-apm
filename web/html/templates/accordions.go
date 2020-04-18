package htmlTemplates

import "fmt"

func CreateAccordion(id, content string) string {
	return fmt.Sprintf(AccordionTemplate, id, content)
}

const AccordionTemplate = `<div class="accordion" id="%[1]s">
%[2]s
</div>
`

func CreateCard(id string, parentId string, name string, headerContent string, content string) string {
	return fmt.Sprintf(CardTemplate, "id"+id, name, headerContent, content, parentId)
}

const CardTemplate = `<div class="card">
    <div class="card-header" id="%[1]s">
      <h2 class="mb-0">
        <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#%[2]s" aria-expanded="false" aria-controls="%[2]s">
          %[3]s
        </button>
      </h2>
    </div>
    <div id="%[2]s" class="collapse" aria-labelledby="%[2]s" data-parent="#%[5]s">
      <div class="card-body">
        %[4]s
      </div>
    </div>
  </div>
`

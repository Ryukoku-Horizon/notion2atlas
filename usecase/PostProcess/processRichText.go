package postprocess

import (
	"notion2atlas/domain"
	"strings"
)

func processRichText(richTexts []domain.RichTextEntity, pageEntities []domain.PageEntity, cId string) []domain.AtlRichTextEntity {
	atlParents := []domain.AtlRichTextEntity{}
	for _, item := range richTexts {
		if item.Href == nil {
			atlParents = append(atlParents, item.ToAtlEntity(false))
			continue
		}
		href := *item.Href
		if strings.HasPrefix(href, "/posts/curriculums/") {
			params := strings.Split(href, "/")
			pageId := params[3]
			var curriculumId string = ""
			for _, page := range pageEntities {
				if page.Id == pageId {
					curriculumId = page.CurriculumId
					break
				}
			}
			if curriculumId == "" {
				atlParents = append(atlParents, item.ToAtlEntity(false))
				continue
			}
			isSameCurriculum := curriculumId == cId
			atlParents = append(atlParents, item.ToAtlEntity(isSameCurriculum))
			continue
		}
		atlParents = append(atlParents, item.ToAtlEntity(false))
	}
	return atlParents
}

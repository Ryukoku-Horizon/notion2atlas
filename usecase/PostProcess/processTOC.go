package postprocess

import "notion2atlas/domain"

func processTOC(atlDataArr []domain.BlockEntity) domain.AtlBlockEntityData {
	headerInfo := []domain.HeaderInfo{}
	for _, block := range atlDataArr {
		switch block.Type {
		case "heading_1":
			text := block.Data.Header.GetCombinedPlainText()
			headerInfo = append(headerInfo, domain.HeaderInfo{HeaderType: 1, BlockId: block.Id, Text: text})
		case "heading_2":
			text := block.Data.Header.GetCombinedPlainText()
			headerInfo = append(headerInfo, domain.HeaderInfo{HeaderType: 2, BlockId: block.Id, Text: text})
		case "heading_3":
			text := block.Data.Header.GetCombinedPlainText()
			headerInfo = append(headerInfo, domain.HeaderInfo{HeaderType: 3, BlockId: block.Id, Text: text})
		}
	}
	toc := domain.AtlTableOfContentsEntity(headerInfo)
	data := domain.AtlBlockEntityData{
		Type:            "table_of_contents",
		TableOfContents: &toc,
	}
	return data
}

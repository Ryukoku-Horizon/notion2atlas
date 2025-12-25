package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/gateway"
)

func addOgpDataToPage() error {
	tmpPages, err := filemanager.ReadJson[[]domain.PageEntity](constants.TMP_PAGE_PATH)
	if err != nil {
		fmt.Println("error in postprocess/addOgpDataToPage/filemanager.ReadJson")
		return err
	}
	atlPageEntities := []domain.AtlPageEntity{}
	for _, page := range tmpPages {
		var filepath = fmt.Sprintf("%s/%s.json", constants.TMP_DIR, page.Id)
		blocks, err := filemanager.ReadJson[[]domain.BlockEntity](filepath)
		if err != nil {
			fmt.Println("error in postprocess/addOgpDataToPage/filemanager.ReadJson")
			return err
		}
		firstText := ""
		for _, block := range blocks {
			switch block.Data.Type {
			case "paragraph":
				firstText = block.Data.Paragraph.GetConcatenatedText()
			case "todo":
				firstText = block.Data.Todo.GetConcatenatedText()
			case "header":
				firstText = block.Data.Header.GetCombinedPlainText()
			case "callout":
				filepath = block.Data.Callout.GetConcatenatedText()
			}
			if firstText != "" {
				break
			}
		}
		imagePath := fmt.Sprintf("%s/ogp/%s.png", constants.DEPLOY_URL, page.Id)
		ogpData := domain.PageOgp{
			FirstText: firstText,
			ImagePath: imagePath,
		}
		atlPageEntity := domain.NewAtlPageEntity(
			page.Id,
			page.CurriculumId,
			page.IconType,
			page.IconUrl,
			page.CoverUrl,
			page.CoverType,
			page.Order,
			page.ParentId,
			page.Title,
			page.Type,
			ogpData,
		)
		atlPageEntities = append(atlPageEntities, atlPageEntity)
	}
	err = gateway.UpsertFile(domain.PAGE, "id", atlPageEntities)
	if err != nil {
		fmt.Println("error in postprocess/addOgpDataToPage/gateway.UpsertFile")
		return err
	}
	err = filemanager.DelFile(constants.TMP_PAGE_PATH)
	if err != nil {
		fmt.Println("error in postprocess/addOgpDataToPage/filemanager.DelFile")
		return err
	}
	return nil
}

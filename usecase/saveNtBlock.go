package usecase

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
)

func saveNtBlockInPage(bp domain.BasePage, pageBuffer []domain.PageEntity, resourceType string) ([]domain.PageEntity, error) {
	filemanager.CreateDirIfNotExist(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, bp.GetId()))
	err := filemanager.ClearDir(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, bp.GetId()))
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/filemanager.ClearDir in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	urls, pageEntity, err := saveBasePage(bp)
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/saveBasePage in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	basePage := *pageEntity
	urlRewritedEntity, err := basePage.ChangePageEntityUrl(urls.IconUrl, urls.CoverUrl)
	if err != nil {
		fmt.Println("error in usecase/saveNtBlockInPage/basePage.ChangePageEntityUrl")
		return pageBuffer, err
	}
	pageBuffer = append(pageBuffer, *urlRewritedEntity)
	blocks, err := GetChildren(bp.GetId())
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/GetChildren in curriculum/" + bp.GetTitle())
		return nil, err
	}
	var blockBuffer = []domain.BlockEntity{}
	blockBuffer, pageBuffer, err = saveNtChildrenBlocks(bp.GetId(), bp.GetId(), blocks, blockBuffer, pageBuffer, resourceType)
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/saveNtChildrenBlocks in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	err = FlushBlockBuffer(blockBuffer, bp.GetId())
	if err != nil {
		fmt.Println("error in usecase/InsertCurriculumBlocks/FlushBlockBuffer in curriculum/" + bp.GetTitle())
		return pageBuffer, err
	}
	return pageBuffer, nil
}

func saveNtChildrenBlocks(currId string, pageId string, blocks []domain.NTBlockEntity, blockBuffer []domain.BlockEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.BlockEntity, []domain.PageEntity, error) {
	var err error = nil
	for i, block := range blocks {
		blockBuffer, pageBuffer, err = saveNtBlock(block, currId, pageId, i, fmt.Sprintf("%d/%d", i+1, len(blocks)), blockBuffer, pageBuffer, resourceType)
		if err != nil {
			fmt.Println("error in usecase/saveNtChildrenBlocks/saveNtBlock")
			return blockBuffer, pageBuffer, err
		}
	}
	return blockBuffer, pageBuffer, nil
}

func saveNtBlock(block domain.NTBlockEntity, curriculumId string, pageId string, i int, p string, buffer []domain.BlockEntity, pageBuffer []domain.PageEntity, resourceType string) ([]domain.BlockEntity, []domain.PageEntity, error) {
	type_ := block.Type
	fmt.Println(p, type_)
	var err error = nil
	buffer, pageBuffer, err = GetBlockEntities(block, buffer, curriculumId, pageId, i, pageBuffer, resourceType)
	if err != nil {
		fmt.Println("error in usecase/saveNtBlock/GetBlockEntities in " + type_)
		return buffer, pageBuffer, err
	}
	if type_ == "child_page" {
		filemanager.CreateDirIfNotExist(fmt.Sprintf("%s/%s", constants.ASSETS_DIR, block.Id))
		children, err := GetChildren(block.Id)
		if err != nil {
			fmt.Println("error in usecase/saveNtBlock/GetChildren in " + type_)
			return buffer, pageBuffer, err
		}
		var newBuffer = []domain.BlockEntity{}
		newBuffer, pageBuffer, err = saveNtChildrenBlocks(curriculumId, block.Id, children, newBuffer, pageBuffer, resourceType)
		if err != nil {
			fmt.Println("error in usecase/saveNtBlock/saveNtChildrenBlocks in " + type_)
			return buffer, pageBuffer, err
		}
		err = FlushBlockBuffer(newBuffer, block.Id)
		if err != nil {
			fmt.Println("error in usecase/saveNtBlock/FlushBlockBuffer in " + type_)
			return buffer, pageBuffer, err
		}
	} else if (type_ == "synced_block" && block.SyncedBlock.SyncedFrom == nil) || type_ != "synced_block" {
		if block.HasChildren {
			children, err := GetChildren(block.Id)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/GetChildren in " + type_)
				return buffer, pageBuffer, err
			}
			buffer, pageBuffer, err = saveNtChildrenBlocks(curriculumId, pageId, children, buffer, pageBuffer, resourceType)
			if err != nil {
				fmt.Println("error in usecase/saveNtBlock/saveNtChildrenBlocks in " + type_)
				return buffer, pageBuffer, err
			}
		}
	}
	return buffer, pageBuffer, nil
}

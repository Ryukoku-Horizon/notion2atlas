package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func saveNtData[T domain.BasePage](item domain.Entity, resourceType domain.ResourceType) error {
	var err error = nil
	curr := item.(T)
	pageBuffer := []domain.PageEntity{}
	pageBuffer, err = saveNtBlockInPage(curr, pageBuffer, resourceType.GetStr())
	if err != nil {
		fmt.Println("error in usecase/saveNtData/InsertCurriculumBlocks New")
		return err
	}
	err = FlushPageBuffer(pageBuffer, curr.GetId())
	if err != nil {
		fmt.Println("error in usecase/saveNtData/FlushPageBuffer")
		return err
	}
	err = UpsertBasePage(curr.GetId(), curr, resourceType)
	if err != nil {
		fmt.Println("error in usecase/saveNtData/UpsertCurriculum")
		return err
	}
	fmt.Println("✅ complete read : " + curr.GetTitle())
	return nil
}

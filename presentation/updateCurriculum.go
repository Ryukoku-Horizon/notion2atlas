package presentation

import (
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/usecase"
	"os"
)

func updateCurriculum() (*usecase.NDE, error) {
	db_id := os.Getenv("NOTION_DB_ID_HORIZON")
	publishedRecords, err := usecase.GetDBQuery(db_id)
	if err != nil {
		fmt.Println("error in presentation/saveCurriculum/usecase.GetDBQuery")
		return nil, err
	}
	var curriculums []domain.CurriculumEntity
	for _, query := range publishedRecords {
		curr, err := query.ToCurriculumEntity()
		if err != nil {
			fmt.Println("error in presentation/saveCurriculum/converter.Query2CurriculumEntity")
			return nil, err
		}
		if curr == nil {
			return nil, fmt.Errorf("unexpected: curriculumEntity is nil")
		}
		curriculums = append(curriculums, *curr)
	}
	oldDataAddress, err := usecase.GetCurriculumFile()
	if err != nil {
		fmt.Println("error in presentation/saveCurriculum/usecase.GetCurriculumFile")
		return nil, err
	}
	if oldDataAddress == nil {
		return nil, fmt.Errorf("oldDataAddress is nil")
	}
	oldData := *oldDataAddress
	nde, err := usecase.ProcessNTData[domain.CurriculumEntity, domain.CurriculumEntity](oldData, curriculums, domain.CURRICULUM)
	if err != nil {
		fmt.Println("error in presentation/updateCurriculum/usecase.ProcessNTData")
		return nil, err
	}

	fmt.Println("✅ completed: update curriculums")
	return nde, nil
}

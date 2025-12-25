package usecase

import (
	"fmt"
	"notion2atlas/domain"
)

func saveBasePage(bp domain.BasePage) (*URLModel, *domain.PageEntity, error) {
	pageEntity, err := bp.ToPageEntity()
	if err != nil {
		fmt.Println("error in usecase/saveBasePage/bp.ToPageEntity")
		return nil, nil, err
	}
	err = DownloadPageImg(pageEntity)
	if err != nil {
		fmt.Println("error in usecase/saveBasePage/DownloadPageImg")
		return nil, nil, err
	}
	urls := GetPathRewritedUrl(pageEntity)
	return &urls, pageEntity, nil
}

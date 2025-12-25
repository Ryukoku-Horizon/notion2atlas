package domain

import (
	"fmt"
	"notion2atlas/constants"
)

type ResourceType int

const (
	Resources ResourceType = iota
	CURRICULUM
	PAGE
	CATEGORY
	INFO
	ANSWER
	SYNCED
	TMP_PAGE
)

func (r ResourceType) GetStr() string {
	switch r {
	case CURRICULUM:
		return "curriculum"
	case PAGE:
		return "page"
	case CATEGORY:
		return "category"
	case INFO:
		return "info"
	case ANSWER:
		return "answer"
	case SYNCED:
		return "synced"
	default:
		return ""
	}
}

func (r ResourceType) GetFilePathFromResourceType() (string, error) {
	switch r {
	case CURRICULUM:
		return constants.CURRICULUM_PATH, nil
	case PAGE:
		return constants.PAGE_PATH, nil
	case CATEGORY:
		return constants.CATEGORY_PATH, nil
	case INFO:
		return constants.INFO_PATH, nil
	case ANSWER:
		return constants.ANSWER_PATH, nil
	case SYNCED:
		return constants.SYNCED_PATH, nil
	case TMP_PAGE:
		return constants.TMP_PAGE_PATH, nil
	default:
		return "", fmt.Errorf("unexpected resourceType")
	}
}

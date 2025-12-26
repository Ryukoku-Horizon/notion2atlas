package domain

type AtlPageEntity struct {
	Id             string   `json:"id"`
	CurriculumId   string   `json:"curriculumId"`
	IconType       string   `json:"iconType"`
	IconUrl        string   `json:"iconUrl"`
	CoverUrl       string   `json:"coverUrl"`
	CoverType      string   `json:"coverType"`
	Order          int      `json:"order"`
	ParentId       string   `json:"parentId"`
	Title          string   `json:"title"`
	Type           string   `json:"type"`
	Ogp            PageOgp  `json:"ogp"`
	Visibility     []string `json:"visibility"`
	Tag            []string `json:"tag"`
	Category       []string `json:"category"`
	LastEditedTime string   `json:"last_edited_time"`
}

func (p AtlPageEntity) GetId() string {
	return p.Id
}

func NewAtlPageEntity(
	Id string,
	CurriculumId string,
	IconType string,
	IconUrl string,
	CoverUrl string,
	CoverType string,
	Order int,
	ParentId string,
	Title string,
	Type string,
	Ogp PageOgp,
	Visibility []string,
	Tag []string,
	Category []string,
	LastEditedTime string,
) AtlPageEntity {
	return AtlPageEntity{
		Id:             Id,
		CurriculumId:   CurriculumId,
		IconType:       IconType,
		IconUrl:        IconUrl,
		CoverType:      CoverType,
		CoverUrl:       CoverUrl,
		Order:          Order,
		ParentId:       ParentId,
		Title:          Title,
		Type:           Type,
		Ogp:            Ogp,
		Visibility:     Visibility,
		Tag:            Tag,
		Category:       Category,
		LastEditedTime: LastEditedTime,
	}
}

type PageOgp struct {
	FirstText string `json:"first_text"`
	ImagePath string `json:"image_path"`
}

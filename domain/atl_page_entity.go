package domain

type AtlPageEntity struct {
	Id           string  `json:"id"`
	CurriculumId string  `json:"curriculumId"`
	IconType     string  `json:"iconType"`
	IconUrl      string  `json:"iconUrl"`
	CoverUrl     string  `json:"cover"`
	CoverType    string  `json:"coverType"`
	Order        int     `json:"order"`
	ParentId     string  `json:"parentId"`
	Title        string  `json:"title"`
	Type         string  `json:"type"`
	Ogp          PageOgp `json:"ogp"`
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
) AtlPageEntity {
	return AtlPageEntity{
		Id,
		CurriculumId,
		IconType,
		IconUrl,
		CoverType,
		CoverUrl,
		Order,
		ParentId,
		Title,
		Type,
		Ogp,
	}
}

type PageOgp struct {
	FirstText string `json:"first_text"`
	ImagePath string `json:"image_path"`
}

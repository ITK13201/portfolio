package domain

type AboutTopic struct {
	Title         string `json:"title,omitempty"`
	DescriptionJp string `json:"description-jp,omitempty"`
	DescriptionEn string `json:"description-en,omitempty"`
	ImagePath     string `json:"image-path,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}

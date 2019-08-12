package model

type ImageInfo struct {
	ImageURL      string `json:"image_url"`
	ImageType     string `json:"image_type"`
	Text          string `json:"text"`
	WebPageURL    string `json:"web_page_url"`
	ImageThumbURL string `json:"thumb_url"`
}

package models

type BannerModel struct {
	Model
	Title string `json:"title"` // 标题
	Cover string `json:"cover"` // 封面url
	Url   string `json:"url"`   // 跳转url
}

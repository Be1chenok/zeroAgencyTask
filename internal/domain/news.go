package domain

//go:generate reform

//reform:news
type News struct {
	Id         int    `json:"Id" reform:"id,pk"`
	Title      string `json:"Title" reform:"title"`
	Content    string `json:"Content" reform:"content"`
	Categories []int  `json:"Categories" reform:"-"`
}

//reform:news_categories
type NewsCategories struct {
	Id         int `reform:"id,pk"`
	NewsId     int `reform:"news_id"`
	CategoryId int `reform:"category_id"`
}

type NewsSearchParams struct {
	Offset int
	Limit  int
}

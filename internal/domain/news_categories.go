package domain

//go:generate reform

//reform:news_categories
type NewsCategories struct {
	Id         int `reform:"id,pk"`
	NewsId     int `reform:"news_id"`
	CategoryId int `reform:"category_id"`
}

package domain

//go:generate reform

//reform:news
type News struct {
	Id         int    `json:"Id" reform:"id,pk"`
	Title      string `json:"Title" reform:"title"`
	Content    string `json:"Content" reform:"content"`
	Categories []int  `json:"Categories" reform:"-"`
}

type NewsSearchParams struct {
	Offset int
	Limit  int
}

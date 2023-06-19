package schemas

type ListArticleIn struct {
	//CategoryId string `uri:"category_id"` //只能用ShouldBindUri获取
	PageQuery
	Keyword string `json:"keyword" form:"keyword"`
}

type List2ArticleIn struct {
	PageQuery
	CategoryIds string `form:"category_ids"`
	Keyword     string `json:"keyword" form:"keyword"`
}

type CreateArticleIn struct {
	Title   string `json:"title" form:"title" binding:"required,min=3,max=75"`
	Content string `json:"content" form:"content" binding:"required"`
	//CategoryId  string `json:"category_id" uri:"category_id"`
	Keywords    string `json:"keywords" form:"keywords"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	Images      string `json:"images" form:"images"`
	Sort        int    `json:"sort" form:"sort"`
	Source      string `json:"source" form:"source"`
	Author      string `json:"author" form:"author"`
	Url         string `json:"url" form:"url"`
	Status      int    `json:"status" form:"status" binding:"oneof=0 1"`
}

type DetailArticleIn struct {
	CurrentArticle
}

type ChangeArticleHeaderIn struct {
	CurrentArticle
	CurrentCategory
}

type ChangeArticleIn struct {
	Title   string `json:"title" form:"title" binding:"omitempty,min=3,max=75"`
	Content string `json:"content" form:"content"`
	//CategoryId  string `json:"category_id" uri:"category_id"`
	Keywords    string `json:"keywords" form:"keywords"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	Images      string `json:"images" form:"images"`
	Sort        int    `json:"sort" form:"sort"`
	Source      string `json:"source" form:"source"`
	Author      string `json:"author" form:"author"`
	Url         string `json:"url" form:"url"`
	Status      int    `json:"status" form:"status" binding:"oneof=0 1"`
}

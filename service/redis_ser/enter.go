package redis_ser

const (
	articleLookPrefix = "article_look"
	articleDiggPrefix = "article_digg"
)

// NewArticleLook 文章浏览量
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}

// NewArticleDigg 文章点赞
func NewArticleDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}

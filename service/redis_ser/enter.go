package redis_ser

const (
	articleLookPrefix = "article_look"
)

// NewArticleLook 文章浏览量
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}

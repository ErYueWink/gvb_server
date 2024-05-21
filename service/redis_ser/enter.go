package redis_ser

const (
	articleLookPrefix    = "article_look"
	articleDiggPrefix    = "article_digg"
	articleCommentPrefix = "article_comment"
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

// NewArticleComment 文章评论
func NewArticleComment() CountDB {
	return CountDB{
		Index: articleCommentPrefix,
	}
}

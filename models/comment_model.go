package models

// CommentModel 评论表
type CommentModel struct {
	MODEL `json:",select(c)"`

	SubComments        []CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(c)"`     // 子评论列表
	ParentCommentModel *CommentModel  `gorm:"foreignKey:ParentCommentID" json:"comment_model"`              // 父级评论
	ParentCommentID    *uint          `gorm:"comment:父评论id" json:"parent_comment_id,select(c)"`             // 父评论id
	Content            string         `gorm:"size:256;comment:评论内容" json:"content,select(c)"`               // 评论内容
	DiggCount          int            `gorm:"size:8;default:0;comment:点赞数" json:"digg_count,select(c)"`     // 点赞数
	CommentCount       int            `gorm:"size:8;default:0;comment:子评论数" json:"comment_count,select(c)"` // 子评论数
	ArticleID          string         `gorm:"comment:文章id;size:32" json:"article_id"`                       // 文章id
	User               UserModel      `gorm:"comment:关联的用户" json:"user,select(c)"`                          // 关联的用户
	UserID             uint           `gorm:"comment:关联的用户id" json:"user_id,select(c)"`                     // 评论的用户
}

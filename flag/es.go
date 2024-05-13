package flag

import "gvb_server/models"

// ESCreateIndex 生成表结构
func ESCreateIndex() {
	models.ArticleModel{}.CreateIndex()
	//models.FullTextModel{}.CreateIndex()
}

package core

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"time"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.Es.URL()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.Es.UserName, global.Config.Es.Password),
		elastic.SetHealthcheckTimeout(time.Minute),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	return c
}

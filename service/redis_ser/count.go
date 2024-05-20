package redis_ser

import "gvb_server/global"

type CountDB struct {
	Index string
}

// Set 设置某一个值，重复执行，重复累加
func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()
	return err
}

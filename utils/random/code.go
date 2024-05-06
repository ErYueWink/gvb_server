package random

import (
	"fmt"
	"math/rand"
	"time"
)

func Code() string {
	// 设置随机数种子，确保每次生成的随机数都是不一样的
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(10000))
}

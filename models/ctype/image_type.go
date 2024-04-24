package ctype

import "encoding/json"

type ImageType int

const (
	LOCAL ImageType = 1
	QINIU ImageType = 2
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case LOCAL:
		str = "本地"
	case QINIU:
		str = "七牛"
	default:
		str = "其他"
	}
	return str
}

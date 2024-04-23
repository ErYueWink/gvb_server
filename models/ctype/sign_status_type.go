package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ       SignStatus = 1
	SignGitee    SignStatus = 2
	SignEmail    SignStatus = 3
	SignGithub   SignStatus = 4
	SignNoPublic SignStatus = 5
)

func (s SignStatus) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "SignQQ"

	}
}

package types

import "github.com/leandroyou/http-car/singleton"

type APIkey string

func (key APIkey) IsValid() bool {
	return key == APIkey(singleton.MyAPIKey)
}

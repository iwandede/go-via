package lib

import (
	"github.com/rs/xid"
)

func GenerateID() string {
	guid := xid.New()
	return guid.String()
}

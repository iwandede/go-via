package lib

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Pagination struct {
	From         int64  `json:"from"`
	Limit        int64  `json:"limit"`
	BeforeCursor string `json:"before_cursor"`
	AfterCursor  string `json:"after_cursor"`
}

func decodeCursor(encodedCursor string) (res time.Time, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	res, err = time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return
	}
	uuid = arrStr[1]
	return
}

func encodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

package lib

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func Pagination(r *url.URL) (*PaginationDTO, error) {
	var (
		Limit  int64 = 10
		Offset int64 = 0
	)
	qs := r.Query()

	NextCursor := qs.Get("next_cursor")
	PrevCursor := qs.Get("next_cursor")

	if Limit > 0 {
		Offset = Offset + (Limit + 1)
		NextCursor = fmt.Sprintf("?limit=%d&offset=%d", Limit, Offset)
		PrevCursor = fmt.Sprintf("?limit=%d&offset=%d", Limit, (Offset - Limit))
	}

	return &PaginationDTO{
		BeforeCursor: EncodeCursor(NextCursor),
		AfterCursor:  EncodeCursor(PrevCursor),
	}, nil
}

func DecodeCursor(encodedCursor string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return "", err
	}

	return ToString(str), nil
}

func EncodeCursor(uuid string) string {
	key := fmt.Sprintf("%s", uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

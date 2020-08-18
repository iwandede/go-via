package lib

import "strconv"

func ToString(n interface{}, p ...int) string {
	var t string

	switch n.(type) {
	case bool:
		t = strconv.FormatBool(n.(bool))
	case int:
		t = strconv.Itoa(n.(int))
	case int64:
		t = strconv.FormatInt(n.(int64), 10)
	case float32:
		if len(p) > 0 {
			t = strconv.FormatFloat(float64(n.(float32)), 'f', p[0], 64)
		} else {
			t = strconv.FormatFloat(float64(n.(float32)), 'f', -1, 64)
		}
	case float64:
		if len(p) > 0 {
			t = strconv.FormatFloat(n.(float64), 'f', p[0], 64)
		} else {
			t = strconv.FormatFloat(n.(float64), 'f', -1, 64)
		}
	case byte:
		t = string(n.(byte))
	case []byte:
		t = string(n.([]byte))
	case string:
		t = n.(string)
	}

	return t
}

func ToInt64(params string) int64 {
	number, _ := strconv.ParseInt(params, 10, 64)
	return number
}

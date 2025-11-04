package utils

import "encoding/json"

func MarshalToStringNoError(v any) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

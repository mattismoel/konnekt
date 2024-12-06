package prnt

import "encoding/json"

func Pretty(v any) string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}

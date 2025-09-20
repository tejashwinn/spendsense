package util

import "fmt"

func ParseUint(s string) uint {
	var v uint
	fmt.Sscanf(s, "%d", &v)
	return v
}

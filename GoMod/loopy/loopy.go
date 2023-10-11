package loopy

import "strings"

func Loopit(hello_world string, no_required int) string {
	hello_list := []string{}
	n := 0
	for n < no_required {
		hello_list = append(hello_list, hello_world)
		n += 1
	}
	return strings.Join(hello_list, "\n")

}

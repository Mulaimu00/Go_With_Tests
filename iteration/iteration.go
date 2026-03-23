package iteration

import "strings"


func Repeat(n string, repeatCount int) string {
	var block strings.Builder

	for i:=0;i<repeatCount;i++{
		block.WriteString(n)
	}
	return block.String()
}

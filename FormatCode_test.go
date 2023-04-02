package csy_test

import (
	"fmt"
	"github.com/front-ck996/csy"
	"testing"
)

func TestFormatCode(t *testing.T) {
	in := []byte(`
package p
var ()
func f() {
	for _ = range v {
	}
}
`[1:])
	fmt.Println(csy.FormatCode(in))
}

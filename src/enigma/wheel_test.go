package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

/*
=== RUN   TestWheel
index 0: 1 -> B
--- PASS: TestWheel (0.00s)
*/
func TestWheel(t *testing.T) {
	w := enigma.NewWheel(enigma.CharacterSet)
	a := w.GetIndex('B')
	b := w.GetChar(a)
	fmt.Printf("index 0: %v -> %v\n", a, string(b))
}

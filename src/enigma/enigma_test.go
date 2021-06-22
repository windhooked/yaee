package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

func TestEnigmaM4(t *testing.T) {
	m4 := enigma.M4
	m4.Setting([]string{"I", "II", "III"},
		[]byte{'H', 'F', 'K', 'D'},
		enigma.PB_30)

	code := m4.Step('A')

	fmt.Printf("%v", string(code))
}

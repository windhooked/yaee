package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

func TestStaticWheel(t *testing.T) {
	wS := enigma.ETW_M4
	var result []byte
	for _, v := range enigma.ETW_M4.Lut() { //
		a := wS.Encode(v)
		fmt.Printf("enc %v -> %v\n", string(v), string(a))
		result = append(result, a)
	}

	if string(result) != string(wS.Lut()) {
		t.Fatalf("encoding for static wheel not correct")
	}

}

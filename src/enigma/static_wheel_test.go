package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

func TestStaticWheel(t *testing.T) {
	wS := enigma.ETW_M4
	var result []byte
	for _, v := range []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ") { //
		a := wS.GetIndex(v)
		b := wS.GetChar(a)
		fmt.Printf("enc %v -> %v, dec %v\n", string(v), string(a), string(b))
		result = append(result, a)
	}

	if string(result) != string(wS.Lut()) {
		t.Fatalf("encoding for static wheel not correct")
	}

}

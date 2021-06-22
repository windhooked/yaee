package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

func TestWheel(t *testing.T) {
	wI := enigma.W_I
	var result []byte
	for _, v := range enigma.ETW_M4.Lut() { //
		if v == 'Q' {
			fmt.Printf("transtion from Q - R")
		}
		a, notch := wI.Encode(v, true) // wheel one will always icrement, validate this assumtion?
		if v == 'Q' {
			if notch == false {
				t.Fatalf("bump not set after %v", string(v))
			}
		}
		if notch == true {
			fmt.Printf(" increment next\n")
		}
		fmt.Printf("enc %v -> %v\n", string(v), string(a))
		result = append(result, a)
	}

	if string(result) != string(wI.Lut()) {
		t.Fatalf("encoding for wheel I not correct")
	}

}

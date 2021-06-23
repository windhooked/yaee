package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

//This test should pass as it translates to a 1 to 1 mapping
func TestRotorWheel(t *testing.T) {
	wI := enigma.W_I
	wI.Build()
	wI.SetInner('Z')
	wI.SetRing('B') // step before A //Message Key
	var encoded []byte
	var decoded []byte
	//var testChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var testChars = []byte("AAAAAAAAAAAAAAAAAAAAAAAAAA")
	for k, v := range testChars {
		if k == 25 {
			//fmt.Printf("transtion from Q - R")
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
		//fmt.Printf("enc %v -> %v\n", string(v), string(a))
		encoded = append(encoded, a)

		z := wI.Decode(a)
		fmt.Printf("dec %v -> %v\n", string(a), string(z))
		decoded = append(decoded, z)
	}

	if string(encoded) != string(wI.Lut()) {
		//t.Fatalf("encoding for wheel I not correct %v %v", string(encoded), string(wI.Lut()))
	}
	t.Logf("encoded %v decoded %v", string(encoded), string(decoded))

}

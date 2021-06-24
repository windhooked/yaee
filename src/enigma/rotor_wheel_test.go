package enigma_test

import (
	"fmt"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

//This test should pass as it translates to a 1 to 1 mapping
func TestRotorWheel(t *testing.T) {
	wI := enigma.NewRotorWheel(1, enigma.W_I)
	wI.SetInnerOffset('A')
	wI.SetRingPosition('Z') // step before A //Message Key
	var encoded []byte
	var decoded []byte
	//var testChars = []byte("AAAAAAAAAAAAAAAAAAAAAAAAAA")
	var testChars = []byte(enigma.CharacterSet)
	for k, v := range testChars {
		if k == 25 {
			//fmt.Printf("transtion from Q - R")
		}
		notch := wI.Step(true)  // wheel one will always icrement
		a := wI.Encode(v - 'A') //
		if v == 'Q' {
			if notch == false {
				t.Fatalf("bump not set after %v\n", string(v))
			}
		}
		if notch == true {
			fmt.Printf(" increment next\n")
		}
		//fmt.Printf("enc %v -> %v\n", string(v), string(a))
		encoded = append(encoded, a+'A')

		z := wI.Decode(a)
		fmt.Printf("dec %v -> %v\n", string(a+'A'), string(z+'A'))
		decoded = append(decoded, z+'A')
	}

	if string(encoded) != string(wI.Lut()) {
		//t.Fatalf("encoding for wheel I not correct %v %v", string(encoded), string(wI.Lut()))
	}
	t.Logf("encoded %v decoded %v", string(encoded), string(decoded))

}

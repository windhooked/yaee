package enigma

import "fmt"

type (
	Enigma struct {
		Plugboard   SB
		StaticWheel ETW
		Wheels      []RotorWheel
		Reflector   UKW
	}
)

const (
	numChars = 26
)

var (
	M4 = &Enigma{
		StaticWheel: ETW_M4,
		Wheels:      []RotorWheel{W_I, W_II, W_III},
		Reflector:   UKW_A,
	}
)

func (h *Enigma) Setting(rotor []string, dial []byte, pb []PB) {
	//Wire plugboard
	h.Plugboard.WireUp(pb)

	//insert rotors

	h.Wheels = nil
	for k := range rotor {
		switch rotor[k] {
		case "I":
			h.Wheels = append(h.Wheels, W_I)
		case "II":
			h.Wheels = append(h.Wheels, W_II)
		case "III":
			h.Wheels = append(h.Wheels, W_III)
		case "IV":
			h.Wheels = append(h.Wheels, W_IV)
		case "V":
			h.Wheels = append(h.Wheels, W_V)
		case "VI":
			h.Wheels = append(h.Wheels, W_VI)
		case "VII":
			h.Wheels = append(h.Wheels, W_VII)
		case "VIII":
			h.Wheels = append(h.Wheels, W_III)
		default:
			fmt.Printf("failed to find config for rotor type %v", rotor[k])
		}
		h.Wheels[k].SetRing(dial[k])
	}

}

/*
keyboard - plugboard - static wheel - w0, w1, w2 - reflector

*/
func (h *Enigma) Step(i byte) (o byte) {

	//plugboard ->
	n := h.Plugboard.Encode(i)

	//Input Rotor ->
	n = h.StaticWheel.Encode(n)

	// Main Rotors ->
	next := true
	// always increment first wheel
	n, next = h.Wheels[0].Encode(n, next)
	for k := range h.Wheels[1:] {
		n, next = h.Wheels[k].Encode(n, next)
	}

	// Reflector ->
	n = h.Reflector.Encode(n)

	// Decode back, through wheels <-
	for k := range h.Wheels {
		k = len(h.Wheels) - 1 - k
		// start from the end
		n = h.Wheels[k].Decode(n)
	}

	//plugboard <-
	o = h.Plugboard.Decode(n)

	return

}

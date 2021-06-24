package enigma

import (
	"fmt"
)

type (
	Enigma struct {
		Plugboard   *SB
		StaticWheel *ETW
		Rotors      []*RotorWheel
		Reflector   *UKW
	}
	EnigmaSetting struct {
		//Rotor placement I II III etc.
		Walzenlage []string
		// Rotor Dial Setting
		Ringstellung []byte
		//Plugboard, Steckerverbindung
		Steckerverbindung [][]byte
		//StaticWheel ETW, usually 1-1
		Eintrittswalze string
		//Reflector UKW A, B, C
		Reflector string
		// Rotor Innter offset setting
		Rotorstellung []byte
	}
)

const (
	numChars = 26
)

/*
keyboard - plugboard - static wheel - w0, w1, w2 - reflector
*/

func NewEnigma(m EnigmaSetting) *Enigma {

	h := new(Enigma)

	//Wire plugboard
	h.Plugboard = NewPlugboard(m.Steckerverbindung)

	// insert static input wheel, scambled on some machines
	h.StaticWheel = &ETW{W: NewWheel(m.Eintrittswalze)}

	//inser reflector
	h.Reflector = &UKW{W: NewWheel(m.Reflector)}

	//insert rotors
	for k, v := range m.Walzenlage {
		switch v {
		case "I":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_I))
		case "II":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_II))
		case "III":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_III))
		case "IV":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_IV))
		case "V":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_V))
		case "VI":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_VI))
		case "VII":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_VII))
		case "VIII":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_VIII))
		case "B":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_Beta))
		case "G":
			h.Rotors = append(h.Rotors, NewRotorWheel(k, W_Gamma))
		default:
			fmt.Printf("failed to find config for rotor type %v", v)
		}
		h.Rotors[k].SetRingPosition(m.Ringstellung[k]) // ring setting visible from outside
		h.Rotors[k].SetInnerOffset(m.Rotorstellung[k]) // inner rotor offset
	}
	return h

}

func (h *Enigma) Codec(message string) string {
	var o []byte
	for _, v := range []byte(message) {
		c := h.Step(v)
		o = append(o, c)
	}
	return string(o)
}

func (h *Enigma) Step(i byte) (o byte) {
	index := i - 'A' // index of char

	//plugboard ->
	index = h.Plugboard.Encode(index)

	//Input Rotor ->
	index = h.StaticWheel.Encode(index)

	// Main Rotors ->
	next := true
	// always increment first wheel
	//index, next = h.Rotors[0].Encode(index, next)
	for k := range h.Rotors {
		next = h.Rotors[k].Step(next)
	}

	for k := range h.Rotors {
		index = h.Rotors[k].Encode(index)
	}

	// Reflector ->
	out := h.Reflector.Encode(index)
	fmt.Printf("Reflect in=%2v:%v out=%2v:%v\n", index, string(index+'A'), out, string(out+'A'))
	index = out

	// Decode back, through Rotors <-
	for k := range h.Rotors {
		k = len(h.Rotors) - 1 - k
		// start from the end
		index = h.Rotors[k].Decode(index)
	}

	//plugboard <-
	index = h.Plugboard.Decode(index)
	o = index + 'A'

	return
}

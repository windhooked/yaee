package enigma

import "fmt"

/*
http://users.telenet.be/d.rijmenants/en/enigmatech.htm
//https://de.wikipedia.org/wiki/Enigma-M4

            A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
I           E K M F L G D Q V Z N T O W Y H X U S P A I B R C J
II          A J D K S I R U X B L H W T M C Q G Z N P Y F V O E
III         B D F H J L C P R T X V Z N Y E I W G A K M U S Q O
IV          E S O V P Z J A Y Q U I R H X L N F T G K D C M W B
V           V Z B R G I T Y U P S D N H L X A W M J Q O F E C K
VI          J P G V O U M F Y Q B E N H Z R D K A S X L I C T W
VII         N Z J H G R C X M Y S W B O U F A I V L P E K Q D T
VIII        F K Q H T L X O C B J S P D Z R A M E W N I U Y G V

// did not auto increment
            A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
Beta        L E Y J V C N I X W P B Q M D R T A K Z G F U H O S
Gamma       F S O K A N U E R H M B T I Y C W L Q P Z X V G J D

*/
var (
	// used to translate from and to char <-> index
	W_I    = RotorWheel{codeWheel: W{setting: []byte("EKMFLGDQVZNTOWYHXUSPAIBRCJ")}, notch: []byte("Q")}  // Q->R
	W_II   = RotorWheel{codeWheel: W{setting: []byte("AJDKSIRUXBLHWTMCQGZNPYFVOE")}, notch: []byte("E")}  // E->F
	W_III  = RotorWheel{codeWheel: W{setting: []byte("BDFHJLCPRTXVZNYEIWGAKMUSQO")}, notch: []byte("V")}  // V->W
	W_IV   = RotorWheel{codeWheel: W{setting: []byte("ESOVPZJAYQUIRHXLNFTGKDCMWB")}, notch: []byte("J")}  // J->K
	W_V    = RotorWheel{codeWheel: W{setting: []byte("VZBRGITYUPSDNHLXAWMJQOFECK")}, notch: []byte("Z")}  // Z->A
	W_VI   = RotorWheel{codeWheel: W{setting: []byte("JPGVOUMFYQBENHZRDKASXLICTW")}, notch: []byte("ZM")} // Z->A M-N
	W_VII  = RotorWheel{codeWheel: W{setting: []byte("NZJHGRCXMYSWBOUFAIVLPEKQDT")}, notch: []byte("ZM")} // Z->A M-N
	W_VIII = RotorWheel{codeWheel: W{setting: []byte("FKQHTLXOCBJSPDZRAMEWNIUYGV")}, notch: []byte("ZM")} //  Z auf A und von M auf N
)

type (
	//
	RotorWheel struct {
		codeWheel     W
		innerPosition byte   // adjustable inner Setting (offset), stored as index
		notch         []byte // notch position
		notchIndex    []uint8
		rotorPosition uint8 //rotor position
		//lutIn         []byte //lookup index to char
		//lutOut        []byte //lookup char to index
		charIndex W //lookup
	}
)

//func init() {
//	W_I.Build()
//W_II.Build()
//W_II.Build()
//W_IV.Build()
//W_V.Build()
//W_VI.Build()
//W_VII.Build()
//W_VIII.Build()
//}

func (h *RotorWheel) Build() {
	h.codeWheel.Build()
	h.charIndex = CharacterSet
	h.charIndex.Build()
	for _, v := range h.notch {
		h.notchIndex = append(h.notchIndex, h.charIndex.GetIndex(v))
	}
	//h.lutIn = make([]uint8, 255)  //numChars+1)
	//h.lutOut = make([]uint8, 255) //numChars+1)
	//for k, v := range ring.setting {
	//	h.lutIn[k] = v
	//	h.lutOut[v] = byte(k)
	//}
}

//http://users.telenet.be/d.rijmenants/en/enigmatech.htm
func (h *RotorWheel) Encode(in byte, bump bool) (out byte, notch bool) {
	// the rotor first increments by 1
	if bump {
		h.rotorPosition += 1
		if h.rotorPosition >= 26 {
			h.rotorPosition = 0
		}
	}
	// chek if the input index alligns with the notch, if yes pump the next rotor
	i := h.charIndex.GetIndex(in)
	for _, v := range h.notchIndex {
		if v == i {
			notch = true
		}
	}

	return h.transposeIn(in), notch

}

func (h *RotorWheel) Decode(in byte) (out byte) {
	return h.transposeOut(in)
}

// (mapping[(k + shift + 26) % 26] - shift + 26) % 26;
func (h *RotorWheel) transposeIn(in byte) (out byte) {
	i := h.charIndex.GetIndex(in)
	// calculate the offset for ring setting and rotor position
	//offset := (i + h.innerSetting + h.rotorPosition) % 26
	offset := (i + h.rotorPosition) % 26
	fmt.Printf("offset: %v\n", offset)
	c := h.charIndex.GetChar(offset) //get the character at offset
	iCode := h.charIndex.GetIndex(c) // get the index for new character
	// now code the new char
	iCode = (iCode + h.innerPosition) % 26
	out = h.codeWheel.GetChar(iCode) // get the coded char for that offset
	return
}

func (h *RotorWheel) transposeOut(in byte) (out byte) {
	i := h.codeWheel.GetIndex(in)
	// calculate the offset for ring setting and rotor position
	//offset := (i + h.innerSetting + h.rotorPosition) % 26
	offset := (i - h.rotorPosition) % 26
	fmt.Printf("offset: %v\n", offset)
	c := h.codeWheel.GetChar(offset) //get the character at offset
	iCode := h.codeWheel.GetIndex(c) // get the index for new character
	// now code the new char
	iCode = (iCode - h.innerPosition) % 26
	out = h.charIndex.GetChar(iCode) // get the coded char for that offset
	return
}

/*
func (h *RotorWheel) EncodeOld(in byte, bump bool) (out byte, notch bool) {
	notch = false
	// check if notch position is support two notch rotors VI, VII, VII
	if in == h.notch[0] || (len(h.notch) > 1 && in == h.notch[1]) {
		notch = true
	}

	// translate ringsetting
	i := h.lutOut[in] // get the index for input char
	i = i + h.index   // offset with rotational index
	if i > numChars {
		i = i - numChars
	}
	in = h.lutIn[i] // lookup new char

	// now encode with wheel
	out = h.codeWheel.Encode(in)
	// if first wheel increment on every
	if bump {
		h.index++
	}

	if h.index < 0 {
		// should never be
		h.index = 0

	}
	if h.index > numChars {
		// start over
		h.index = 0
	}
	return
}
*/
func (h *RotorWheel) Lut() []byte {
	return h.codeWheel.Lut()
}
func (h *RotorWheel) SetInner(in byte) {
	h.innerPosition = h.charIndex.GetIndex(in)
}
func (h *RotorWheel) SetRing(in byte) {
	h.rotorPosition = h.charIndex.GetIndex(in)
}

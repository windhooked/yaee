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
	W_I     = RotorWheelSetting{Name: "I", Code: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", Notch: []byte("Q")}     // Q->R
	W_II    = RotorWheelSetting{Name: "II", Code: "AJDKSIRUXBLHWTMCQGZNPYFVOE", Notch: []byte("E")}    // E->F
	W_III   = RotorWheelSetting{Name: "III", Code: "BDFHJLCPRTXVZNYEIWGAKMUSQO", Notch: []byte("V")}   // V->W
	W_IV    = RotorWheelSetting{Name: "IV", Code: "ESOVPZJAYQUIRHXLNFTGKDCMWB", Notch: []byte("J")}    // J->K
	W_V     = RotorWheelSetting{Name: "V", Code: "VZBRGITYUPSDNHLXAWMJQOFECK", Notch: []byte("Z")}     // Z->A
	W_VI    = RotorWheelSetting{Name: "VI", Code: "JPGVOUMFYQBENHZRDKASXLICTW", Notch: []byte("ZM")}   // Z->A M-N
	W_VII   = RotorWheelSetting{Name: "VII", Code: "NZJHGRCXMYSWBOUFAIVLPEKQDT", Notch: []byte("ZM")}  // Z->A M-N
	W_VIII  = RotorWheelSetting{Name: "VIII", Code: "FKQHTLXOCBJSPDZRAMEWNIUYGV", Notch: []byte("ZM")} //  Z auf A und von M auf N
	W_Beta  = RotorWheelSetting{Name: "B", Code: "LEYJVCNIXWPBQMDRTAKZGFUHOS"}
	W_Gamma = RotorWheelSetting{Name: "G", Code: "FSOKANUERHMBTIYCWLQPZXVGJD"}
)

type (
	//
	RotorWheel struct {
		Name          string //used for debugging
		Pos           int    //used for debugging
		codeWheel     *W
		innerPosition byte   // adjustable inner Setting (offset), stored as index
		notch         []byte // notch position
		notchIndex    []uint8
		rotorPosition uint8 //rotor position
		charIndex     *W    //lookup
	}
	RotorWheelSetting struct {
		Name     string
		Position uint8
		Code     string
		Notch    []byte
	}
)

func NewRotorWheel(pos int, w RotorWheelSetting) *RotorWheel {
	h := &RotorWheel{
		codeWheel: NewWheel(w.Code),
		charIndex: NewWheel(CharacterSet),
		notch:     w.Notch,
		Name:      w.Name,
		Pos:       pos,
	}
	for _, v := range w.Notch {
		h.notchIndex = append(h.notchIndex, h.charIndex.GetIndex(v))
	}
	return h
}

func (h *RotorWheel) Step(bump bool) (notch bool) {
	// the rotor first increments by 1
	if bump {
		h.rotorPosition += 1
		// the rotor first increments by 1
		if h.rotorPosition >= 26 {
			h.rotorPosition = 0
		}
		// chek if the input index alligns with the notch, if yes pump the next rotor
		for _, v := range h.notchIndex {
			if v == h.rotorPosition {
				notch = true
			}
		}
	}
	fmt.Printf(">Rotor%v:% 3v: pos=%2v, step=% 5v, carry=%v\n", h.Pos, h.Name, h.rotorPosition, bump, notch)
	return notch
}

//http://users.telenet.be/d.rijmenants/en/enigmatech.htm
func (h *RotorWheel) Encode(in uint8) (out uint8) {

	out = h.transpose(in, false)

	fmt.Printf(">Rotor%v:% 3v: pos=%2v, in=%2v:%v out=%2v:%v\n", h.Pos, h.Name, h.rotorPosition, in, string(in+'A'), out, string(out+'A'))
	return out

}

func (h *RotorWheel) Decode(in uint8) (out uint8) {
	out = h.transpose(in, true)
	fmt.Printf("<Rotor%v:% 3v: pos=%2v,  in=%2v:%v out=%2v:%v \n", h.Pos, h.Name, h.rotorPosition, in, string(in+'A'), out, string(out+'A'))
	return out
}

func (h *RotorWheel) transpose(in byte, reverse bool) (out byte) {
	//shift the input charater by rotor position and inner offset
	in = ((in) - h.rotorPosition + h.innerPosition + byte(CharacterSetCount)) % byte(CharacterSetCount)

	if !reverse {
		out = h.codeWheel.Encode(in) // get the coded char index for that offset
	} else {
		out = h.codeWheel.Decode(in) // get the decoded char index for that offset
	}

	//shift the resulting index by rotor position and inner offset
	out = (out + h.rotorPosition - h.innerPosition + byte(CharacterSetCount)) % byte(CharacterSetCount)
	return
}

func (h *RotorWheel) Lut() []byte {
	return h.codeWheel.Lut()
}

// Must be a character , not the index
func (h *RotorWheel) SetInnerOffset(in byte) {
	h.innerPosition = h.charIndex.GetIndex(in)
}

// Must be a character , not the index
func (h *RotorWheel) SetRingPosition(in byte) {
	h.rotorPosition = h.charIndex.GetIndex(in)
}

func (h *RotorWheel) wrap(index, bufferLength uint8) uint8 {
	return ((index % bufferLength) + bufferLength) % bufferLength
}

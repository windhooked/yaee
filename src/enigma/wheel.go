package enigma

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
	ring   = W{setting: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")}
	W_I    = Wheel{W: W{setting: []byte("EKMFLGDQVZNTOWYHXUSPAIBRCJ")}, notch: []byte("Q"), ringSetting: 'A'}  // Q->R
	W_II   = Wheel{W: W{setting: []byte("AJDKSIRUXBLHWTMCQGZNPYFVOE")}, notch: []byte("E"), ringSetting: 'A'}  // E->F
	W_III  = Wheel{W: W{setting: []byte("BDFHJLCPRTXVZNYEIWGAKMUSQO")}, notch: []byte("V"), ringSetting: 'A'}  // V->W
	W_IV   = Wheel{W: W{setting: []byte("ESOVPZJAYQUIRHXLNFTGKDCMWB")}, notch: []byte("J"), ringSetting: 'A'}  // J->K
	W_V    = Wheel{W: W{setting: []byte("VZBRGITYUPSDNHLXAWMJQOFECK")}, notch: []byte("Z"), ringSetting: 'A'}  // Z->A
	W_VI   = Wheel{W: W{setting: []byte("JPGVOUMFYQBENHZRDKASXLICTW")}, notch: []byte("ZM"), ringSetting: 'A'} // Z->A M-N
	W_VII  = Wheel{W: W{setting: []byte("NZJHGRCXMYSWBOUFAIVLPEKQDT")}, notch: []byte("ZM"), ringSetting: 'A'} // Z->A M-N
	W_VIII = Wheel{W: W{setting: []byte("FKQHTLXOCBJSPDZRAMEWNIUYGV")}, notch: []byte("ZM"), ringSetting: 'A'} //  Z auf A und von M auf N
)

type (
	//Walze
	Wheel struct {
		W
		ringSetting byte   // Setting
		notch       []byte // Notch position
		index       uint8
		lutIn       []byte
		lutOut      []byte
	}
)

func init() {
	W_I.Build()
	W_II.Build()
	W_II.Build()
	W_IV.Build()
	W_V.Build()
	W_VI.Build()
	W_VII.Build()
	W_VIII.Build()
}

func (h *Wheel) Build() {
	h.W.Build()
	h.lutIn = make([]uint8, 255)  //numChars+1)
	h.lutOut = make([]uint8, 255) //numChars+1)
	for k, v := range ring.setting {
		h.lutIn[k] = v
		h.lutOut[v] = byte(k)
	}
}

//http://users.telenet.be/d.rijmenants/en/enigmatech.htm
func (h *Wheel) Encode(in byte, bump bool) (out byte, notch bool) {
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
	out = h.W.Encode(in)
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

func (h *Wheel) Decode(in byte) (out byte) {
	//k := strings.Index(string(h.lut), string(in))
	//return h.lut[k]
	return h.W.Decode(in)
}
func (h *Wheel) Lut() []byte {
	return h.W.Lut()
}
func (h *Wheel) SetRing(in byte) {
	h.ringSetting = in
	h.index = in - 'A'
}

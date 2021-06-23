package enigma

// the wheel is a generic pattern used by any coding component in the enigma
// the coding components: static input wheel, rotors, plugboard and reflector, all
// have a input to output code associated, translating 26 input to 26 output
var CharacterSet = W{setting: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")}

type (
	W struct {
		setting        []byte
		lutIndexToChar []byte
		lutCharToIndex []byte
	}
)

func (h *W) Build() {
	h.lutIndexToChar = make([]uint8, 255) //numChars+1)
	h.lutCharToIndex = make([]uint8, 255) //numChars+1)
	for k, v := range h.setting {
		h.lutIndexToChar[k] = v
		h.lutCharToIndex[v] = uint8(k)
	}
}

//func (h *W) Encode(in uint8) (out byte) {
//	return h.lutIndexToChar[in]
//}

//func (h *W) EncodeOffset(in byte, offset uint8) (out byte) {
//	return h.lutIn[in-'A']
//}

//func (h *W) Decode(in byte) (out uint8) {
//k := strings.Index(string(h.lutIn), string(in))
// find the input index, and then the character
//	return h.lutCharToIndex[in-'A']
//}

// GetIndex returns the index for given character
// eg. A == 1, Z == 26
func (h *W) GetIndex(in byte) (out uint8) {
	return h.lutCharToIndex[in]
}

// GetChar returns the character for a given index
// eg. 1 == A, 26 == Z
func (h *W) GetChar(in uint8) (out byte) {
	return h.lutIndexToChar[in]
}

//func (h *W) DecodeOffset(in byte, offset uint8) (out byte) {
//	return h.lutIn[h.lutOut[in]]
//}
func (h *W) Lut() []byte {
	return h.setting
}

func (h *W) wrap(n uint8) uint8 {
	if n > numChars {
	}
	return n
}

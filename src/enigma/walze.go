package enigma

type (
	W struct {
		setting []byte
		lutIn   []byte
		lutOut  []byte
	}
)

func (h *W) Build() {
	h.lutIn = make([]uint8, 255)  //numChars+1)
	h.lutOut = make([]uint8, 255) //numChars+1)
	for k, v := range h.setting {
		h.lutIn[k] = v
		h.lutOut[v] = byte(k)
	}
}

func (h *W) Encode(in byte) (out byte) {
	return h.lutIn[in-'A']
}

func (h *W) Decode(in byte) (out byte) {
	//k := strings.Index(string(h.lutIn), string(in))
	// find the input index, and then the character
	return h.lutIn[h.lutOut[in]]
}
func (h *W) Lut() []byte {
	return h.lutIn[0:26]
}

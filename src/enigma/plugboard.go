package enigma

/*
https://de.wikipedia.org/wiki/Enigma-M4

----------------------------------------------------------------------
| Mo-  |                                                    | Grund- |
| nats-|      S t e c k e r v e r b i n d u n g e n         | stel-  |
| tag  |                                                    | lung   |
----------------------------------------------------------------------
|  30. |18/26 17/4 21/6 3/16 19/14 22/7 8/1 12/25 5/9 10/15 |H F K D |
|  29. |20/13 2/3 10/4 21/24 12/1 6/5 16/18 15/8 7/11 23/26 |O M S R |
|  28. |9/14 4/5 18/24 3/16 20/26 23/21 12/19 13/2 22/6 1/8 |E Y D X |
|  27. |16/2 25/21 6/20 9/17 22/1 15/4 18/26 8/23 3/14 5/19 |T C X K |
|  26. |20/13 26/11 3/4 7/24 14/9 16/10 8/17 12/5 2/6 15/23 |Y S R B |
*/
var (
	PB_30 = []PB{
		//|  30. |18/26 17/4 21/6 3/16 19/14 22/7 8/1 12/25 5/9 10/15 |H F K D |
		{A: 18, B: 26},
		{A: 17, B: 4},
		{A: 21, B: 6},
		{A: 3, B: 16},
		{A: 19, B: 14},
		{A: 22, B: 7},
		{A: 8, B: 1},
		{A: 12, B: 25},
		{A: 5, B: 9},
		{A: 10, B: 15},
	}
)

type (
	SB struct {
		Setting []PB
		lutIn   []uint8 // lookup table
		lutOut  []uint8 // lookup table
	}
	PB struct {
		A uint8 // from
		B uint8 //to
	}
)

func (h *SB) WireUp(pb []PB) {
	h.Setting = pb
	h.lutIn = make([]uint8, numChars+1)
	h.lutOut = make([]uint8, numChars+1)
	for _, v := range pb {
		h.lutIn[v.A] = v.B
		h.lutOut[v.B] = v.A
	}
}

// Encode on the plugboard reroutes if plug wires are inserted, else input bridges to output
func (h *SB) Encode(n uint8) (out uint8) {
	out = h.lutIn[n-'A']
	if out == 0 {
		out = n
	}
	return
}
func (h *SB) Decode(n uint8) (out uint8) {
	out = h.lutOut[n-'A']
	if out == 0 {
		out = n
	}
	return
}

package enigma_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

//https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030669
//Reflector: C, Greek: B, Wheels: 568, Wheel positions: DGOE, Rings: AAEL, Plugs: AE BF CM DQ HU JN LX PR SZ VW
const test_P1030669 = "VFOFZTNTOQXHMHYSUARPWEDAEEOXNYDZQZHXMFXGMRPCOFERVIVUQNGCSAOVXDZWRUGVADACKFUOOTDXQZBNXDGVXBFPOEVRPBECSYYSIABWAWGCWCFZROYAXSRGVNSLUUIPMTQIKLEZTANXBANMTFKZJNOITINZVCIEGBXADZTMKYPWTTDZXZKDIBZITQRESNHLQIITTPUNKRAZTBSOMIMLLWTLEKVDSFQMFBBECFHEDCAQWIPINCLAUVBSJKCMOXXMJGEPMIOFUEOXPGQUIYOWVPDCNSHW"
const result_P1030669 = "SSDCHEFFUNFXUUUFLOTTXXEINSKKTTTFFFEINSACHTMITZWOSTELLKARTENAUSRUESTUNGSKAGERRAKUNDNORWEGENMITHAFENPLAENENUNDEINSEXEMPLARLFDXBEFXBBBDDDUUUOOOPPPNRXSIEBENVOMZWOVIERXVIERXVIERFUNFXSOFORTNEUSTADTINMARSCHSETZENXZWOKKTTTFFFEINBNEUNVORLPEUFIGKIELBLEIBENXDHEIKKBESTARTEHXNGERBETENXXFXDXUUUAUSBXKF"

// test on https://cryptii.com/pipes/enigma-machine

//https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030675
// longer message

func TestEnigma(t *testing.T) {
	testCipher := strings.ToUpper(strings.ReplaceAll(test_P1030669, " ", ""))

	m4 := enigma.NewEnigma(enigma.EnigmaSetting{
		Walzenlage: []string{"VIII", "VI", "V", "B"},
		//Ringstellung: []byte{'E', 'O', 'G', 'D'}, //
		Ringstellung: []byte{'L', 'E', 'A', 'A'}, //
		Steckerverbindung: [][]byte{
			//{'A', 'E'}, {'D', 'Q'}, {'R', 'C'}, {'V', 'B'}, {'M', 'T'}, {'O', 'G'}, {'P', 'F'}, {'Y', 'L'}, {'J', 'W'}, {'I', 'Z'},
			//{'B', 'Q'}, {'C', 'R'}, {'D', 'I'}, {'E', 'J'}, {'K', 'W'}, {'M', 'T'}, {'O', 'S'}, {'P', 'X'}, {'U', 'Z'}, {'G', 'H'},
			//{'A', 'T'}, {'C', 'L'}, {'D', 'H'}, {'E', 'P'}, {'F', 'G'}, {'I', 'O'}, {'J', 'N'}, {'K', 'Q'}, {'M', 'U'}, {'R', 'X'},
			{'A', 'E'}, {'B', 'F'}, {'C', 'M'}, {'D', 'Q'}, {'H', 'U'}, {'J', 'N'}, {'L', 'X'}, {'P', 'R'}, {'S', 'Z'}, {'V', 'W'},
		},
		Eintriswalze: enigma.ETW_M4,
		Reflector:    enigma.UKW_Casar,
		//Rotorstellung: []byte{'L', 'E', 'A', 'A'},
		Rotorstellung: []byte{'E', 'O', 'G', 'D'},
	})

	//code := m4.Step(testCipher[0])
	code := m4.Codec(testCipher)

	fmt.Printf(">> %v", string(code))
}

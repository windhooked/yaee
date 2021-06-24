package enigma_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/windhooked/yaee/src/enigma"
)

//PASS tag v1
func TestEnigmaAAA123(t *testing.T) {
	//path should be
	// plugboard AA FF
	//ETW AA    FF
	//I BK ,    DG
	//II CD ,   DC
	//III DH ,  HD
	//UKW B ->  DH ^
	//https://piotte13.github.io/enigma-cipher/

	m1 := enigma.NewEnigma(enigma.EnigmaSetting{
		Walzenlage:        []string{"I", "II", "III"},
		Ringstellung:      []byte{'A', 'A', 'A'}, //
		Steckerverbindung: [][]byte{},
		Eintrittswalze:    enigma.ETW_M4,
		Reflector:         enigma.UKW_B,
		Rotorstellung:     []byte{'A', 'A', 'A'},
	})

	code := m1.Codec(enigma.CharacterSet)
	if "FUVEPUMWARVQKEFGHGDIJFMFXI" == string(code) {
		t.Logf("result FUVEPUMWARVQKEFGHGDIJFMFXI == %v", string(code))
	} else {
		t.Fatalf("cipher text does not match for setting 123 AAA")
	}
}

//PASS tag v1
func TestEnigmaAAA123Plugboard(t *testing.T) {
	//https://piotte13.github.io/enigma-cipher/

	settings := enigma.EnigmaSetting{
		Walzenlage:   []string{"I", "II", "III"},
		Ringstellung: []byte{'A', 'A', 'A'}, //
		Steckerverbindung: [][]byte{
			{'A', 'E'},
			{'Z', 'L'},
		},
		Eintrittswalze: enigma.CharacterSet,
		Reflector:      enigma.UKW_B,
		Rotorstellung:  []byte{'A', 'A', 'A'},
	}

	// encode A-Z
	m10 := enigma.NewEnigma(settings)
	code := m10.Codec(enigma.CharacterSet)
	// decode
	m11 := enigma.NewEnigma(settings)
	dCode := m11.Codec(code)

	if "VUVAGUMWERVRKAFGHGDIJFMFXF" == string(code) {
		t.Logf("result VUVAGUMWERVRKAFGHGDIJFMFXF == %v", string(dCode))
	} else {
		t.Fatalf("cipher text does not match for setting 123 AAA with plugboard setting AE,ZL")
	}

}

//https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030669
//Reflector: C, Greek: B, Wheels: 568, Wheel positions: DGOE, Rings: AAEL, Plugs: AE BF CM DQ HU JN LX PR SZ VW
const test_P1030669 = "VFOFZTNTOQXHMHYSUARPWEDAEEOXNYDZQZHXMFXGMRPCOFERVIVUQNGCSAOVXDZWRUGVADACKFUOOTDXQZBNXDGVXBFPOEVRPBECSYYSIABWAWGCWCFZROYAXSRGVNSLUUIPMTQIKLEZTANXBANMTFKZJNOITINZVCIEGBXADZTMKYPWTTDZXZKDIBZITQRESNHLQIITTPUNKRAZTBSOMIMLLWTLEKVDSFQMFBBECFHEDCAQWIPINCLAUVBSJKCMOXXMJGEPMIOFUEOXPGQUIYOWVPDCNSHW"
const result_P1030669 = "SSDCHEFFUNFXUUUFLOTTXXEINSKKTTTFFFEINSACHTMITZWOSTELLKARTENAUSRUESTUNGSKAGERRAKUNDNORWEGENMITHAFENPLAENENUNDEINSEXEMPLARLFDXBEFXBBBDDDUUUOOOPPPNRXSIEBENVOMZWOVIERXVIERXVIERFUNFXSOFORTNEUSTADTINMARSCHSETZENXZWOKKTTTFFFEINBNEUNVORLPEUFIGKIELBLEIBENXDHEIKKBESTARTEHXNGERBETENXXFXDXUUUAUSBXKF"

// test on https://cryptii.com/pipes/enigma-machine

//https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030675
// longer message

//Test Fails
func TestEnigma_P1030669(t *testing.T) {
	testCipher := strings.ToUpper(strings.ReplaceAll(test_P1030669, " ", ""))

	m4 := enigma.NewEnigma(enigma.EnigmaSetting{
		Walzenlage:   []string{"VIII", "VI", "V", "B"},
		Ringstellung: []byte{'L', 'E', 'A', 'A'}, //
		Steckerverbindung: [][]byte{
			//{'A', 'E'}, {'D', 'Q'}, {'R', 'C'}, {'V', 'B'}, {'M', 'T'}, {'O', 'G'}, {'P', 'F'}, {'Y', 'L'}, {'J', 'W'}, {'I', 'Z'},
			//{'B', 'Q'}, {'C', 'R'}, {'D', 'I'}, {'E', 'J'}, {'K', 'W'}, {'M', 'T'}, {'O', 'S'}, {'P', 'X'}, {'U', 'Z'}, {'G', 'H'},
			//{'A', 'T'}, {'C', 'L'}, {'D', 'H'}, {'E', 'P'}, {'F', 'G'}, {'I', 'O'}, {'J', 'N'}, {'K', 'Q'}, {'M', 'U'}, {'R', 'X'},
			{'A', 'E'}, {'B', 'F'}, {'C', 'M'}, {'D', 'Q'}, {'H', 'U'}, {'J', 'N'}, {'L', 'X'}, {'P', 'R'}, {'S', 'Z'}, {'V', 'W'},
		},
		Eintrittswalze: enigma.ETW_M4,
		Reflector:      enigma.UKW_Casar,
		Rotorstellung:  []byte{'E', 'O', 'G', 'D'},
	})

	//code := m4.Step(testCipher[0])
	code := m4.Codec(testCipher)

	fmt.Printf(">> %v", string(code))
}

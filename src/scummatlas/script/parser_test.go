package script

import "testing"
import "io/ioutil"

func readScriptOrDie(filename string, t *testing.T) []byte {
	data, err := ioutil.ReadFile("./testdata/" + filename + ".dump")
	if err != nil {
		t.Errorf("Error reading the file")
	}
	var initialOffset int
	blockType := string(data[0:4])
	switch blockType {
	case "SCRP", "VERB":
		initialOffset = 8
	case "LSCR":
		initialOffset = 9
	}
	return data[initialOffset:]
}

func checkScriptLengthAndOpcodes(file string, expectedOpcodes []byte, t *testing.T) {
	data := readScriptOrDie(file, t)
	script := ParseScriptBlock(data)
	if len(script) != len(expectedOpcodes) {
		t.Errorf("File %v, length mismatch, got %d and expected %d",
			file, len(script), len(expectedOpcodes))
	}
	for i, _ := range script {
		if len(script) <= i || len(expectedOpcodes) <= i {
			return
		}
		if script[i].opCode != expectedOpcodes[i] {
			t.Errorf("File %v, expecting opcode %x but got %x in position %d",
				file, expectedOpcodes[i], script[i].opCode, i)
		}
	}
}

func TestRoomScript1(t *testing.T) {
	checkScriptLengthAndOpcodes("monkey2_11_202",
		[]byte{0x1A, 0x80, 0xE8, 0x28, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey2_11_200",
		[]byte{
			0x13, 0x11, 0x2D, 0x01, 0x2A}, t)

	checkScriptLengthAndOpcodes("monkey2_11_210",
		[]byte{0x40, 0x93, 0x91, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xC0, 0x93, 0x2A, 0x24}, t)

	checkScriptLengthAndOpcodes("monkey2_11_209",
		[]byte{0x1A, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey2_11_208",
		[]byte{0x1A, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x1C, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey2_11_203",
		[]byte{0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0x18}, t)

	checkScriptLengthAndOpcodes("monkey1_11_200",
		[]byte{0x40, 0x1A, 0x05, 0x5D, 0x2E, 0x1C, 0x2E, 0x2A, 0x80, 0x68, 0x28, 0x1C, 0x2E, 0x0A, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0x1A, 0x2A, 0x80, 0xA8, 0x2A, 0x80, 0x68, 0x28, 0x48, 0x0A, 0x80, 0x68, 0x28, 0x48, 0x28, 0x1A, 0x0A, 0x33, 0x80, 0x80, 0x80, 0x33, 0x07, 0x5D, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x18, 0xD8, 0xAE, 0xD8, 0xC0}, t)

	checkScriptLengthAndOpcodes("monkey1_12_206",
		[]byte{0x1D, 0x1A, 0xC8, 0xDD, 0x18, 0x1D, 0x5D, 0xDD, 0x18, 0xDD, 0x46, 0x44, 0x9D, 0xD8, 0x18, 0x9D, 0xD8, 0x18, 0xD8}, t)

	checkScriptLengthAndOpcodes("monkey1_2_200",
		[]byte{0xF5, 0x48, 0xC3, 0x78, 0x9E, 0x18, 0xB6, 0xAE, 0x37, 0x18, 0x42}, t)

	checkScriptLengthAndOpcodes("monkey2_g_17",
		[]byte{0x48, 0x1A, 0xAB, 0xAB, 0xAB, 0xAB, 0x1A, 0x0A, 0x0A, 0x1A, 0x18, 0x48, 0x1A, 0xA8, 0xFA, 0x1A, 0xAB, 0xAB, 0xAB, 0xAB, 0x1A, 0x62, 0x1A, 0x18, 0x48, 0x1A, 0xA8, 0xFA, 0x1A, 0x2C, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x2C, 0x0A, 0x18}, t)

	checkScriptLengthAndOpcodes("monkey2_g_1",
		[]byte{0x40, 0x58, 0x18, 0x33, 0x72, 0x33, 0x72, 0x80, 0x1C, 0x4C, 0x48, 0x4C, 0x4C, 0x9A, 0x44, 0x80, 0x04, 0x18, 0x2E, 0x58, 0x3C, 0x33, 0x72, 0x33, 0xC0, 0x1A, 0x24, 0x33}, t)

	checkScriptLengthAndOpcodes("monkey1_g_10",
		[]byte{0x1D, 0x48, 0x25, 0x07, 0x5D, 0x40, 0x11, 0x58, 0x18, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x80, 0x80, 0x80, 0xD8, 0xC0}, t)

	checkScriptLengthAndOpcodes("monkey1_g_7",
		[]byte{0x48, 0x68, 0x28, 0x0A, 0x18, 0x48, 0x4A, 0x2E, 0x4A, 0x2E, 0x13, 0x13, 0x4A, 0x2E, 0x4A, 0x80, 0x68, 0x28, 0x80, 0x1A, 0x0C, 0x0C, 0x18, 0x80, 0x6A, 0x2E, 0x6A}, t)

	checkScriptLengthAndOpcodes("monkey1_g_1",
		[]byte{0x1A, 0x83, 0x48, 0x93, 0xC3, 0xA3, 0xED, 0xE1, 0xAC, 0xD1, 0x91, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xAD, 0x18, 0x46, 0x44}, t)

	checkScriptLengthAndOpcodes("monkey1_g_88",
		[]byte{0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xDD, 0x93, 0x0A, 0x0A, 0x0A, 0x48, 0x93, 0x18, 0x93, 0xA8, 0x0A, 0x18}, t)

	checkScriptLengthAndOpcodes("monkey1_g_96",
		[]byte{0x62, 0x40, 0x9E, 0xAE, 0x91, 0x80, 0x91, 0x2D, 0x01, 0x13, 0x28, 0x1A, 0x58, 0x18, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0x14, 0x1E, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0xAE, 0x11, 0x33, 0x72, 0x1C, 0x80, 0xA8, 0x58, 0x14, 0x33, 0xD2, 0x18, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0x14, 0x1E, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0xAE, 0x11, 0xD2, 0x01, 0x5D, 0x54, 0x91, 0xC0, 0x0A, 0x2A, 0x80, 0x9A, 0x80, 0xF8, 0x0A}, t)
}

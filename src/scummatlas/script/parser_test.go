package script

import "testing"
import "io/ioutil"
import "strings"

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

func checkScriptResult(rawscript []byte, expected string, t *testing.T) {
	script := ParseScriptBlock(rawscript)
	if len(script) == 0 {
		t.Errorf("Empty result, was expecting %v", expected)
		return
	}
	result := strings.TrimSpace(script.Print())
	if result != expected {
		t.Errorf("Expression not valid, expected %v but got %v",
			expected, result)
	}
}

func TestExpression(t *testing.T) {
	checkScriptResult(
		[]byte{0xac, 0x76, 0x00, 0x81, 0x03, 0x40, 0x01, 0x01, 0x00, 0x03, 0xff},
		"var[118] = expression((local[3] - 1))", t)

	checkScriptResult(
		[]byte{0xac, 0x64, 0x00, 0x01, 0xc8, 0x00, 0x81, 0x05, 0x40, 0x02, 0x01, 0x01, 0x00, 0x03, 0xff},
		"var[100] = expression(((200 + local[5]) - 1))", t)
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
		[]byte{0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18,
			0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48,
			0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18,
			0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0x18, 0x2A, 0x80, 0x68,
			0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80,
			0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28,
			0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0x18}, t)

	checkScriptLengthAndOpcodes("monkey1_11_200",
		[]byte{0x40, 0x1A, 0x05, 0x5D, 0x2E, 0x1C, 0x2E, 0x2A, 0x80, 0x68, 0x28, 0x1C, 0x2E, 0x0A, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC,
			0x1A, 0x2A, 0x80, 0xA8, 0x2A, 0x80, 0x68, 0x28, 0x48, 0x0A, 0x80, 0x68, 0x28, 0x48, 0x28, 0x1A, 0x0A, 0x33, 0x80, 0x80,
			0x80, 0x33, 0x07, 0x5D, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x18, 0xD8, 0xAE, 0xD8, 0xC0}, t)

	checkScriptLengthAndOpcodes("monkey1_12_206",
		[]byte{0x1D, 0x1A, 0xC8, 0xDD, 0x18, 0x1D, 0x5D, 0xDD, 0x18, 0xDD, 0x46, 0x44, 0x9D, 0xD8, 0x18, 0x9D, 0xD8, 0x18, 0xD8}, t)

	checkScriptLengthAndOpcodes("monkey1_2_200",
		[]byte{0xF5, 0x48, 0xC3, 0x78, 0x9E, 0x18, 0xB6, 0xAE, 0x37, 0x18, 0x42}, t)

	checkScriptLengthAndOpcodes("monkey2_g_17",
		[]byte{0x48, 0x1A, 0xAB, 0xAB, 0xAB, 0xAB, 0x1A, 0x0A, 0x0A, 0x1A, 0x18, 0x48, 0x1A, 0xA8, 0xFA, 0x1A, 0xAB, 0xAB, 0xAB, 0xAB,
			0x1A, 0x62, 0x1A, 0x18, 0x48, 0x1A, 0xA8, 0xFA, 0x1A, 0x2C, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x7A, 0x2C,
			0x0A, 0x18}, t)

	checkScriptLengthAndOpcodes("monkey2_g_1",
		[]byte{0x40, 0x58, 0x18, 0x33, 0x72, 0x33, 0x72, 0x80, 0x1C, 0x4C, 0x48, 0x4C, 0x4C, 0x9A, 0x44, 0x80, 0x04, 0x18, 0x2E, 0x58,
			0x3C, 0x33, 0x72, 0x33, 0xC0, 0x1A, 0x24, 0x33}, t)

	checkScriptLengthAndOpcodes("monkey1_g_10",
		[]byte{0x1D, 0x48, 0x25, 0x07, 0x5D, 0x40, 0x11, 0x58, 0x18, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0xD8, 0xAE, 0x2E, 0x11,
			0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE,
			0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE,
			0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE, 0x2E, 0x11, 0x2E, 0x11, 0xD8, 0xAE, 0xD8, 0xAE, 0xD8, 0xAE,
			0x2E, 0x11, 0x80, 0x80, 0x80, 0xD8, 0xC0}, t)
	checkScriptLengthAndOpcodes("monkey1_g_7",
		[]byte{0x48, 0x68, 0x28, 0x0A, 0x18, 0x48, 0x4A, 0x2E, 0x4A, 0x2E, 0x13, 0x13, 0x4A, 0x2E, 0x4A, 0x80, 0x68, 0x28, 0x80,
			0x1A, 0x0C, 0x0C, 0x18, 0x80, 0x6A, 0x2E, 0x6A}, t)

	checkScriptLengthAndOpcodes("monkey1_g_1",
		[]byte{0x1A, 0x83, 0x48, 0x93, 0xC3, 0xA3, 0xED, 0xE1, 0xAC, 0xD1,
			0x91, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xAD, 0x18, 0x46, 0x44}, t)

	checkScriptLengthAndOpcodes("monkey1_g_88",
		[]byte{0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xDD, 0x93, 0x0A, 0x0A, 0x0A, 0x48, 0x93, 0x18, 0x93, 0xA8, 0x0A, 0x18}, t)

	checkScriptLengthAndOpcodes("monkey1_g_96",
		[]byte{0x62, 0x40, 0x9E, 0xAE, 0x91, 0x80, 0x91, 0x2D, 0x01, 0x13, 0x28, 0x1A, 0x58, 0x18, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78,
			0x14, 0x1E, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0xAE, 0x11, 0x33, 0x72, 0x1C, 0x80, 0xA8, 0x58, 0x14, 0x33, 0xD2, 0x18,
			0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0x14, 0x1E, 0x1A, 0x21, 0x1C, 0x80, 0xC6, 0x78, 0xAE, 0x11, 0xD2, 0x01, 0x5D, 0x54,
			0x91, 0xC0, 0x0A, 0x2A, 0x80, 0x9A, 0x80, 0xF8, 0x0A}, t)

	checkScriptLengthAndOpcodes("monkey1_g_53",
		[]byte{0x62, 0x14, 0x2C, 0x2C, 0x80, 0xA8, 0x16, 0x48, 0x14, 0x18, 0x48, 0x14, 0x18, 0x48, 0x14, 0x18, 0x48, 0x14, 0xAE, 0x14, 0x18, 0x48, 0x14,
			0x18, 0x48, 0x14, 0xAE, 0x18, 0x9A, 0x33, 0xA8, 0x12, 0x1A, 0x1A, 0x1A, 0x1A, 0x18, 0xA8, 0x91, 0x2E, 0x72, 0xAD, 0x81,
			0x93, 0x1A, 0x1A, 0x1A, 0x1A, 0x18, 0x72, 0xAD, 0x81, 0x93, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x0A, 0x80, 0x80, 0x80, 0x14,
			0xAE, 0x0A, 0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46,
			0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A, 0x1A, 0x16, 0x80, 0xA8, 0x14, 0xAE, 0x48, 0x14, 0xAE, 0x14, 0x18, 0x48, 0x14,
			0xAE, 0x14, 0x18, 0x48, 0x14, 0x18, 0x48, 0x14, 0xAE, 0x18, 0x58, 0x18, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14,
			0xAE, 0x14, 0xAE, 0x58, 0x14, 0x1A, 0x0A, 0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A,
			0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A, 0xAC, 0xFA, 0x5A, 0x3A, 0x46, 0x1A, 0x1A, 0x16, 0x80, 0xA8, 0x14, 0xAE, 0x48, 0x14,
			0xAE, 0x18, 0x18, 0x48, 0x14, 0xAE, 0x14, 0xAE, 0x18, 0x18, 0x48, 0x18, 0x48, 0x14, 0xAE, 0x18, 0x58, 0x18, 0x14, 0xAE,
			0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE, 0x14, 0xAE,
			0x14, 0xAE, 0x58, 0x14, 0x1A, 0x80, 0xAE, 0x33, 0xA8, 0x18, 0xA8, 0xAD, 0x81, 0x93, 0x72, 0x91, 0x2E, 0x91, 0x18, 0xAD,
			0x81, 0x93, 0x72, 0x9A, 0x0A, 0x2A, 0x14, 0x2C, 0x2C}, t)

	checkScriptLengthAndOpcodes("monkey1_g_12",
		[]byte{0x2C, 0x2C, 0x1A, 0x48, 0x1A, 0x18, 0x1A, 0x38, 0x3A, 0x14, 0x18, 0x0A, 0x80, 0x68, 0x28, 0x27, 0x33, 0x27, 0x9A, 0x17,
			0xA8, 0x1A, 0x18, 0x1A, 0x9A, 0x17, 0xA8, 0x1A, 0x18, 0x1A, 0x9A, 0x17, 0xA8, 0x1A, 0x18, 0x1A, 0x9A, 0x17, 0xA8, 0x1A,
			0x18, 0x1A, 0x38, 0x27, 0x48, 0x4C, 0x18, 0x4C, 0x1A, 0x1A, 0x1A, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0x27, 0x1A, 0x1A, 0x1A, 0x1A, 0x0C, 0x0C, 0xA8, 0x0C, 0x2C, 0x72, 0x14, 0x0C, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x1A, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x1A, 0x1A, 0x0C, 0x0C, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x26, 0x26, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x1A, 0x1A,
			0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x1A, 0xAC, 0x1A, 0x0A, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0x0C, 0x2C, 0x2C, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x0A, 0x0A, 0x33, 0x27, 0x27, 0x27,
			0x27, 0x27, 0x27, 0x27, 0xCC, 0xCC, 0xCC, 0x28, 0x0A, 0xCC, 0xCC, 0xCC, 0xCC, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27,
			0xCC, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x27, 0x1A, 0x33, 0x48, 0x1A, 0x0A, 0x80, 0x68, 0x28, 0x1A, 0x48, 0x0A, 0x80,
			0x68, 0x28, 0x18, 0x1A, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C,
			0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C,
			0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C,
			0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C,
			0x1A, 0x1A, 0x1A, 0x72, 0x33, 0x0A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x14, 0x1A,
			0x1A, 0x1A, 0x13, 0x08, 0x1A, 0x72, 0x28, 0x14, 0x98, 0x18, 0x38, 0x04, 0x9A, 0x5B, 0x1B, 0xAC, 0x9A, 0x9A, 0x5B, 0x1B,
			0x5B, 0xBA, 0x9A, 0x9A, 0x5B, 0x1B, 0x5B, 0xBA, 0x48, 0x18, 0x48, 0x1A, 0x18, 0x48, 0x1A, 0x1A, 0x18, 0x48, 0x1A, 0x1A,
			0x1A, 0x18, 0x48, 0x1A, 0x18, 0x48, 0x1A, 0x1A, 0x18, 0x48, 0x1A, 0x1A, 0x1A, 0x18, 0x48, 0x1A, 0x1A, 0x1A, 0x1A, 0x48,
			0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x25, 0x29,
			0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x37, 0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x72, 0x25, 0x29, 0x37, 0x1A, 0x18,
			0x48, 0x72, 0x25, 0x29, 0x1A, 0x72, 0x25, 0x29, 0x37, 0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x25, 0x29, 0x1A, 0x72,
			0x25, 0x29, 0x37, 0x1A, 0x48, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x18, 0x48, 0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A,
			0x1A, 0x18, 0x48, 0x72, 0x25, 0x29, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x72, 0x25, 0x29, 0x18, 0x48, 0x1A, 0x72, 0x25,
			0x29, 0x18, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x1A, 0x72, 0x25, 0x29, 0x1A, 0x48, 0x1A, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25,
			0x72, 0x25, 0x25, 0x25, 0x72, 0x25, 0x5D, 0x72, 0x25, 0x72, 0x25, 0x37, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x28, 0x18,
			0x48, 0x1A, 0x1A, 0x1A, 0x72, 0x25, 0x72, 0x25, 0x25, 0x25, 0x1A, 0x1A, 0x1A, 0x48, 0x1A, 0x1A, 0x72, 0x25, 0x25, 0x25,
			0x07, 0x72, 0x25, 0x72, 0x25, 0x25, 0x25, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x87, 0x46, 0x44, 0x07, 0x5D, 0x5D, 0x5D, 0x5D,
			0x5D, 0x5D, 0x5D, 0x07, 0x07, 0x07, 0x07, 0x07, 0x48, 0x1A, 0x1A, 0x72, 0x25, 0x25, 0x25, 0x07, 0x72, 0x25, 0x72, 0x25,
			0x72, 0x25, 0x25, 0x25, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x87, 0x46, 0x44, 0x07, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D,
			0x07, 0x07, 0x07, 0x07, 0x07, 0x48, 0x1A, 0x1A, 0x07, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x25, 0x25, 0x1A,
			0x1A, 0x1A, 0x1A, 0x1A, 0x87, 0x46, 0x44, 0x07, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x07, 0x07, 0x07, 0x07, 0x07,
			0x48, 0x1A, 0x1A, 0x07, 0x72, 0x25, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x25, 0x25, 0x1A, 0x1A, 0x1A,
			0x1A, 0x1A, 0x87, 0x46, 0x44, 0x07, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x5D, 0x07, 0x07, 0x07, 0x07, 0x07, 0x48, 0x1A,
			0x1A, 0x07, 0x72, 0x25, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x25, 0x25, 0x1A, 0x1A, 0x1A, 0x48, 0x1A,
			0x1A, 0x48, 0x72, 0x25, 0x25, 0x5D, 0x72, 0x25, 0x54, 0x5D, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x48, 0x72, 0x25, 0x72,
			0x25, 0x1A, 0x1A, 0x1A, 0x48, 0x10, 0x88, 0x18, 0x72, 0x25, 0x25, 0x25, 0x72, 0x25, 0x25, 0x25, 0x72, 0x25, 0x72, 0x25,
			0x25, 0x1A, 0x1A, 0x1A, 0x1A, 0x48, 0x72, 0x25, 0x25, 0x25, 0x72, 0x25, 0x25, 0x72, 0x25, 0x25, 0x1A, 0x1A, 0x1A, 0x48,
			0x72, 0x25, 0x25, 0x25, 0x25, 0x72, 0x25, 0x25, 0x1A, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x48, 0x1A, 0x1A, 0x48, 0x72, 0x25,
			0x07, 0x1A, 0x48, 0x72, 0x25, 0x1A, 0x48, 0x1A, 0x1A, 0x48, 0x72, 0x25, 0x54, 0x5D, 0x1A, 0x1A, 0x1A, 0x48, 0x1A, 0x1A,
			0x72, 0x25, 0x54, 0x5D, 0x1A, 0x1A, 0x1A, 0x1A, 0x48, 0x72, 0x25, 0x1A, 0x48, 0x1A, 0x1A, 0x72, 0x25, 0x72, 0x25, 0x25,
			0x1A, 0x44, 0xAC, 0x48, 0x72, 0x25, 0x25, 0x1A, 0x48, 0x72, 0x25, 0x72, 0x25, 0x1A, 0x48, 0x72, 0x25, 0x1A, 0x48, 0x1A,
			0x1A, 0x1A, 0x5D, 0x1A, 0x48, 0x72, 0x25, 0x1A, 0x1A, 0x48, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x1A, 0x48, 0x72, 0x25,
			0x29, 0x1A, 0x72, 0x25, 0x29, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x1A, 0x48, 0x72, 0x25, 0x72, 0x25, 0x72, 0x1A,
			0xED, 0x81, 0x91, 0x0A, 0xD2, 0x62, 0x48, 0x1A, 0x72, 0x25, 0x72, 0x37, 0x72, 0x25, 0x1A, 0x48, 0x1A, 0x1A, 0x72, 0x25,
			0x1A, 0x48, 0x1A, 0x1A, 0x37, 0x1A, 0x48, 0x72, 0x25, 0x25, 0x1A, 0x48, 0x37, 0x72, 0x25, 0x1A, 0x48, 0x1A, 0x1A, 0x48,
			0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x72, 0x25, 0x1A, 0x48, 0x72, 0x25, 0x72, 0x25, 0x1A, 0x48,
			0x1A, 0x72, 0x25, 0x1A, 0x48, 0x1A, 0x72, 0x25, 0x1A, 0x48, 0x1A, 0x72, 0x25, 0x1A, 0x1A, 0x1A, 0x46, 0x44, 0x1A, 0x48,
			0x1A, 0x1A, 0x72, 0x25, 0x1A, 0x1A, 0x48, 0x1A, 0x1A, 0x72, 0x25, 0x1A, 0x1A, 0x1A, 0x48, 0x1A, 0x1A, 0x1A, 0x46, 0x44,
			0x1A, 0x1A, 0x46, 0x44, 0x1A, 0x48, 0x37, 0x72, 0x25, 0x25, 0x25, 0x72, 0x25, 0x1A, 0x48, 0x72, 0x25, 0x72, 0x25, 0x5D,
			0x72, 0x25, 0x07, 0x5D, 0x1A, 0x44, 0x18, 0x78, 0x18, 0x72, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x0C, 0x13, 0x48, 0x1A, 0x1A,
			0x48, 0x72, 0x42, 0x48, 0x72, 0x25, 0x29, 0x1A, 0x48, 0x72, 0x0A, 0x1A, 0x80, 0x20, 0xED, 0x81, 0x18, 0xAD, 0x81, 0x91,
			0x0A, 0xD2, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey1_g_77",
		[]byte{0xF1, 0x08, 0xD8, 0x18, 0x48, 0xFB, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x28, 0x40, 0x91, 0x91, 0xE3,
			0x08, 0x91, 0xAE, 0xA8, 0x1A, 0x54, 0x28, 0x1A, 0xD8, 0xAE, 0x93, 0x91, 0x80, 0x80, 0x80, 0x80, 0x91, 0x1A, 0xC0, 0x68,
			0xA8, 0x62, 0x48, 0x2A, 0x18, 0x48, 0x0A, 0x18, 0x48, 0x2A, 0x28, 0x1A, 0x19, 0x18, 0x40, 0x58, 0x18, 0x91, 0x2E, 0x28,
			0x1A, 0x94, 0xAE, 0x94, 0xAE, 0x2E, 0x91, 0x1C, 0x80, 0x80, 0x80, 0x80, 0x91, 0x91, 0x80, 0x80, 0x80, 0x80, 0x1A, 0x91,
			0x1C, 0x80, 0x80, 0x46, 0x44, 0x80, 0x80, 0x91, 0x91, 0x2E, 0x94, 0xAE, 0xC0, 0x0A}, t)

	checkScriptLengthAndOpcodes("monkey1_g_11",
		[]byte{0x40, 0x14, 0x93, 0x91, 0x80, 0x80, 0xA5, 0x80, 0x80, 0x9E, 0xAE, 0x91, 0x80, 0x80, 0x91, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x93, 0x07, 0x5D, 0xD8, 0x30, 0x30, 0xC0, 0x29}, t)

	checkScriptLengthAndOpcodes("monkey1_g_28",
		[]byte{0x5A, 0x2C, 0x2C, 0x08, 0x19, 0x0A, 0x0A, 0x1A, 0xA8, 0xFA, 0x48, 0xFA, 0x1A, 0x48, 0xAB, 0xAB, 0xAB, 0xAB, 0x48, 0xAB, 0xAB, 0xAB, 0xAB, 0x60}, t)

	checkScriptLengthAndOpcodes("monkey1_g_35",
		[]byte{0x28, 0x9A, 0x8F, 0x48, 0x9D, 0x40, 0x87, 0xA8, 0x9D, 0x87, 0x48, 0x0A, 0x18, 0x48, 0x0A, 0x18, 0x48, 0x0A, 0x18, 0x48, 0x0A, 0x18, 0x48, 0x0A, 0x18, 0x0A, 0xC0, 0x18, 0xD8}, t)

	checkScriptLengthAndOpcodes("monkey2_g_7",
		[]byte{0x48, 0x33, 0x33, 0x33, 0x33, 0x18, 0x33, 0x33, 0x33, 0x33, 0x0A, 0x0A, 0x9A, 0x19}, t)

	checkScriptLengthAndOpcodes("monkey2_g_50",
		[]byte{0x40, 0x58, 0x18, 0x14, 0xAE, 0x14, 0xAE, 0xD8, 0xAE, 0x0C, 0x11, 0x13, 0x0A, 0x2D, 0x0E, 0x2A, 0x1E, 0x14, 0x4C, 0x4C,
			0x4C, 0x2E, 0x1C, 0x4C, 0x4C, 0x4C, 0xAE, 0x14, 0x11, 0x09, 0xAE, 0x09, 0x14, 0x11, 0xAE, 0xAE, 0x11, 0x14, 0xAE, 0x14,
			0xAE, 0x11, 0x07, 0x2E, 0x11, 0x80, 0x07, 0x13, 0x0A, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
			0x80, 0x80, 0x80, 0x80, 0x80, 0x07, 0x13, 0x0A, 0x80, 0x11, 0x2E, 0x07, 0x2E, 0x11, 0x1A, 0x11, 0x13, 0x11, 0x80, 0x1C,
			0x80, 0x80, 0x80, 0x91, 0x91, 0x80, 0x80, 0x80, 0x80, 0x13, 0x2D, 0x72, 0x11, 0x13, 0x2D, 0x01, 0x1E, 0xAE, 0x72, 0x2A,
			0x2D, 0x2A, 0x11, 0x11, 0x11, 0x13, 0x2D, 0x01, 0x11, 0x91, 0x1E, 0xAE, 0x05, 0x1C, 0x2D, 0x5D, 0xAE, 0x2E, 0x11, 0x91,
			0x80, 0x91, 0x11, 0x80, 0x13, 0x11, 0x80, 0x11, 0x11, 0x14, 0xAE, 0x80, 0x13, 0x11, 0x11, 0x80, 0x11, 0x11, 0x80, 0x2E,
			0x14, 0xAE, 0xAE, 0x11, 0x14, 0x1E, 0xAE, 0x11, 0x80, 0x80, 0x80, 0x11, 0x80, 0x80, 0x11, 0x80, 0x11, 0x80, 0x80, 0x80,
			0x11, 0x11, 0xAE, 0x14, 0xAE, 0x11, 0x14, 0xAE, 0x36, 0x14, 0xAE, 0x14, 0x2E, 0x11, 0xAE, 0x4C, 0x4C, 0xD8, 0x91, 0xAE,
			0x2D, 0xAE, 0x14, 0xAE, 0xA8, 0x14, 0xD2, 0x2D, 0x2D, 0x91, 0x91, 0x8E, 0x13, 0x11, 0x11, 0x01, 0x1A, 0x05, 0x5D, 0x07,
			0x07, 0x4C, 0x4C, 0x4C, 0x18, 0x58, 0x11, 0x2A, 0x1C, 0x4C, 0x62, 0x48, 0x33}, t)

	checkScriptLengthAndOpcodes("monkey1_g_42",
		[]byte{0x1A, 0x1A, 0x1A, 0x2C, 0x2C, 0x1A, 0xFA, 0x46, 0x44, 0x26, 0x26}, t)

	/*
			checkScriptLengthAndOpcodes("monkey2_g_3",
				[]byte{0xC8, 0xC8, 0x46, 0x44, 0x1A, 0xA8, 0x42, 0x18, 0x1A, 0x9A, 0x9A, 0x9A, 0x04, 0x9D, 0xCD, 0xAE, 0xC9, 0x8A, 0x0A, 0x62,
					0x48, 0x44, 0x9D, 0x90, 0x48, 0xD8, 0x18, 0x28, 0xF6, 0xAE, 0xF4, 0x04, 0xB7, 0x18, 0x18, 0x90, 0x48, 0xD8, 0x18, 0xA8,
					0x79, 0x62, 0x28, 0xCD, 0xAE, 0xAC, 0xEC, 0xDA, 0xF4, 0x84, 0x8B, 0xA8, 0xC9, 0xC9, 0x9A, 0xB7, 0x0A, 0x18, 0x18, 0x9D,
					0x8A, 0x18, 0x42, 0x18, 0xCB, 0x28, 0x9A, 0x9A, 0x9D, 0x08, 0x42, 0x18, 0x18, 0x28, 0x90, 0x48, 0x9D, 0xFE, 0xAE, 0x08,
					0x42, 0x18, 0x18, 0x18, 0xF6, 0xAE, 0xF4, 0x04, 0x08, 0x42, 0x18, 0x18, 0x18, 0x1A, 0x08, 0x48, 0x08, 0x90, 0x48, 0x9D,
					0xF9, 0x59, 0x18, 0x46, 0x90, 0x88, 0x90, 0x48, 0x9D, 0xF9, 0x59, 0x18, 0x46, 0x48, 0x42, 0x9A, 0x90, 0x08, 0xA8, 0x90,
					0x48, 0x9A, 0x90, 0x48, 0x28, 0x9D, 0x80, 0x18, 0x9D, 0xFE, 0xAE, 0x18, 0x18, 0xF6, 0xAE, 0xF4, 0x44, 0xC9, 0xAE, 0xD8,
					0x0A, 0x0A, 0x19, 0x62, 0xA8, 0x48, 0x18, 0x9D, 0x1A, 0x18, 0x9D, 0x1A, 0x18, 0x9D, 0x1A, 0x18, 0x9D, 0x1A, 0x18, 0x9D,
					0x1A, 0x18, 0x9D, 0x1A, 0x18, 0x9D, 0x1A, 0x18, 0x1A, 0xA8, 0x90, 0x48, 0x08, 0x08, 0x08, 0x40, 0xD1, 0x2E, 0x91, 0xC0,
					0x9A, 0xF7, 0x0A, 0x1A}, t)

		checkScriptLengthAndOpcodes("monkey2_g_23",
			[]byte{0x44, 0xA8, 0x18, 0x44, 0x44, 0xAC, 0xAC, 0x9A, 0x48, 0x9A, 0x1A, 0x48, 0x9A, 0x1A, 0x18, 0x18, 0x18, 0xF5, 0xA8, 0x9A,
				0x1A, 0x08, 0xD5, 0xA8, 0x9A, 0x48, 0x9D, 0x18, 0xD5, 0x9D, 0x18, 0x1A, 0x48, 0xA8, 0x9D, 0x18, 0xD5, 0x9D, 0x18, 0x1A,
				0x88, 0x9A, 0xA8, 0x88, 0x9A, 0x18, 0x1A, 0x18, 0x9A, 0xA8, 0xFA, 0x1A, 0x9A, 0xA8, 0x44, 0xB7, 0x18, 0xCA, 0x08, 0xA8,
				0xFA, 0x18, 0x1A, 0xFA, 0x18, 0x1A, 0x9A, 0xE8, 0x28, 0x0A, 0x80, 0x18}, t)
	*/
}

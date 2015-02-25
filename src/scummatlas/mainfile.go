package scummatlas

import (
	"fmt"
)

type MainScummData struct {
	Data []byte
}

type RoomOffset struct {
	Number int
	Offset int
}

type Image bool
type Script bool
type Box bool
type BoxMatrix bool

type Room struct {
	Width    int
	Height   int
	ObjCount int
	//ColorCycle ColorCycle
	//TranspColor TranspColor
	//Palette Palette
	Image         Image
	ObjectImage   Image
	ObjectScripts []Script
	ExitScript    Script
	EntryScript   Script
	LocalScript   Script
	BoxData       []Box
	BoxMatrix     BoxMatrix
}

func (d *MainScummData) GetRoomCount() int {
	blockName := string(d.Data[0:4])
	blockSize := BE32(d.Data[4 : 4+4])
	if blockName != "LECF" {
		panic("No main container in the file")
	}
	fmt.Println(blockName, blockSize)

	blockName = string(d.Data[8 : 8+4])
	if blockName != "LOFF" {
		panic("No room offset table in the file")
	}
	blockSize = BE32(d.Data[12 : 12+4])
	fmt.Println(blockName, blockSize)
	roomCount := int(d.Data[16])
	fmt.Println("roomCount", roomCount)
	return roomCount
}

func (d *MainScummData) GetRoomsOffset() (offsets []RoomOffset) {
	count := d.GetRoomCount()
	currentOffset := 17
	var out []RoomOffset
	for i := 0; i < count; i++ {
		count := int(d.Data[currentOffset])
		offset := LE32(
			d.Data[currentOffset+1 : currentOffset+5])
		roomOffset := RoomOffset{count, offset}
		out = append(out, roomOffset)
		currentOffset += 5
	}
	return out
}

func (d *MainScummData) ParseRoom(offset int) Room {
	blockName := string(d.Data[offset : offset+4])
	//blockSize := string(d.Data[offset : offset+4])
	if blockName != "ROOM" {
		panic("No room header found")
	}

	return Room{}
}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	srcAddress := uint16(1)     // src address
	srcUnit := uint8(2)         // src unit
	dstAddress := uint16(1)     // destination address
	dstUnit := uint8(0)         // destination unit
	transport := uint16(0)      // transport layer bytes
	appRef := uint8(1)          // app ref to record order of messages
	msgType := uint8(7)         // message type
	interfaceUid := uint16(1)   // interface uuid for keep alive interval
	interfaceMember := uint8(1) // interface member
	DataLen := uint8(4)         // Length of data payload
	Data := uint32(600000)      // value in milliseconds for sending keep alive signal

	// create our bytes buffer
	buf := new(bytes.Buffer)
	fmt.Println("Adding source address bytes to buf:")
	err := binary.Write(buf, binary.BigEndian, srcAddress)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding source unit bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, srcUnit)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding dstAddress bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, dstAddress)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding dstUnit bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, dstUnit)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding transport bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, transport)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding appRef bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, appRef)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding msgType bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, msgType)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding interfaceUid bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, interfaceUid)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding interfaceMember bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, interfaceMember)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding DataLen bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, DataLen)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("Adding Data bytes to buf:")
	err = binary.Write(buf, binary.BigEndian, Data)
	if err != nil {
		panic(fmt.Errorf("binary.Write failed: %v", err))
	}
	fmt.Println(buf.Bytes())

	fmt.Println("What does JSON look like as bytes?")

	someJson := "{\"Bird\":10,\"Cat\":\"Fuzzy\"}"
	fmt.Println([]byte(someJson))

}

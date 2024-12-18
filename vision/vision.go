package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

type KeyPart struct {
	offset int
	length int16
}

type Key struct {
	nparts     int
	duplicates bool
	length     int
	boh        int16
	parts      []KeyPart
}

type VisionFile struct {
	vFile             *os.File
	vRAFile           *os.File
	vVersion          rune
	blockingFactor    int
	blockSize         int
	blockSize_4       int
	numOfRecords      int64
	deletedRecords    int64
	totOpens          int64
	userCount         int64
	segmentSize       int64
	maxRec            int
	minRec            int
	nKeys             int
	keys              []Key
	compressed        bool
	hasDuplicates     bool
	buffer            []byte
	node              []byte
	dummyInt          []byte
	nodePnt           int
	validRecordsNum   int64
	deletedRecordsNum int64
	fName             string
	nDataSegments     int
	currDataSegment   int
	firstDataBlock    int64
	collatingSequence []byte
	blanks            string
}

// func (vf *VisionFile) readInt48(raf *os.File) (int64, error) {
// 	var a rune
// 	var Return int64
// 	var err error
// 	if a, err = readRune(raf); err != nil {
// 		return 0, err
// 	}
// 	var tmp int32
// 	if err = binary.Read(raf, binary.LittleEndian, &tmp); err != nil {
// 		return 0, err
// 	}
// 	Return = int64(tmp) & 0xFFFFFFFF
// 	Return |= int64(a) << 32
// 	return Return, nil
// }

func (vf *VisionFile) readInt48(raf *os.File) (int64, error) {
	var a rune
	var Return int64
	var err error
	var b byte
	if b, err = readByte(raf); err != nil {
		return 0, err
	}
	a = rune(b)
	var tmp int32
	if err = binary.Read(raf, binary.LittleEndian, &tmp); err != nil {
		return 0, err
	}
	Return = int64(tmp) & 0xFFFFFFFF
	Return |= int64(a) << 32
	return Return, nil
}

func readByte(r io.Reader) (byte, error) {
	var b byte
	err := binary.Read(r, binary.LittleEndian, &b)
	return b, err
}

func (vf *VisionFile) VisionFile(fileName string) error {
	var err error
	vf.fName = fileName
	vf.vFile, err = os.Open(fileName)
	if err != nil {
		return err
	}
	vf.vRAFile, err = os.Open(fileName)
	if err != nil {
		return err
	}
	vf.vVersion, _, err = vf.vRAFile.ReadRune()
	if err != nil {
		return err
	}

	// func (vf *VisionFile) VisionFile(fileName string) error {
	// 	var err error
	// 	vf.fName = fileName
	// 	vf.vFile, err = os.Open(fileName)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	vf.vRAFile, err = os.Open(fileName)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	vf.vVersion, err = vf.readRune(vf.vRAFile)
	// 	if err != nil {
	// 		return err
	// 	}
	vf.blockingFactor, err = vf.myReadInt(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.blockSize = vf.blockingFactor * 512
	vf.blockSize_4 = vf.blockSize / 4
	vf.numOfRecords, err = vf.readInt48(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.deletedRecords, err = vf.readInt48(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.totOpens, err = vf.readInt48(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.userCount, err = vf.readInt48(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.segmentSize, err = vf.readInt48(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.maxRec, err = vf.myReadInt(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.minRec, err = vf.myReadInt(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.nKeys, err = vf.myReadInt(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.keys = make([]Key, vf.nKeys)
	for i := 0; i < vf.nKeys; i++ {
		vf.keys[i].nparts, err = vf.myReadInt(vf.vRAFile)
		if err != nil {
			return err
		}
		vf.keys[i].duplicates, err = vf.myReadBool(vf.vRAFile)
		if err != nil {
			return err
		}
		vf.keys[i].length, err = vf.myReadShort(vf.vRAFile)
		if err != nil {
			return err
		}
		vf.keys[i].boh, err = vf.myReadShort(vf.vRAFile)
		if err != nil {
			return err
		}
		vf.keys[i].parts = make([]KeyPart, vf.keys[i].nparts)
		for j := 0; j < vf.keys[i].nparts; j++ {
			vf.keys[i].parts[j].offset, err = vf.myReadInt(vf.vRAFile)
			if err != nil {
				return err
			}
			vf.keys[i].parts[j].length, err = vf.myReadShort(vf.vRAFile)
			if err != nil {
				return err
			}
		}
	}
	vf.compressed, err = vf.myReadBool(vf.vRAFile)
	if err != nil {
		return err
	}

	vf.hasDuplicates, err = vf.myReadBool(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.buffer = make([]byte, vf.blockSize)
	vf.node = make([]byte, 4)
	vf.dummyInt = make([]byte, 4)
	vf.nodePnt = 0
	vf.validRecordsNum = vf.numOfRecords - vf.deletedRecords
	vf.deletedRecordsNum = vf.deletedRecords
	vf.nDataSegments = int(vf.segmentSize / vf.blockSize)
	vf.currDataSegment = 0
	vf.firstDataBlock, err = vf.readInt48(vf.vRAFile)
	if err != nil {
		return err
	}
	vf.collatingSequence = make([]byte, 48)
	if _, err = io.ReadFull(vf.vRAFile, vf.collatingSequence); err != nil {
		return err
	}
	vf.blanks = "                                              "
	return nil
}

func (vf *VisionFile) fmt(Return string, len int, left bool) string {
	sLen := len(Return)
	if sLen < len {
		if left {
			Return = Return + vf.blanks[0:len-sLen]
		} else {
			Return = vf.blanks[0:len-sLen] + Return
		}
	}
	return Return
}

func (vf *VisionFile) fmt(s string, len int) string {
	return vf.fmt(s, len, false)
}

func (vf *VisionFile) fmt(n int64, len int, left bool) string {
	s := fmt.Sprintf("%d", n)
	return vf.fmt(s, len, left)
}

func (vf *VisionFile) fmt(n int64, len int) string {
	return vf.fmt(n, len, false)
}

func (vf *VisionFile) printInfo() {
	fmt.Println(vf.vFile.Name() + "  [vision version " + string(vf.vVersion) + "]")
	// ...
	// Tradução do código original omitida para brevidade
	// ...
}

func (vf *VisionFile) getValidRecordsNum() int64 {
	return vf.numOfRecords
}

func (vf *VisionFile) readNode3() (int, error) {
	// ...
	// Tradução do código original omitida para brevidade
	// ...
	return 0, nil
}

func (vf *VisionFile) myReadChar3() (int, error) {
	// ...
	// Tradução do código original omitida para brevidade
	// ...
	return 0, nil
}

// func (vf *VisionFile) myRead3(b []byte, offs int, len int) (int, error) {
// 	// ...

func main() {
	vf := &VisionFile{}
	err := vf.VisionFile("aigefcop")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	vf.printInfo()
}

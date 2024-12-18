package main

import (
	"bufio"
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
	vRAFile           *bufio.Reader
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

func (vf *VisionFile) readInt48(r io.Reader) (int64, error) {
	var a byte
	var Return int64
	var err error
	if err = binary.Read(r, binary.LittleEndian, &a); err != nil {
		return 0, err
	}
	var tmp int32
	if err = binary.Read(r, binary.LittleEndian, &tmp); err != nil {
		return 0, err
	}
	Return = int64(tmp) & 0xFFFFFFFF
	Return |= int64(a) << 32
	return Return, nil
}

func (vf *VisionFile) myReadInt(r io.Reader) (int, error) {
	var value int32
	err := binary.Read(r, binary.LittleEndian, &value)
	return int(value), err
}

func (vf *VisionFile) myReadShort(r io.Reader) (int16, error) {
	var value int16
	err := binary.Read(r, binary.LittleEndian, &value)
	return value, err
}

func (vf *VisionFile) myReadBool(r io.Reader) (bool, error) {
	var value byte
	err := binary.Read(r, binary.LittleEndian, &value)
	return value != 0, err
}

func (vf *VisionFile) VisionFile(fileName string) error {
	var err error
	vf.fName = fileName
	vf.vFile, err = os.Open(fileName)
	if err != nil {
		return err
	}
	vf.vRAFile = bufio.NewReader(vf.vFile) // Usando bufio.Reader para ler runes

	// Lendo a versão do arquivo (primeiro caractere)
	vf.vVersion, _, err = vf.vRAFile.ReadRune()
	if err != nil {
		return err
	}

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
		vf.keys[i].length, err = int(vf.myReadShort(vf.vRAFile))
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
			var keyPartLength int16
			keyPartLength, err = vf.myReadShort(vf.vRAFile)
			if err != nil {
				return err
			}
			vf.keys[i].parts[j].length = keyPartLength
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
	vf.nDataSegments = int(vf.segmentSize / int64(vf.blockSize)) // Convertendo para int64
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

func (vf *VisionFile) fmtStr(s string, length int, left bool) string {
	sLen := len(s)
	if sLen < length {
		if left {
			s = s + vf.blanks[:length-sLen]
		} else {
			s = vf.blanks[:length-sLen] + s
		}
	}
	return s
}

func (vf *VisionFile) fmtInt(n int64, length int, left bool) string {
	s := fmt.Sprintf("%d", n)
	return vf.fmtStr(s, length, left)
}

func (vf *VisionFile) printInfo() {
	fmt.Println(vf.fName + "  [vision version " + string(vf.vVersion) + "]")
	// Outros detalhes podem ser impressos aqui
}

func (vf *VisionFile) getValidRecordsNum() int64 {
	return vf.numOfRecords
}

func main() {
	vf := &VisionFile{}
	err := vf.VisionFile("aigefcop")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	vf.printInfo()
}

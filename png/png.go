package png

import (
	"bufio"
	"os"

	"github.com/daxschoof/png-stenography/utils"
)

type Png struct {
	Name string
	ByteArr []byte
}

func (p *Png) ReadFileIntoBinary() error {
	file, err := os.Open(p.Name)
	if err != nil {
		return err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return err
	}

 	fileSize := stats.Size()
	byteArr := make([]byte, fileSize)

	buf := bufio.NewReader(file)
	_, err = buf.Read(byteArr)
	if err != nil {
		return err
	}

	p.ByteArr = byteArr
	
	return nil
}

func (p *Png) CheckIfFileIsPng() bool {
	if len(p.ByteArr) < 8 {
		return false
	}

	firstEightPng := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	return utils.CheckSliceEquality(firstEightPng, p.ByteArr[0:8])
}
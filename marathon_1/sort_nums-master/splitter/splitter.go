package splitter

import (
	"io"
	"os"
	"sort"
	"strconv"
	"../fileoperations"
	"fmt"
)

//把二进制的文件，划分成多个小的文件
func SplistByteFileToChunks(file *os.File, bufferSize int64) (int, error) {
	bufferData := make([]byte, bufferSize)
	counter := 0
	for {
		n, err := file.Read(bufferData)
		if err != nil && err != io.EOF {
			fmt.Print("file read error,err = ",err)
			return -1, err
		}
		if n == 0 {
			break
		}
		tmpFilename := "tmp_" + strconv.Itoa(counter)
		counter++
		tmpFile, err := fileoperations.CreateFile(tmpFilename)
		defer  tmpFile.Close()
		tmpFile.Write(bufferData)
	}
	return counter,nil
}

// sortBuffer sorts a slice of integers and returns it
func sortBuffer(ints []int) ([]int, error) {
	sort.Ints(ints)
	return ints, nil
}

package Util

import (
	//"time"
	//"crypto/rand"
	//"math/rand"
	"../fileoperations"
	"../quicksort/quicksort"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"sync"
)

/**
author:ortonwu
产生一个二进制文件
args :	filename: 文件名称
		length : int32个数
*/
func CreatBinFile(filename string, length int32) {
	//随机种子
	//t := time.Now().Nanosecond()
	//rand.Seed(int64(t))
	//创建文件
	fp, err := os.Create(path.Join("./", filename))
	fmt.Print("create result:", fp, err)
	defer fp.Close()
	buf := new(bytes.Buffer)
	var i int32
	for i = 0; i < length; i++ {
		binary.Write(buf, binary.LittleEndian, i)
		fp.Write(buf.Bytes())
	}
}

//解析二进制文件
//默认小端进行解析
//解析后的整数输出到临时文件中
func ParseBinFile(sourceFileName string, targetFileName string, splitSymbol string,wg *sync.WaitGroup) {
	defer wg.Done()
	sourceFp, _ := os.Open(path.Join("./", sourceFileName))
	defer sourceFp.Close()
	targetFp, _ := os.Create(path.Join("./", targetFileName))
	defer targetFp.Close()
	data := make([]byte, 4)
	var k uint32
	var ints []uint32

	for {
		data = data[:cap(data)]
		n, err := sourceFp.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}
		data = data[:n]
		binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &k)
		ints = append(ints, k)
	}
	//进行排序
	quicksort.QuickSort(ints)
	fileoperations.WriteBufferToFile(ints, targetFp)
}

// sortBuffer sorts a slice of integers and returns it
func sortBuffer(ints []int) ([]int, error) {
	sort.Ints(ints)
	return ints, nil
}

package main

import (
	"./sort_nums-master/Util"
	"./sort_nums-master/fileoperations"
	"./sort_nums-master/splitter"
	"./sort_nums-master/merger"
	"./sort_nums-master/timers"
	"fmt"
	"strconv"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	var wg2 sync.WaitGroup

	//将二进制文件解析出来
	filePtr := "rand.txt"
	inputFile, _ := fileoperations.OpenFile(&filePtr)
	//var fileLength int64 = 81920
	var fileLength int64 = 1024*8*10*5
	//将字节文件转为多个小文件
	var splitTime time.Time  = time.Now();
	filenum, e := splitter.SplistByteFileToChunks(inputFile, fileLength)
	timers.TimeTrack(splitTime,"SplistByteFileToChunks()")
	if e != nil {
		fmt.Print("splitByteFileToChunks error")
		return
	}
	var tranAndSortTime time.Time  = time.Now();
	//将二进制小文件转换为十进制小文件(同时排序)
	wg2.Add(filenum)
	for i := 0; i < filenum; i++ {
		 go Util.ParseBinFile("tmp_"+strconv.Itoa(i), "tmp_oct_"+strconv.Itoa(i), ",",&wg2)
	}
	timers.TimeTrack(tranAndSortTime,"ParseBinFile()")
	//合并文件
	wg2.Wait() //需要等待所有的小文件都排序完毕后再merge
	var mergeTime time.Time  = time.Now()
	go merger.MergeRuns(filenum,&wg)
	timers.TimeTrack(mergeTime,"MergeRun()")
	//等待合并完成
	wg.Wait()
	return
}


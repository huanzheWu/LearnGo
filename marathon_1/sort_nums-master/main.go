package main

import (
	"../sort_nums-master/fileoperations"
	"../sort_nums-master/merger"
	"../sort_nums-master/splitter"
	"../sort_nums-master/timers"
	"log"
	"time"
)

func main() {
	var err error
	//filePtr := flag.String("file", "", "Input file")
	//	memoryPtr := flag.Int64("memory", 0, "Memory (in bytes) that will be used to hold and sort a series of numbers")

	//flag.Parse()

	//if *memoryPtr <= 0 {
	//	log.Fatal("Not valid memory value")
	//}

	filePtr := "test.txt"
	defer timers.TimeTrack(time.Now(), "Sorting")
	//打开新文件
	inputFile, err := fileoperations.OpenFile(&filePtr)
	if err != nil {
		log.Fatal(err)
	}
	//获取该文件信息（需要知道文件的大小）
	fi, err := inputFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	//将文件分割成多个临时文件，并存储在磁盘上
	//var memoryPtr int64 = 100
	calculatedChunkNum, err := splitter.SplitFileToChunks(inputFile, memoryPtr)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("File %s is %d bytes, will be split to %d chunks\n", *filePtr, fi.Size(), calculatedChunkNum)
	log.Printf("File %s is %d bytes, will be split to %d chunks\n", filePtr, fi.Size(), calculatedChunkNum)

	//合并多个有序文件
	err = merger.MergeRuns(int(calculatedChunkNum))
	if err != nil {
		log.Fatal(err)
	}
	//删除多个临时文件
	//err = fileoperations.CleanupTempFiles(int(calculatedChunkNum))
	if err != nil {
		log.Fatal(err)
	}
}

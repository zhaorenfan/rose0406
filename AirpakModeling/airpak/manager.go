/*
 * @Descripttion:
 * @version:
 * @Author: 招人烦
 * @Date: 2022-09-02 19:40:11
 * @LastEditors: 招人烦
 * @LastEditTime: 2022-09-03 17:49:25
 */
package airpak

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Parse(d *Data, path string) /*error*/ {
	file, err := os.Open(path)
	if err != nil {
		panic("输出文件为空")
		//return err
	}
	defer file.Close()

	//逐行读取
	_reader := bufio.NewReader(file)
	for {
		str, err := _reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", str)
		//逐行解析字符串
		arr := strings.Fields(str)

		fmt.Println(arr, "长度：", len(arr))
		if l := len(arr); l < 1 {
			continue
		}
		airpakType := arr[0]
		switch airpakType {
		case "DOMAIN":
			if l := len(arr); l != 2 {
				panic("DOMAIN描述错误，请检查")
			}
			fmt.Println("发现一个DOMAIN")
			d.AddObject(Domain{arr[1]})
		case "BLOCK":
			if l := len(arr); l != 7 {
				//[BLOCK block.4 FLUID PSTART {-7,0,-2} PEND {-5,-4,0}]
				panic("BLOCK描述错误, 请检查")
			}
			fmt.Println("发现一个BLOCK")
			bktypes := arr[2]
			bktype := FLUID
			switch bktypes {
			case "FLUID":
				bktype = FLUID
			}
			block := &Block{
				Name:      arr[1],
				BlockType: bktype,
				Point1:    Str2Point(arr[4]),
				Point2:    Str2Point(arr[6]),
			}
			d.AddObject(block)
		case "WALL":
			if l := len(arr); l != 6 {
				panic("WALL描述错误，请检查")
			}
			fmt.Println("发现一个WALL")
			wall := &Wall{
			 	Name:   arr[1],
			 	Point1: Str2Point(arr[3]),
			 	Point2: Str2Point(arr[5]),
			}
			d.AddObject(wall)
		case "OPENING":
			if l := len(arr); l != 6 {
				panic("OPENING描述错误，请检查")
			}
			fmt.Println("发现一个OPENING")
			opening := &Opening{
			 	Name:   arr[1],
			 	Point1: Str2Point(arr[3]),
			 	Point2: Str2Point(arr[5]),
			}
			d.AddObject(opening)
		case "PRISM":
			// PRISM prism.1 FLUID VERT1 {-4,3,0} VERT2 {0,3,6} VERT3 {0,3,6} HEIGHT 1
			if l := len(arr); l != 11 {
				panic("PRISM描述错误，请检查")
			}
			fmt.Println("发现一个PRISM")
			bktypes := arr[2]
			bktype := FLUID
			switch bktypes {
			case "FLUID":
				bktype = FLUID
			}
			h, err :=strconv.ParseFloat(arr[10], 64)
			if err!=nil{
				panic("PRISM高度错误，请检查")
			}
			prism := &Prism{
				Name:      arr[1],
				BlockType: bktype,
				Point1:    Str2Point(arr[4]),
				Point2:    Str2Point(arr[6]),
				Point3:    Str2Point(arr[8]),
				Height:    float32(h),
			}
			d.AddObject(prism)
		}

	}
	//return err
}

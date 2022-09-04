/*
 * @Descripttion:
 * @version:
 * @Author: 招人烦
 * @Date: 2022-09-02 19:40:11
 * @LastEditors: 招人烦
 * @LastEditTime: 2022-09-04 14:11:03
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
			//WALL wall.1 INCLINED PSTART {31.75,-9.35,0} PEND {33.94,-11.54,2.65} AXIS 2 ANGLE 45.0
			fmt.Println("发现一个WALL")
			// if l := len(arr); l != 6 {
			// 	panic("WALL描述错误，请检查")
			// }
			switch arr[2]{
			case "RECT":
				wall := &Wall{
					Name:   arr[1],
					WallType: RECT,
					Point1: Str2Point(arr[4]),
					Point2: Str2Point(arr[6]),
			   }
			   d.AddObject(wall)
			case "INCLINED":
				axis,err:= strconv.Atoi(arr[8])
				if err!=nil{
					panic("INCLINED旋转轴转换失败")
				}
				angle,err := strconv.ParseFloat(arr[10],64)
				if err!=nil{
					panic("INCLINED倾角转换失败")
				}
				wall := &Wall{
					Name:     arr[1],
					WallType: INCLINED,
					Point1:   Str2Point(arr[4]),
					Point2:   Str2Point(arr[6]),
					Axis:     axis,
					Angle:    float32(angle),
				}
			   d.AddObject(wall)
			case "POLYGON":
				// WALL wall.1 POLYGON NVERTS 4 VERT1 {31,-9.,0} VERT2 {33.9,-11.54,0} VERT3 {33.9,-11.5,2.65} VERT4 {31.7,-9.3,2.65}
				nverts,err:= strconv.Atoi(arr[4])
				if err!=nil{
					panic("POLYGON点数量转换失败")
				}
				switch nverts{
				case 3:
					wall := &Wall{
						Name:     arr[1],
						WallType: POLYGON,
						Point1:   Str2Point(arr[6]),
						Point2:   Str2Point(arr[8]),
						Point3:   Str2Point(arr[10]),
						NVerts:   nverts,
						//Point4:   Str2Point(arr[12]),
					}
					d.AddObject(wall)
				case 4:
					wall := &Wall{
						Name:     arr[1],
						WallType: POLYGON,
						Point1:   Str2Point(arr[6]),
						Point2:   Str2Point(arr[8]),
						Point3:   Str2Point(arr[10]),
						NVerts:   nverts,
						Point4:   Str2Point(arr[12]),
					}
					d.AddObject(wall)
				}
				
			}
			
			
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
			fmt.Println(h)
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

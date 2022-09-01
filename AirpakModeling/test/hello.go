/*
 * @Descripttion: 
 * @version: 
 * @Author: 周宏敞
 * @Date: 2022-08-29 17:52:44
 * @LastEditors: 周宏敞
 * @LastEditTime: 2022-09-01 22:45:01
 */
package main

import (
	"fmt"
	"airpak"
)






func main(){
	fmt.Printf("Hello, world. 你好， 世界！\n")
	domain :=&airpak.Domain{
		Name: "room.1",
	}
	block :=airpak.Block{
		Name: "block.1",
		BlockType: airpak.FLUID,
		Point1: airpak.Point{X: -1., Y: 0., Z: 0.},
		Point2: airpak.Point{X: 0., Y: 2., Z: 1.},
	}
	wall :=airpak.Wall{
		Name: "wall.1",
		Point1: airpak.Point{X: -1., Y: 0., Z: 1.},
		Point2: airpak.Point{X: 0., Y: 2., Z: 1.},
	}
	open :=airpak.Opening{
		Name: "open.1",
	}


	//Airpak对象的切片
	Objects:=[] airpak.AirPaKObj{}
	Objects = append(Objects, domain)
	Objects = append(Objects, block)
	Objects = append(Objects, wall)
	Objects = append(Objects, open)

	for _, obj := range Objects{
		fmt.Printf("%s", obj.Text())
	}
	
}

//在每个文件夹下 go mod init 包名

//root文件夹下分别
//go work init .\airpak\
//go work use .\test\

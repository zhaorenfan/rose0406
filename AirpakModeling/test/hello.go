/*
 * @Descripttion:
 * @version:
 * @Author: 招人烦
 * @Date: 2022-08-29 17:52:44
 * @LastEditors: 周宏敞
 * @LastEditTime: 2022-09-03 00:19:19
 */
package main

import (
	"airpak"
	"fmt"
	"flag"
)

var(
	file string
	export string
)

func init(){
	flag.StringVar(&file, "i", "", "输入文件，默认为空")
	flag.StringVar(&export, "o", "", "输出文件，默认为空")
}

func main() {
	//需要先解析
	flag.Parse()
	if i:=flag.NFlag();i!=2{
		panic("参数不等于2")
	}

	if len(file)==0{
		panic("输入文件为空")
	}
	if len(export)==0{
		panic("输出文件为空")
	}

	fmt.Printf("Airpak3.0.16建模程序！\n")
	//file := "D:\\MYProjects\\Rhinos\\10-零碳\\block.txt"
	//export := "D:\\MYProjects\\Airpaks\\08-zero\\zero\\model"
	data := &airpak.Data{}
	airpak.Parse(data, file)
	


	fmt.Println("数据：")
	fmt.Println(len(data.Objs))
	objs := data.Objs
	for _, obj := range objs {
		fmt.Println(obj.Text())
	}
	data.ExportModel(export)

}

//在每个文件夹下 go mod init 包名

//root文件夹下分别
//go work init .\airpak\
//go work use .\test\

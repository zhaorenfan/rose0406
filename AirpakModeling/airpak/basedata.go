/*
 * @Descripttion: 
 * @version: 
 * @Author: 招人烦
 * @Date: 2022-09-02 19:40:11
 * @LastEditors: 招人烦
 * @LastEditTime: 2022-09-03 17:50:55
 */
package airpak

import (
	"fmt"
	"io"
	"os"
)

type Data struct {
	Objs []AirPaKObj
}

func (d *Data) AddObject(obj AirPaKObj) {

	d.Objs = append(d.Objs, obj)
}

// https://blog.csdn.net/weixin_48575553/article/details/124857550
func (d *Data) ExportModel(path string) error {
	fmt.Println("导出Airpak的model文件")
	var (
		file *os.File
		err  error
	)
	//文件是否存在
	if Exists(path) {
		//使用追加模式打开文件
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666) //os.O_APPEND
		if err != nil {
			fmt.Println("打开文件错误：", err)
			return err
		}
	} else {
		//不存在创建文件
		file, err = os.Create(path)
		if err != nil {
			fmt.Println("创建失败", err)
			return err
		}
	}

	//不要忘了文件延迟关闭
	defer file.Close()
	//写入文件
	head := "#@ Airpak 3.0.16 model file\n\n"
	n, err := io.WriteString(file, head)
	if err != nil {
		fmt.Println("写入错误：", err)
		return err
	}
	fmt.Println("写入成功：字符数n=", n)

	for _, obj := range d.Objs {
		n, err := io.WriteString(file, obj.Text()+"\n")
		if err != nil {
			fmt.Println("写入错误：", err)
			return err
		}
		fmt.Println("写入成功：字符数n=", n)
	}

	//读取文件
	// fileContent,err:=ioutil.ReadFile(path)
	// if err!= nil{
	// 	fmt.Println("读取错误：",err)
	// 	return err
	// }
	// fmt.Println("读取成功，文件内容：",string(fileContent))
	return err
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

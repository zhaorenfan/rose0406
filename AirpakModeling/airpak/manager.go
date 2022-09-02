package airpak

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Parse(d *Data, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
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
		}

	}
	return err
}

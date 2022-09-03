/*
 * @Descripttion:
 * @version:
 * @Author: 招人烦
 * @Date: 2022-09-01 17:56:51
 * @LastEditors: 招人烦
 * @LastEditTime: 2022-09-03 17:51:05
 */
package airpak

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	FLUID = iota
	SOLID
)

type Vector3d struct {
	X float32
	Y float32
	Z float32
}

func Vector3d2Str(v3 Vector3d) string {
	return strconv.FormatFloat(float64(v3.X), 'f', 3, 32) + " " +
		strconv.FormatFloat(float64(v3.Y), 'f', 3, 32) + " " +
		strconv.FormatFloat(float64(v3.Z), 'f', 3, 32)
}

type Point Vector3d

func Point2Arr(p Point) [3]float32 {
	return [3]float32{p.X, p.Y, p.Z}
}

func Point2Str(p Point) string {
	return Vector3d2Str(Vector3d(p))
}

func Str2Point(ctx string) Point {
	/* 字符串分割-多个分隔符
	s := "11111|2222||3333,4444,,555"
	seps := ",|"
	rs2 := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(seps, r)
		})
	fmt.Println(rs2)
	*/
	//{1,2,3}
	seps := ",{}"
	arr := strings.FieldsFunc(ctx, func(r rune) bool {
		return strings.ContainsRune(seps, r)
	})
	fmt.Println(arr)
	if l := len(arr); l != 3 {
		panic("字符串转三维点失败，请检查点坐标数量")
	}
	x, err := strconv.ParseFloat(arr[0], 32)
	if err != nil {
		panic("字符串转三维点失败，x坐标字符串转float32失败")
	}
	y, err := strconv.ParseFloat(arr[1], 32)
	if err != nil {
		panic("字符串转三维点失败，y坐标字符串转float32失败")
	}
	z, err := strconv.ParseFloat(arr[2], 32)
	if err != nil {
		panic("字符串转三维点失败，z坐标字符串转float32失败")
	}
	return Point{float32(x), float32(y), float32(z)}

}

func PointSub(p2, p1 Point) Vector3d {
	return Vector3d{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}
}

// 判断两个点在yz xz xy哪个平面上，返回值分别为0，1，2
// 否则，则返回-1
func WhichPlane(p1, p2 Point) int {
	diff := PointSub(p2, p1)
	dx := float64(diff.X)
	dy := float64(diff.Y)
	dz := float64(diff.Z)

	if math.Abs(dx) == 0 && math.Abs(dy) != 0 && math.Abs(dz) != 0 {
		return 0
	} else if math.Abs(dx) != 0 && math.Abs(dy) == 0 && math.Abs(dz) != 0 {
		return 1
	} else if math.Abs(dx) != 0 && math.Abs(dy) != 0 && math.Abs(dz) == 0 {
		return 2
	} else {
		return -1
	}

}

// Airpak的对象接口
type AirPaKObj interface {
	Text() string
}

// Airpak的对象结构
// domian
type Domain struct {
	Name string
}

func (d Domain) Text() string {
	return ("object domain " + d.Name + "\n" +
		"    current_genus default\n" +
		"    shape body_shape shape_none\n" +
		"        setval \n" +
		"    end shape\n" +
		"    current_stype none\n" +
		"    creation_order 1\n" +
		"    grid_priority 0\n" +
		"end object\n")
}

type Block struct {
	Name      string
	BlockType int //枚举
	Point1    Point
	Point2    Point
}

func (b Block) Text() string {

	sub := PointSub(b.Point2, b.Point1)
	diff := Vector3d2Str(sub)

	blockType := ""
	switch b.BlockType {
	case FLUID:
		blockType = "fluid"
	case SOLID:
		blockType = "solid"
	default:
		blockType = "fluid"
	}
	return "object block " + b.Name + "\n" +
		"    current_stype hexa\n" +
		"    block_type " + blockType + "\n" +
		"    shape body_shape shape_hexa\n" +
		"        setval point1 {" + Point2Str(b.Point1) + "} point2 {" + Point2Str(b.Point2) + "} diff {" + diff + "} volume_flag {1} diff_flag {0} \n" +
		"    end shape\n" +
		"    grid_priority 3\n" +
		"    creation_order 4\n" +
		"    current_genus default\n" +
		"end object\n"
}

type Wall struct {
	Name   string
	Point1 Point
	Point2 Point
}

func (w Wall) Text() string {
	sub := PointSub(w.Point2, w.Point1)
	diff := Vector3d2Str(sub)

	plane := WhichPlane(w.Point1, w.Point2)
	if plane == -1 {
		msg := "object name:" + w.Name + " " + "点坐标错误导致未发现在yz/xz/xy平面任意一个"
		panic(msg)
	}

	return "object wall " + w.Name + "\n" +
		"    shape body_shape shape_quad\n" +
		"        setval point1 {" + Point2Str(w.Point1) + "} point2 {" + Point2Str(w.Point2) + "} diff {" + diff + "} volume_flag {1} split_flag {0} plate_flag {1} diff_flag {0} plane {" + strconv.Itoa(plane) + "} iradius {0} thickness {0} \n" +
		"    end shape\n" +
		"    thermal_type temp\n" +
		"    grid_priority 4\n" +
		"    forced_flow_dir 0\n" +
		"    thermal_itemp 293.15\n" +
		"    current_genus default\n" +
		"    current_stype quad\n" +
		"    thermal_rtype reftemp\n" +
		"    creation_order 2\n" +
		"end object\n"
}

type Opening struct {
	Name string
	Point1 Point
	Point2 Point
}

func (o Opening) Text() string {

	sub := PointSub(o.Point2, o.Point1)
	diff := Vector3d2Str(sub)

	plane := WhichPlane(o.Point1, o.Point2)
	if plane == -1 {
		msg := "object name:" + o.Name + " " + "点坐标错误导致未发现在yz/xz/xy平面任意一个"
		panic(msg)
	}

	return "object opening " + o.Name + "\n" +
		"    xvecf 0\n" +
		"    current_stype quad\n" +
		"    shape body_shape shape_quad\n" +
		"        setval point1 {"+ Point2Str(o.Point1) +"} point2 {"+ Point2Str(o.Point2)+"} diff {"+ diff +"} volume_flag {0} split_flag {16} plate_flag {0} diff_flag {0} plane {"+ strconv.Itoa(plane) +"} iradius {0} thickness {0} \n" +
		"    end shape\n" +
		"    zvecf 1\n" +
		"    free_magnitude 1.0\n" +
		"    creation_order 5\n" +
		"    current_genus free\n" +
		"    yvecf 0\n" +
		"end object\n"
}

//棱柱
type Prism struct {
	Name string
	BlockType int //枚举
	Point1 Point
	Point2 Point
	Point3 Point
	Height float32
}

func (p Prism) Text() string {

	plane := WhichPlane(p.Point1, p.Point2)
	if plane == -1 {
		msg := "object name:" + p.Name + " " + "点坐标错误导致未发现在yz/xz/xy平面任意一个"
		panic(msg)
	}
	blockType := ""
	switch p.BlockType {
	case FLUID:
		blockType = "fluid"
	case SOLID:
		blockType = "solid"
	default:
		blockType = "fluid"
	}
	return  "object block "+p.Name+"\n" +
			"	current_stype polygon\n"+
			"	block_type "+blockType+"\n"+
			"	shape body_shape shape_polygon\n"+
			"		setval nverts 3\n"+
			"		setval volume_flag {1} split_flag {0} changes {0} nverts {3} plane {"+strconv.Itoa(plane)+"} height {"+strconv.Itoa(int(p.Height))+"}  vert1 {"+ Point2Str(p.Point1)+"} vert2 {"+ Point2Str(p.Point2)+"} vert3 {"+ Point2Str(p.Point3)+"}\n"+
			"	end shape\n"+
			"	creation_order 30\n"+
			"	current_genus default\n"+
			"end object\n"
}



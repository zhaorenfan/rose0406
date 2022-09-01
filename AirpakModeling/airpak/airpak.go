/*
 * @Descripttion: 
 * @version: 
 * @Author: 周宏敞
 * @Date: 2022-09-01 17:56:51
 * @LastEditors: 周宏敞
 * @LastEditTime: 2022-09-01 22:51:41
 */
package airpak

import (
	"strconv"
	"math"
)

const (
	FLUID = iota
	SOLID
)

type Vector3d struct{
	X float32
	Y float32
	Z float32
}

func Vector3d2Str(v3 Vector3d) string{
	return 	strconv.FormatFloat(float64(v3.X), 'f', 3, 32)+" "+
			strconv.FormatFloat(float64(v3.Y), 'f', 3, 32)+ " "+
			strconv.FormatFloat(float64(v3.Z), 'f', 3, 32)
}

type Point Vector3d

func Point2Arr(p Point) [3]float32{
	return [3]float32{p.X, p.Y, p.Z}
}

func Point2Str(p Point) string{
	return 	Vector3d2Str(Vector3d(p))
}

func PointSub(p2, p1 Point) Vector3d{
	return Vector3d{ p2.X - p1.X,p2.Y - p1.Y,p2.Z - p1.Z}
}

//判断两个点在yz xz xy哪个平面上，返回值分别为0，1，2 
//否则，则返回-1
func WhichPlane(p1, p2 Point) int {
	diff := PointSub(p2, p1)
	dx := float64(diff.X)
	dy := float64(diff.Y)
	dz := float64(diff.Z)

	if math.Abs(dx)==0 && math.Abs(dy)!=0 && math.Abs(dz)!=0{
		return 0
	} else if math.Abs(dx)!=0 && math.Abs(dy)==0 && math.Abs(dz)!=0{
		return 1
	} else if math.Abs(dx)!=0 && math.Abs(dy)!=0 && math.Abs(dz)==0{
		return 2
	} else {
		return -1
	}

}

//Airpak的对象接口
type AirPaKObj interface{
	Text() string
}

//Airpak的对象结构
//domian
type Domain struct{
	Name string
}

func (d Domain) Text() string{
	return ("object domain "+ d.Name + "\n"+
    		"    current_genus default\n"+
    		"    shape body_shape shape_none\n"+
        	"        setval \n"+
    		"    end shape\n"+
    		"    current_stype none\n"+
    		"    creation_order 1\n"+
    		"    grid_priority 0\n"+
			"end object\n")
}

type Block struct{
	Name string
	BlockType int    //枚举
	Point1 Point
	Point2 Point
}




func (b Block) Text() string{

	sub := PointSub(b.Point2, b.Point1)
	diff := Vector3d2Str(sub)

	blockType := ""
	switch b.BlockType{
	case FLUID:
		blockType = "fluid"
	case SOLID:
		blockType = "solid"
	default:
		blockType = "fluid"
	}
	return  "object block " + b.Name +"\n"+
			"    current_stype hexa\n"+
			"    block_type " + blockType +"\n"+
			"    shape body_shape shape_hexa\n"+
			"        setval point1 {" + Point2Str(b.Point1) + "} point2 {" + Point2Str(b.Point2) + "} diff {"+ diff  +"} volume_flag {1} diff_flag {0} \n"+
			"    end shape\n"+
			"    grid_priority 3\n"+
			"    creation_order 4\n"+
			"    current_genus default\n"+
			"end object\n"
}

type Wall struct{
	Name string
	Point1 Point
	Point2 Point
}

func (w Wall) Text() string{
	sub := PointSub(w.Point2, w.Point1)
	diff := Vector3d2Str(sub)

	plane := WhichPlane(w.Point1, w.Point2)
	if plane==-1{
		msg := "object name:"+ w.Name + " " + "点坐标错误导致未发现在yz/xz/xy平面任意一个"
		panic(msg)
	}

	return 	"object wall " + w.Name + "\n" +
			"    shape body_shape shape_quad\n" +
			"        setval point1 {"+ Point2Str(w.Point1) +"} point2 {"+ Point2Str(w.Point2) +"} diff {"+ diff + "} volume_flag {1} split_flag {0} plate_flag {1} diff_flag {0} plane {"+strconv.Itoa(plane)+"} iradius {0} thickness {0} \n"+
			"    end shape\n"+
			"    thermal_type temp\n"+
			"    grid_priority 4\n"+
			"    forced_flow_dir 0\n"+
			"    thermal_itemp 293.15\n"+
			"    current_genus default\n"+
			"    current_stype quad\n"+
			"    thermal_rtype reftemp\n"+
			"    creation_order 2\n"+
			"end object\n"
}

type Opening struct{
	Name string
}

func (o Opening) Text() string{
	return 	"object opening " + o.Name + "\n" +
			"    xvecf 0\n"+
			"    current_stype quad\n"+
			"    shape body_shape shape_quad\n"+
			"        setval point1 {-1 0 0} point2 {0 2 0} diff {1 2 0} volume_flag {0} split_flag {16} plate_flag {0} diff_flag {0} plane {2} iradius {0} thickness {0} \n"+
			"    end shape\n"+
			"    zvecf 1\n"+
			"    free_magnitude 1.0\n"+
			"    creation_order 5\n"+
			"    current_genus free\n"+
			"    yvecf 0\n"+
			"end object\n"
}
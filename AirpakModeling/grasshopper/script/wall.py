"""Provides a scripting component.
    Inputs:
        name: wall的name
        type: one of RECT、POLYGON、INCLINED
    Output:
        a: The a output variable
"""

__author__ = "zhaorenfan"
__version__ = "2022.09.02"

ghenv.Component.Name = "WALL"
ghenv.Component.NickName = "w"
ghenv.Component.Message = "VER 0.0.1\nJULY_2_2022"
ghenv.Component.Category = "ZRF"
ghenv.Component.SubCategory = "0 | ZRF"

import rhinoscriptsyntax as rs
import Rhino as rc
import ghpythonlib.components as gh
import math

if len(name)==0:
    name="wall"
a = []
pts = []
i = 1

if type == "RECT":
    for surf in surfs:
        _, _, Vt = gh.DeconstructBrep(surf)
        # print(Vt)
        p1 = Vt[0]
        
        xmin = p1[0]
        ymin = p1[1]
        zmin = p1[2]
        xmax = p1[0]
        ymax = p1[1]
        zmax = p1[2]
        for p in Vt:
            xmin = min(xmin, p[0])
            ymin = min(ymin, p[1])
            zmin = min(zmin, p[2])
            xmax = max(xmax, p[0])
            ymax = max(ymax, p[1])
            zmax = max(zmax, p[2])
        pmin = rc.Geometry.Point3d(xmin,ymin,zmin)
        pmax = rc.Geometry.Point3d(xmax,ymax,zmax)
        pts.append(pmin)
        pts.append(pmax)
        
        v = "WALL "+name+"."+str(i)+" "+ type
        i += 1
        v+=" PSTART {"+str(pmin)+"} "+"PEND {"+str(pmax)+"}"
        a.append(v)

if type=="POLYGON":
    for surf in surfs:
        #print(surf)
        F, E, Vt = gh.DeconstructBrep(surf)

        #print("111")
        print(rs.IsSurfaceTrimmed(surf))

        pt_str = ""
        si = 1   # 点的数量
        for p in Vt:
            pts.append(p)
            r = " VERT"+str(si)+" {"+str(p)+"}"
            pt_str+=r
            si+=1
        # print(pt_str)
        v = "WALL "+name+"."+str(i) + " " +type + " NVERTS "+str(si-1)
        i += 1
        v+=pt_str
        a.append(v)

if type == "INCLINED":
    for surf in surfs:
        _, _, Vt = gh.DeconstructBrep(surf)
        # print(Vt)
        pS = Vt[0]
        pE = Vt[0]
        nor = rs.SurfaceNormal(surf, pS)  #平面法向量
        print(nor)
        axis = 0
        angle = 0
        n0 = rc.Geometry.Vector3d(0,0,0)
        if abs(nor[0])<=1e-3:
            axis = 0
            # 与平面xy夹角(0,0,1)
            n0 = rc.Geometry.Vector3d(0,0,1)
            angle = rs.VectorAngle(n0, nor)
            print("aaaaa")

        elif abs(nor[1])<=1e-3:
            axis = 1
            n0 = rc.Geometry.Vector3d(0,0,1)
            angle = rs.VectorAngle(n0, nor)

        elif abs(nor[2])<=1e-3:
            axis = 2
            n0 = rc.Geometry.Vector3d(0,-1,0)
            angle = rs.VectorAngle(n0, nor)
        # angle = rs.VectorAngle(n0, nor)
        print("angle:", angle)
        print("axis:",axis)
        print(nor)
        for p in Vt:
            l1 = rs.VectorSubtract(p, pS)
            l2 = rs.VectorSubtract(pE, pS)
            # 这里有问题
            #if l1>l2:
            if rs.VectorLength(l1)>rs.VectorLength(l2):
                pE = p
        pts.append(pS)
        pts.append(pE)
        diff = rs.VectorSubtract(pE, pS)
        # setval point1 {0.05 0.22 1.78} point2 {1.28 1.45 3.0} diff {1.2 1.2 1.2} diff2 {1.2 1.7 0} volume_flag {1} 
        # split_flag {0} plate_flag {1} diff_flag {2} axis {0} rotate_sign {1} rotate_angle {45} thickness {0}
        v = "WALL "+name+"."+str(i)+" "+ type
        i += 1
        v+=" PSTART {"+str(pS)+"} "+"PEND {"+str(pE)+"}"
#       v+=" PSTART {"+str(pS)+"} "+"PEND {"+str(pE)+"}" + " DIFF {" + str(diff)+"}"
        v+=" AXIS "+str(axis)+" ANGLE "+ str(angle)
        a.append(v)

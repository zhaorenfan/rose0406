"""Provides a scripting component.
    Inputs:
        name: opening的名字
        surfs: opening的brep类型的表面
    Output:
        a: 输出自定义的文本
        pts: 参考点
"""

__author__ = "zhaorenfan"
__version__ = "2022.09.02"

ghenv.Component.Name = "OPENING"
ghenv.Component.NickName = "o"
ghenv.Component.Message = "VER 0.0.1\nSEPT_2_2022"
ghenv.Component.Category = "ZRF"
ghenv.Component.SubCategory = "0 | ZRF"

import rhinoscriptsyntax as rs
import Rhino as rc
import ghpythonlib.components as gh

if len(name)==0:
    name="opening"
a = []
pts = []
i = 1
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
    
    v = "OPENING "+name+"."+str(i)
    i += 1
    v+=" PSTART {"+str(pmin)+"} "+"PEND {"+str(pmax)+"}"
    a.append(v)

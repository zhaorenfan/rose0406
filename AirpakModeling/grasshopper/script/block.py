"""Provides a scripting component.
    Inputs:
        x: The x script variable
        y: The y script variable
    Output:
        a: The a output variable"""

__author__ = "zhaorenfan"
__version__ = "2022.09.02"

ghenv.Component.Name = "BLOCK"
ghenv.Component.NickName = "blk"
ghenv.Component.Message = "VER 0.0.1\nJULY_2_2022"
ghenv.Component.Category = "ZRF"
ghenv.Component.SubCategory = "0 | ZRF"

import rhinoscriptsyntax as rs
import Rhino as rc
import ghpythonlib.components as gh

def findPoints(Vertex):
    print(Vertex)
    xmin = Vertex[0][0]
    ymin = Vertex[0][1]
    zmin = Vertex[0][2]
    xmax = Vertex[0][0]
    ymax = Vertex[0][1]
    zmax = Vertex[0][2]
    for v in Vertex:
        xmin = min(xmin, v[0])
        ymin = min(ymin, v[1])
        zmin = min(zmin, v[2])
        
        xmax = max(xmax, v[0])
        ymax = max(ymax, v[1])
        zmax = max(zmax, v[2])
    pmin = rc.Geometry.Point3d(xmin,ymin,zmin)
    pmax = rc.Geometry.Point3d(xmax,ymax,zmax)
    return pmin, pmax


V = []
a = []
if len(name)==0:
    name = "block"
if len(type) == 0:
    type = "FLUID"

i = 1
print(len(breps))
for brep in breps:
    _, _, Vt = gh.DeconstructBrep(brep)
    print(brep)
    v = "BLOCK "+name+"."+str(i)+" "+type
    i+=1
    p1, p2 = findPoints(Vt)
    v+=" PSTART {"+str(p1)+"} "+"PEND {"+str(p2)+"}"
    a.append(v)
    
    for vertex in Vt:
        p = rc.Geometry.Point3d(vertex)
        V.append(p)


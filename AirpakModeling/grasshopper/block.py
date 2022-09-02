"""Provides a scripting component.
    Inputs:
        x: The x script variable
        y: The y script variable
    Output:
        a: The a output variable"""

__author__ = "ZRF"
__version__ = "2022.09.02"

ghenv.Component.Name = "BLOCK"
ghenv.Component.NickName = 'bck'
ghenv.Component.Message = 'VER 0.0.1\nJULY_2_2022'
ghenv.Component.Category = "ZRF"
ghenv.Component.SubCategory = "0 | ZRF"

import rhinoscriptsyntax as rs
# import scriptcontext as sc
import Rhino as rc
import ghpythonlib.components as gh

def findPoints(Vertex):
    xmin = 0
    ymin = 0
    zmin = 0
    xmax = 0
    ymax = 0
    zmax = 0
    for v in Vertex:
        xmin = min(xmin, v[0])
        ymin = min(ymin, v[1])
        zmin = min(zmin, v[2])
        
        xmax = max(xmax, v[0])
        ymax = max(ymax, v[1])
        zmax = max(zmax, v[2])
    pmin = rc.Geometry.Point3d(xmin, ymin, zmin)
    pmax = rc.Geometry.Point3d(xmax, ymax, zmax)
    return pmin, pmax

V = []
a = []
# print(breps)
nums = len(breps)
# print("numOfBreps:"+ str(nums))

name = "block"
type = "FLUID"


i = 1
#print(brep)
for brep in breps:
    _, _, Vt = gh.DeconstructBrep(brep)
    # print(len(Vt))
    v = "BLOCK "+ name + "."+str(i) + " "+ type
    i+=1
    p1, p2 = findPoints(Vt)
    print(p1, p2)
    # v += " PSTART {"+ str(Vt[0])+"} " +"PEND {"+str(Vt[5])+"}"
    v += " PSTART {"+ str(p1)+"} " +"PEND {"+str(p2)+"}"
    a.append(v)
    # print(type(V))
    
    for vertex in Vt:
        #print(vertex)
        p = rc.Geometry.Point3d(vertex)
        V.append(p)
        

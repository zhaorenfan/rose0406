"""Provides a scripting component.
    Inputs:
        x: The x script variable
        y: The y script variable
    Output:
        a: The a output variable"""

__author__ = "zhaorenfan"
__version__ = "2022.09.03"

ghenv.Component.Name = "PRISM"
ghenv.Component.NickName = "pr"
ghenv.Component.Message = "VER 0.0.1\nSEPT_3_2022"
ghenv.Component.Category = "ZRF"
ghenv.Component.SubCategory = "0 | ZRF"

import rhinoscriptsyntax as rs
import Rhino as rc
import ghpythonlib.components as gh




pts = []
a = []
Height = []
if len(name)==0:
    name = "prism"
if len(type) == 0:
    type = "FLUID"

i = 1
print(len(breps))
for brep in breps:
    F, E, V = gh.DeconstructBrep(brep)
    tri1 = ""
    tri2 = ""
    numOfTri = 0
    for f in F:
        print(f)
        if rs.IsSurfaceTrimmed(f):
            if tri1=="":
                tri1 = f
            else:
                tri2 = f
            numOfTri += 1
    print(numOfTri)
    print(tri1)
    print(tri2)
    _, _, Pts1 = gh.DeconstructBrep(tri1)
    for p in Pts1:
        pts.append(p)
    _, _, Pts2 = gh.DeconstructBrep(tri2)
    p1 = Pts1[0]
    p2 = ""
    for p in Pts2:
        n1 = 0
        n2 = 0
        for x in range(3):
            p1x = p1[x]
            p2x = p[x]
            if p1x==p2x:
                n1+=1
            else:
                n2+=1
        if n1 == 2 and n2 ==1:
            p2 = p
            pts.append(p2)
            break
    print(p1)
    print(p2)
    d = rs.VectorSubtract(p2, p1)
    print(d)
    h = d[0]+d[1]+d[2]
    Height.append(h)
    
    # PRISM prism.1 FLUID VERT1 {-4,3,0} VERT2 {0,3,6} VERT3 {0,3,6} HEIGHT 1
    v = "PRISM "+name+"."+str(i)+" "+type
    i+=1
    v+=" VERT1 {"+str(Pts1[0])+"} "+"VERT2 {"+str(Pts1[1])+"}"+" VERT2 {"+str(Pts1[2])+"}"
    v+=" HEIGHT "+str(h)
    a.append(v)
    


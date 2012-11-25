package hello

import (
	"fmt"
    "net/http"
)

type Color int;
const (
    red = iota   
    green
    black
	blue
	yellow
	orange
	invalid
)

type namedColor struct {
	color Color
	name string
}

var colorNames = [7]namedColor {
	{red, "Red"},
	{green, "Green"},
	{black, "Black"},
	{blue, "Blue"},
	{yellow, "Yellow"},
	{orange, "Orange"},
	{invalid, "Invalid"},
}

func (color Color) getColorName() string {
	return colorNames[color].name
}

type Side int
const (
	frontSide = 0
	backSide = 1
	topSide = 2
	bottomSide = 3
	rightSide = 4
	leftSide =5
	noSide = 6
)

type namedSide struct {
	s Side
	name string
}

var sideNames = [6]namedSide {
	{frontSide, "Front"},
	{backSide, "Back"},
	{topSide, "Top"},
	{bottomSide, "Bottom"},
	{rightSide, "Right"},
	{leftSide, "Left"},
}

func (s Side) getSideName() string {
	return sideNames[s].name
}

const (
	front = 0
	back = 1
	top = 0
	bottom = 1
	right = 0
	left = 1
)

type CubeletID int
type CubeletColors [6]Color // order: front, back, top, bottom, right, left

type Cubelet struct {
	cid CubeletID "ID"
	colors CubeletColors
}

func (c * Cubelet) rotate(s Side) {
	cc := &c.colors
	
	if s == frontSide {
		cc[rightSide], cc[bottomSide], cc[leftSide], cc[topSide] = cc[topSide], cc[rightSide], cc[bottomSide], cc[leftSide]
	} else if s == backSide {
		cc[leftSide], cc[topSide], cc[rightSide], cc[bottomSide] = cc[topSide], cc[rightSide], cc[bottomSide], cc[leftSide]
	} else if s == topSide {
		cc[leftSide], cc[backSide], cc[rightSide], cc[frontSide] = cc[frontSide], cc[leftSide], cc[backSide], cc[rightSide]
	} else if s == bottomSide {
		cc[rightSide], cc[frontSide], cc[leftSide], cc[backSide] = cc[frontSide], cc[leftSide], cc[backSide], cc[rightSide]
	} else if s == rightSide {
		cc[backSide], cc[bottomSide], cc[frontSide], cc[topSide] = cc[topSide], cc[backSide], cc[bottomSide], cc[frontSide]
	} else if s == leftSide {
		cc[frontSide], cc[topSide], cc[backSide], cc[bottomSide] = cc[topSide], cc[backSide], cc[bottomSide], cc[frontSide]
	}
}

type Cube [2][2][2]Cubelet


func init() {
    http.HandleFunc("/", handler)
}
	
func initCube(w http.ResponseWriter, c * Cube) {
	fmt.Fprintf(w, "Init Cube\n")

	// CubeletColors order: front, back, top, bottom, right, left

 	c[front][top][right] = Cubelet { 0, CubeletColors{blue, invalid, black, invalid, orange, invalid} }
 	c[back][top][right] = Cubelet { 1, CubeletColors{invalid, green, black, invalid, orange, invalid} }
 	c[front][bottom][right] = Cubelet { 2, CubeletColors{blue, invalid, invalid, yellow, orange, invalid} }
 	c[back][bottom][right] =  Cubelet { 3, CubeletColors{invalid, green, invalid, yellow, orange, invalid} }
 	c[front][top][left] =  Cubelet { 4, CubeletColors{blue, invalid, black, invalid, invalid, red} }
 	c[back][top][left] =  Cubelet { 5, CubeletColors{invalid, green, black, invalid, invalid, red} }
 	c[front][bottom][left] =  Cubelet { 6, CubeletColors{blue, invalid, invalid, yellow, invalid, red} }
 	c[back][bottom][left] =  Cubelet { 7, CubeletColors{invalid, green, invalid, yellow, invalid, red} }
}

func rotate(w http.ResponseWriter, c * Cube, s Side, t int) {
	t = t % 4
	
	if t < 0 {
		t = t + 4
	}
	
	if t == 0 {
		return
	}
	
	fmt.Fprintf(w, "Rotate %v\n", s.getSideName() )
	
	var cubelets []*Cubelet
	cubelets = getCubeSide(w, c, s)
	
	*cubelets[0], *cubelets[1], *cubelets[2], *cubelets[3] = *cubelets[3], *cubelets[0], *cubelets[1], *cubelets[2]
	
	for _, cubelet := range cubelets {
		cubelet.rotate(s)
	}
	
	rotate(w, c, s, t-1)
}

func printCubelets(w http.ResponseWriter, cubelets []*Cubelet) {
	fmt.Fprintf(w, "printCubelets\n")
	fmt.Fprintf(w, "Length: %v\n", len(cubelets))
	for i, cubelet := range cubelets {
		fmt.Fprintf(w, "Cubelet	%v:	%v\n", i, cubelet)
	}
}

func printCube(w http.ResponseWriter, c Cube) {
	fmt.Fprintf(w, "[front][top][right]:	%v\n",	c[front][top][right])
	fmt.Fprintf(w, "[back][top][right]:	%v\n", 		c[back][top][right])
	fmt.Fprintf(w, "[front][bottom][right]:	%v\n",	c[front][bottom][right])
	fmt.Fprintf(w, "[back][bottom][right]:	%v\n",	c[back][bottom][right])
	fmt.Fprintf(w, "[front][top][left]:	%v\n",		c[front][top][left])
	fmt.Fprintf(w, "[back][top][left]:	%v\n",		c[back][top][left])
	fmt.Fprintf(w, "[front][bottom][left]:	%v\n",	c[front][bottom][left])
	fmt.Fprintf(w, "[back][bottom][left]:	%v\n",	c[back][bottom][left])
}

func getCubeletColor(w http.ResponseWriter, c Cubelet, s Side) Color {
	return c.colors[s]
}

func getFaceColors(w http.ResponseWriter, c * Cube, s Side) [4]Color {
	
	cubeSide := getCubeSide(w, c, s)
	
	var result [4]Color
	
	for i, c := range cubeSide {
		result[i] = getCubeletColor(w, *c, s) 
	}

	return result
}

func getCubeSide(w http.ResponseWriter, c * Cube, s Side) []*Cubelet {
	
	var cubeSide []*Cubelet
	
	if(s==frontSide) {
		cubeSide = append(cubeSide, &c[front][top][right])
		cubeSide = append(cubeSide, &c[front][bottom][right])
		cubeSide = append(cubeSide, &c[front][bottom][left])
		cubeSide = append(cubeSide, &c[front][top][left])
	} else if(s==backSide) {
		cubeSide = append(cubeSide, &c[back][top][left])
		cubeSide = append(cubeSide, &c[back][bottom][left])
		cubeSide = append(cubeSide, &c[back][bottom][right])
		cubeSide = append(cubeSide, &c[back][top][right])
	} else if(s==topSide) {
		cubeSide = append(cubeSide, &c[back][top][right])
		cubeSide = append(cubeSide, &c[front][top][right])
		cubeSide = append(cubeSide, &c[front][top][left])	
		cubeSide = append(cubeSide, &c[back][top][left])	
	} else if(s==bottomSide) {
		cubeSide = append(cubeSide, &c[front][bottom][right])
		cubeSide = append(cubeSide, &c[back][bottom][right])
		cubeSide = append(cubeSide, &c[back][bottom][left])
		cubeSide = append(cubeSide, &c[front][bottom][left])
	} else if(s==rightSide) {
		cubeSide = append(cubeSide, &c[back][top][right])
		cubeSide = append(cubeSide, &c[back][bottom][right])
		cubeSide = append(cubeSide, &c[front][bottom][right])
		cubeSide = append(cubeSide, &c[front][top][right])
	} else if(s==leftSide) {
		cubeSide = append(cubeSide, &c[front][top][left])
		cubeSide = append(cubeSide, &c[front][bottom][left])
		cubeSide = append(cubeSide, &c[back][bottom][left])
		cubeSide = append(cubeSide, &c[back][top][left])
	}
	
	return cubeSide	
}

func printFaceColors(w http.ResponseWriter, c Cube, s Side) {
	for _, c := range getFaceColors(w, &c, s) {
		fmt.Fprintf(w, "%v\n", c.getColorName())
	}	
}

func printAllFaces(w http.ResponseWriter, c Cube) {
	var sideColors [6][4]Color
	
	for i, side := range sideNames {
		sideColors[i] = getFaceColors(w, &c, side.s)
	}
	
	for _, side := range sideNames {
		fmt.Fprintf(w, "%v\t", side.s.getSideName())
	}
	fmt.Fprintf(w, "\n")
	
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Fprintf(w, "%v\t", sideColors[j][i].getColorName())
		}
		fmt.Fprintf(w, "\n")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var c Cube
	
	initCube(w, &c)
	printAllFaces(w, c)
	
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)
	rotate(w, &c, backSide, 1)
	rotate(w, &c, rightSide, 1)

	printAllFaces(w, c)
}

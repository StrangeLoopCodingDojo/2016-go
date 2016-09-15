/*
   Copyright 2016 Google Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package minesweeper

type Grid struct {
	cellMap [][]bool
}

func (g *Grid) NeighboringMinesCount(x int, y int) int {
	i := 0

	for xd := -1; xd < 2; xd++ {
		for yd := -1; yd < 2; yd++ {

			if g.HasMineAt(x+xd, y+yd) {
				i++
			}
		}
	}

	return i
}

func (g *Grid) PlaceMine(x int, y int) {
	g.cellMap[x][y] = true
}

func (g *Grid) HasMineAt(x int, y int) bool {
	if x < 0 || y < 0 || len(g.cellMap[0]) <= y || len(g.cellMap) <= x {
		return false
	}

	return g.cellMap[x][y]
}

func NewGrid(width int, height int) *Grid {
	cellMap := [][]bool{}
	for i := 0; i < width; i++ {
		cellMap = append(cellMap, make([]bool, height))
	}

	return &Grid{cellMap: cellMap}
}

func DoThing() bool {
	return true
}

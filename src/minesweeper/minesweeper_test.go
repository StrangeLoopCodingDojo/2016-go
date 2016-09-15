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

import "testing"

func TestScenario1(t *testing.T) {
	grid := NewGrid(8, 8)
	grid.PlaceMine(5, 2)

	if !grid.HasMineAt(5, 2) {
		t.Fatalf("there is no mine")
	}

	if grid.HasMineAt(5, 1) {
		t.Fatalf("there is a mine")
	}

	// Verify: 4,1 says that there is a mine next to it.

	if grid.NeighboringMinesCount(0, 0) != 0 {
		t.Fatalf("there is a mine")
	}

	if grid.NeighboringMinesCount(7, 7) != 0 {
		t.Fatalf("there is a mine")
	}

	if nc := grid.NeighboringMinesCount(4, 1); nc != 1 {
		t.Fatalf("Neighboring count is: %#v doesn't = 1", nc)
	}

	if nc := grid.NeighboringMinesCount(4, 2); nc != 1 {
		t.Fatalf("Neighboring count is: %#v doesn't = 1", nc)
	}

	if nc := grid.NeighboringMinesCount(5, 1); nc != 1 {
		t.Fatalf("Neighboring count is: %#v doesn't = 1", nc)
	}

}

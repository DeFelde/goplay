package main

import (
	"testing"
)

func TestPuzzleSum(t *testing.T) {
	sum := puzzleSum(agent)
	if sum != 1203 {
		t.Errorf("Got %d but expected 1203\n", sum)
	}
}

func Test2PuzzleSum(t *testing.T) {
	sum := puzzleSum("553")
	if sum != 5 {
		t.Errorf("Got %d but expected 5\n", sum)
	}
}

func Test3PuzzleSum(t *testing.T) {
	sum := puzzleSum("1551")
	if sum != 6 {
		t.Errorf("Got %d but expected 6\n", sum)
	}
}

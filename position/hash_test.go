package position

import (
	"testing"
)

type positionHashTest struct {
	x, y, z int
}

var positionHashTests = []positionHashTest{
	positionHashTest{6543, 6543, 1234},
	positionHashTest{11, 22, 35},
}

var savedHashes map[int64]Position

func TestHash(t *testing.T) {
	t.Log("Testing predefined hashes")
	for _, ht := range positionHashTests {
		v := Hash(ht.x, ht.y, ht.z)
		pos := NewPositionFromHash(v)

		if pos.X != ht.x || pos.Y != ht.y || pos.Z != ht.z {
			t.Errorf("Failed to convert back (%d, %d, %d) to (%d, %d, %d)", pos.X, pos.Y, pos.Z, ht.x, ht.y, ht.z)
		}
	}
}

func TestDuplicated(t *testing.T) {
	savedHashes = make(map[int64]Position)
	t.Log("Duplicate Test")
	for x := -100; x <= 100; x++ {
		for y := -100; y <= 100; y++ {
			pos := NewPositionFrom(x, y, 1)

			if saved, ok := savedHashes[pos.Hash()]; ok {
				t.Errorf("Found duplicate for (%v). Have (%v)\n", pos.String(), saved.String())
			} else {
				savedHashes[pos.Hash()] = pos
			}
		}
	}
}

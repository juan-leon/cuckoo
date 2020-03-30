// Copyright (c) 2014-2015 Utkan Güngördü <utkan@freeconsole.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cuckoo

import (
	"math"
	"runtime"
	"testing"
)

var n = int(2e6) // close enough to a power of 2, to test whether the LoadFactor is close to 1 or not.

var (
	gkeys   []Key
	logsize = int(math.Ceil(math.Log2(float64(n))))
)

var (
	mapBytes    uint64
	cuckooBytes uint64
)

func readAlloc() uint64 {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return ms.Alloc
}

func mkmap(n int) ([]Key, uint64) {

	keys := make([]Key, n, n)

	runtime.GC()
	before := readAlloc()
	after := readAlloc()

	return keys, after - before
}

func init() {
	gkeys, mapBytes = mkmap(n)
}

func TestZero(t *testing.T) {
	c := NewCuckoo(logsize)

	for i := 0; i < 10; i++ {
		c.Insert(0)
		ok := c.Search(0)
		if !ok {
			t.Error("search failed")
		}
	}
}

func TestSimple(t *testing.T) {
	c := NewCuckoo(DefaultLogSize)
	for _, k := range gkeys {
		c.Insert(k)
	}

	for _, k := range gkeys {
		ok := c.Search(k)
		if !ok {
			t.Error("not ok:", k)
			return
		}
	}

	if c.Len() != len(gkeys) {
		t.Error("got: ", c.Len(), " expected: ", len(gkeys))
		return
	}

	ndeleted := 0
	maxdelete := len(gkeys) * 95 / 100
	for _, k := range gkeys {
		if ndeleted >= maxdelete {
			break
		}

		c.Delete(k)
		if ok := c.Search(k); ok == true {
			t.Error("found: ", k)
			return
		}

		ndeleted++

		if c.Len() != len(gkeys)-ndeleted {
			t.Error("got: ", c.Len(), " expected: ", len(gkeys)-ndeleted)
			return
		}
	}
}

func TestMem(t *testing.T) {
	runtime.GC()
	before := readAlloc()

	c := NewCuckoo(logsize)
	for _, k := range gkeys {
		c.Insert(k)
	}

	after := readAlloc()

	cuckooBytes = after - before

	t.Log("LoadFactor:", c.LoadFactor())
	t.Log("Built-in map memory usage (MiB):", float64(mapBytes)/float64(1<<20))
	t.Log("Cuckoo hash  memory usage (MiB):", float64(cuckooBytes)/float64(1<<20))
}

# cuckoo
<img src="http://upload.wikimedia.org/wikipedia/commons/thumb/1/16/NeomorphusSalviniSmit.jpg/220px-NeomorphusSalviniSmit.jpg"></img>

Note: this fork makes the filter to act like a Set (as opposed to a map that can
store values).  Some of the documentation below (those that refer to values), is
obsoleted

Package `cuckoo` implements d-ary bucketized [cuckoo hashing](http://en.wikipedia.org/wiki/Cuckoo_hashing) with stash (bucketized cuckoo hashing is also known as splash tables).
This implementation uses configurable number of hash functions and cells per bucket.
Greedy algorithm for collision resolution is a random walk.

## Purpose
Cuckoo is a memory-efficient alternative to the built-in `map[Key]Value` type (where Key is an integer type and Value can be any type) with zero per-item overhead.
Hence, the memory-efficiency is equal to the [load factor](http://en.wikipedia.org/wiki/Hash_table#Key_statistics) (which can be as high as 99%).

## Usage
After cloning the repository, modify the definitions of `Key` and `Value` types to fit your needs. For optimal performance, you should also experiment with the fine-grade parameters of the algorithm listed in `config.go`.

## Documentation
[godoc](http://godoc.org/github.com/salviati/cuckoo)

## License
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see [http://www.gnu.org/licenses/](http://www.gnu.org/licenses/).

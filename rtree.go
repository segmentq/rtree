package rtree

import (
	"github.com/tidwall/rtree"
)

type IndexType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type ValueType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// OneD is an RTree of one ValueType that only takes a single index
type OneD[I IndexType, T ValueType] struct {
	base rtree.RTreeGN[I, T]
}

// NewOneD creates a new instance of OneD
func NewOneD[I IndexType, T ValueType]() OneD[I, T] {
	return OneD[I, T]{}
}

func (tr OneD[I, T]) inTwoD(index I) [2]I {
	return [2]I{0, index}
}

func (tr OneD[I, T]) inOneD(index [2]I) I {
	return index[0]
}

// Insert an item into the structure
func (tr *OneD[I, T]) Insert(min, max I, data T) {
	tr.base.Insert(tr.inTwoD(min), tr.inTwoD(max), data)
}

// Delete an item from the structure
func (tr *OneD[I, T]) Delete(min, max I, data T) {
	tr.base.Delete(tr.inTwoD(min), tr.inTwoD(max), data)
}

// Replace an item in the structure. This is effectively just a Delete
// followed by an Insert. But for some structures it may be possible to
// optimize the operation to avoid multiple passes
func (tr *OneD[I, T]) Replace(
	oldMin, oldMax I, oldData T,
	newMin, newMax I, newData T,
) {
	tr.base.Replace(
		tr.inTwoD(oldMin), tr.inTwoD(oldMax), oldData,
		tr.inTwoD(newMin), tr.inTwoD(newMax), newData,
	)
}

// Search the structure for items that intersects the rect param
func (tr *OneD[I, T]) Search(
	min, max I,
	iter func(min, max I, data T) bool,
) {
	tr.base.Search(tr.inTwoD(min), tr.inTwoD(max), func(min, max [2]I, data T) bool {
		return iter(tr.inOneD(min), tr.inOneD(max), data)
	})
}

// Scan iterates through all data in tree in no specified order.
func (tr *OneD[I, T]) Scan(iter func(min, max I, data T) bool) {
	tr.base.Scan(func(min, max [2]I, data T) bool {
		return iter(tr.inOneD(min), tr.inOneD(max), data)
	})
}

// Len returns the number of items in tree
func (tr *OneD[I, T]) Len() int {
	return tr.base.Len()
}

// Bounds returns the minimum bounding box
func (tr *OneD[I, T]) Bounds() (min, max I) {
	bmin, bmax := tr.base.Bounds()
	return tr.inOneD(bmin), tr.inOneD(bmax)
}

// Nearby performs a kNN-type operation on the index
func (tr *OneD[I, T]) Nearby(
	algo func(min, max I, data T, item bool) (dist float64),
	iter func(min, max I, data T, dist float64) bool,
) {
	falgo := func(min, max [2]I, data T, item bool) (dist float64) {
		return algo(tr.inOneD(min), tr.inOneD(max), data, item)
	}
	fiter := func(min, max [2]I, data T, dist float64) bool {
		return iter(tr.inOneD(min), tr.inOneD(max), data, dist)
	}
	tr.base.Nearby(falgo, fiter)
}

// Copy the tree.
// This is a copy-on-write operation and is very fast because it only performs
// a shadowed copy.
func (tr *OneD[I, T]) Copy() *OneD[I, T] {
	return &OneD[I, T]{base: *tr.base.Copy()}
}

// Clear will delete all items.
func (tr *OneD[I, T]) Clear() {
	tr.base.Clear()
}

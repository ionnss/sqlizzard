package main

import "encoding/binary"

// assert is a function that is used to assert a condition
func assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}

// BNode is a node in the B-tree - a slice of bytes
// all information for a node (header, pointers, offsets, key values) would be stored in this byte slice
type BNode struct {
	data []byte // can be dumped to the disk
}

// BNode types
const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

// BTree is a B-tree
type BTree struct {
	// pointer to the root node page number
	// each page on a disk is 4 KB in size
	// the number root points to a specific page on the disk where the root node is stored
	root uint64

	// callback functionsfor managing on-disk pages
	// get is a function that takes a page number and returns a BNode
	get func(uint64) BNode // dereference a pointer

	// new is a function that takes a BNode and returns a page number
	new func(BNode) uint64 // allocate a new page

	// del is a function that takes a page number and deallocates it
	del func(uint64) // deallocate a page
}

// constants
const HEADER = 4                  // 4 bytes
const BTREE_PAGE_SIZE = 4096      // 4 KB
const BTREE_MAX_KEY_SIZE = 1000   // 1000 bytes 1KB
const BTREE_MAX_VALUE_SIZE = 3000 // 3000 bytes 3 KB

// init is a function that is called when the program starts
// it is used to initialize the BTree
func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VALUE_SIZE
	assert(node1max <= BTREE_PAGE_SIZE)
}

// header
// btype is a function that returns the type of the node
func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data)
}

// nkeys is a function that returns the number of keys in the node
func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[2:4])
}

// setHeader is a function that sets the header of the node
func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

// pointers
func (node BNode) getPtr(idx uint16) uint64 {
	assert(idx < node.nkeys())
	pos := HEADER + 8*idx
	return binary.LittleEndian.Uint64(node.data[pos:])
}

func (node BNode) setPtr(idx uint16, val uint64) {
	assert(idx < node.nkeys())
	pos := HEADER + 8*idx
	binary.LittleEndian.PutUint64(node.data[pos:], val)
}

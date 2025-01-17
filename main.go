package main

// BNode is a node in the B-tree - a slice of bytes
// All information for a node (header, pointers, offsets, key values) would be stored in this byte slice
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

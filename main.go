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
	// pointer (a nonzero page number)
	root uint64
	// callback for managing on-disk pages
	get func(uint64) BNode // deference a pointer
	new func(BNode) uint64 // allocate a new page
	del func(uint64)       // deallocate a page
}

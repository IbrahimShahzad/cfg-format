package grammar

// #cgo CFLAGS: -std=c11 -fPIC -fcommon
// #include "include.h"
import "C"
import "unsafe"

// Language returns the tree-sitter language for Kamailio cfg files.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_kamailio_cfg())
}

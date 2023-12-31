// Package file verifies that a file node is generated by the indexer, and that
// it has the correct relationship to its package.
//
// - File=vname("", "test", "", "basic/filenode.go", "").node/kind file
// - @file defines/binding Pkg
package file

//- Pkg=vname("package", "test", "", "basic", "go").node/kind package
//- File childof Pkg

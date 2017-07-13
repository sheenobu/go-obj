# go-obj

OBJ file loader

Currently supported fields:

 * \# - comments , ignored
 * o - Object Name??
 * v - Vertex
 * vn - Vertex Normal
 * f - Face
 * vt - vertex texture coordinate indices

Everything else is silently ignored

## Usage

Simply `go get github.com/sheenobu/go-obj/obj`.

Much of the code outside of `go-obj/obj`
relies on SDL2 and vendored code but `go-obj/obj` should
be generic and never fail to pull due to Cgo dependencies #5.

## cmd/obj-renderer

This is a standard object renderer, using a simple GLSL shader (embedded) for lighting.

Usage:

	$ obj-renderer <filename>

## TODO

 * obj.Writer interface
 * The gometalinter says all the table based tests are the same. Try to abstract them?
 * Materials aren't supported.
 * Logging


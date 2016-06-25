# go-obj

OBJ file loader

Currently supported fields:

 * # - comments , ignored
 * o - Object Name??
 * v - Vertex
 * vn - Vertex Normal
 * f - Face
 * vt - vertex texture coordinate indices

Everyting else is silently ignored

## TODO

 * Renderers
   * OpenGL
   * ???
 * obj.Writer interface
 * The gometalinter says all the table based tests are the same. Try to abstract them?
 * Materials aren't supported.
 * Logging

## Notes

### Byte Array Usages

The majority of the obj.Reader is designed to use byte arrays, not strings. This is to keep
the memory footprint low as each line is considered one contigious region of memory
and never copied. In my opinion, it's a premature optimization but a good experiment.

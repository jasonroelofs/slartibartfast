// Wrapper package around the details of the mathgl library I'm using
// Basically I don't want to leak implementation details of *any* library I'm using, and
// building my own types here gives me the option of writing my own or adding to
// the underlying mathgl.
//
// I expect the go compiler will only get better at inlining so not worried about any
// function call overhead.
//
// Sorry about the package name, naming it just "math" means that Go pulls in it's
// own "math" package instead
package math3d

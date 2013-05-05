// My own implementation of common structures for 3d math.
// Why build my own? I'm only doing this because all of the others out there already
// written in Go are completely untested, and the last thing I want to do is have to
// track down an errant sign or number in someone elses code that's breaking Matrix
// or Vector math in my app.
//
// And yes there is a bit of NIH here, I also don't want to have someone else's math library
// strewn throughout the code. Also a lot of the packages include 2x2, 3x3, and off-size Matricies
// as well as multiple different sizes of Vector, and I don't really need all that.
//
// I'm mainly using github.com/Jragonmiris/mathgl as the reference for building this package.
//
// Sorry about the package name, naming it just "math" means that Go pulls in it's
// own "math" package instead
package math3d

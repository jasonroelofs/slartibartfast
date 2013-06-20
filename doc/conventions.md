Dependencies
------------

`go get` the dependency, add it to the appropriate array in the Rakefile, then `git submodule add [url] [dir]` to ensure
the version is saved in git.

Shader Conventions
------------------

Reference by a single name, pick up from data/shaders/[name].vert and data/shaders/[name].frag

Always a vert/frag pair until I find a situation where I want just one or the other, or some sort of
mix-and-match.

### Uniforms

The following uniform names will be automatically assigned the appropriate values
if found in the shaders currently in use:

* uniform mat4 projection;
* uniform mat4 view;
* uniform mat4 model;
* uniform mat4 viewProjection;
* uniform mat4 modelViewProjection;


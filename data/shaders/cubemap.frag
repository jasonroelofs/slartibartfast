#version 150

in vec3 cubeUVCoords;
out vec4 fragColor;

uniform samplerCube cubemap;

void main() {
  fragColor = texture(cubemap, cubeUVCoords);
}

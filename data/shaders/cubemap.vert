#version 410

in vec3 vertexPosition;
out vec3 cubeUVCoords;

uniform mat4 modelViewProjection;

void main() {
	gl_Position = modelViewProjection * vec4(vertexPosition, 1.0);
  cubeUVCoords = vertexPosition;
}

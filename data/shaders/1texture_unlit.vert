#version 150

in vec3 vertexPosition;  // Attr 0
in vec2 vertUVCoords;    // Attr 1
out vec2 fragUVCoords;

uniform mat4 modelViewProjection;

void main() {
	gl_Position = modelViewProjection * vec4(vertexPosition, 1.0);
  fragUVCoords = vertUVCoords;
}

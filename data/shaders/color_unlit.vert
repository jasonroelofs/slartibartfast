#version 150

in vec3 vertexPosition;  // Attr 0
in vec3 inColor;         // Attr 1
out vec3 vertColor;

uniform mat4 modelViewProjection;

void main() {
	gl_Position = modelViewProjection * vec4(vertexPosition, 1.0);
	vertColor = inColor;
}

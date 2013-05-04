#version 150

in vec3 vertexPosition;
in vec3 in_color;
out vec3 vert_color;

uniform mat4 modelViewProjection;

void main() {
	gl_Position = modelViewProjection * vec4(vertexPosition, 1.0);
	vert_color = in_color;
}

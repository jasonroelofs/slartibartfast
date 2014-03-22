#version 330

layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec3 inColor;
out vec3 vertColor;

uniform mat4 modelViewProjection;

void main() {
	gl_Position = modelViewProjection * vec4(vertexPosition, 1.0);
	vertColor = inColor;
}

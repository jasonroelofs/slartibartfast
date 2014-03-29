#version 410

layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec2 vertUVCoords;
out vec2 fragUVCoords;

uniform mat4 modelViewProjection;

void main() {
	gl_Position = modelViewProjection * vec4(vertexPosition, 1.0);
  fragUVCoords = vertUVCoords;
}

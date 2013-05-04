#version 150

in vec3 vert_color;
out vec4 frag_color;

void main() {
	frag_color = vec4(vert_color, 0);
}

#version 410

in vec2 fragUVCoords;
out vec3 fragColor;

uniform sampler2D textureSampler;

void main() {
  fragColor = texture(textureSampler, fragUVCoords).rgb;
}

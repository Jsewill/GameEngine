#version 330

uniform mat4 Projection;
uniform mat4 View;
uniform mat4 Translation;
uniform mat4 Rotation;
uniform mat4 Scale;

in vec3 position;
//in vec4 color;
in vec3 normal;

out vec4 fPosition;
out vec4 fColor;
out vec4 fNormal;
out vec4 fView;

void main() {
	mat4 MVP = Projection*View*Translation*Rotation*Scale;
	fColor = vec4(position, 1.0);
	fNormal = MVP*vec4(normal, 1.0);
	fPosition = MVP*vec4(position, 1.0);
	fView = -fPosition;
	gl_Position = fPosition;
}
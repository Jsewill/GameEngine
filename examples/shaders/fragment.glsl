#version 330

in vec4 fPosition;
in vec4 fColor;
in vec4 fNormal;
in vec4 fView;

out vec4 color;
 
void main() {
	float fresnel = normalize(1.0 - max(dot(fView, fNormal), 0.0));
	color = mix(vec4(vec3(0.0), 1.0), normalize(vec4(1.0)-(2.0*fColor)), 1.0);//vec4(normalize(gl_FragCoord.st), 1.0, 1.0);
}
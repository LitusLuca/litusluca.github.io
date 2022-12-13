#type vertex
#version 300 es

layout(location=0) in vec3 a_Pos;

void main() {
    gl_Position = vec4(a_Pos, 1.0);
}

#type fragment
#version 300 es

precision mediump float;

out vec4 FragmentColor;

void main() {
    FragmentColor = vec4(1.0);
}
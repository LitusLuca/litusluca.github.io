#type vertex
#version 300 es

layout(location=0) in vec3 a_Pos;
layout(location=1) in vec2 a_Color;

out vec2 vColor;

uniform mat4 uViewProjection;
uniform mat4 uModel;

void main() {
    vColor = a_Color;
    gl_Position = uViewProjection * uModel * vec4(a_Pos, 1.0);
}

#type fragment
#version 300 es

precision mediump float;

in vec2 vColor;

uniform sampler2D tex;

out vec4 FragmentColor;

void main() {
    //FragmentColor = vec4(vColor,1.0, 1.0);
    FragmentColor = texture(tex, vColor);
}
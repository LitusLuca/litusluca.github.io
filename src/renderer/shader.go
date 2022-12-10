package renderer

import "github.com/litusluca/litusluca.github.io/src/glapi"

type ShaderProgram struct {
	renderID uint32
	name     string
}

func NewShader(name, vecSrc, fragSrc string) *ShaderProgram {
	shader := new(ShaderProgram)
	shader.name = name
	shaderSource := make(map[uint32]string)
	shaderSource[glapi.VERTEX_SHADER] = vecSrc
	shaderSource[glapi.FRAGMENT_SHADER] = fragSrc

	shader.compile(shaderSource)

	return shader
}

func (sp *ShaderProgram) Bind()  {
	sRenderer.gapi.UseProgram(sp.renderID)
}

func (sp *ShaderProgram) compile(shaderSource map[uint32]string)  {
	program := sRenderer.gapi.CreateProgram()
	shaders := make([]uint32, 0, 2)
	
	for t, v := range shaderSource {
		shader := sRenderer.gapi.CreateShader(t)
		sRenderer.gapi.ShaderSource(shader, v)
		sRenderer.gapi.CompileShader(shader)
		isCompiled := sRenderer.gapi.GetShaderiv(shader, glapi.COMPILE_STATUS)
		if isCompiled == glapi.FALSE {
			println(sRenderer.gapi.GetShaderInfoLog(shader))
			sRenderer.gapi.DeleteShader(shader)
			break
		}
		sRenderer.gapi.AttachShader(program, shader)
		shaders = append(shaders, shader)
	}
	sp.renderID = program
	sRenderer.gapi.LinkProgram(program)
	isLinked := sRenderer.gapi.GetProgramiv(program, glapi.LINK_STATUS)
	if isLinked == glapi.FALSE {
		println(sRenderer.gapi.GetProgramInfoLog(program))
		sRenderer.gapi.DeleteShader(program)

		for _, shader := range shaders {
			sRenderer.gapi.DeleteShader(shader)
		}

		return
	}

	for _, shader := range shaders {
		sRenderer.gapi.DetachShader(program, shader)
		sRenderer.gapi.DeleteShader(shader)
	}
}
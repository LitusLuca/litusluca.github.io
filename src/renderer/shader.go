package renderer

import (
	"fmt"
	"io"
	"strings"

	"github.com/litusluca/litusluca.github.io/src/glapi"
	"github.com/litusluca/litusluca.github.io/src/utils/loader"
)

type ShaderProgram struct {
	renderID uint32
	name     string
}

func NewShaderFromFile(shaderSourcePath string) *ShaderProgram {
	source := readFile(shaderSourcePath)
	fmt.Println(source)
	shaderSource := preprocess(source)

	shader := new(ShaderProgram)
	shader.name = shaderSourcePath
	shader.compile(shaderSource)
	return shader
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

func shaderTypeToGL(sType string) uint32 {
	switch sType {
	case "vertex":
		return glapi.VERTEX_SHADER
	case "fragment":
		return glapi.FRAGMENT_SHADER
	case "geometry":
		return glapi.GEOMETRY_SHADER
	default:
		return 0
	}
}

func preprocess(source string) map[uint32]string {
	shaderSource := make(map[uint32]string)
	key := "#type"
	keyOffset := len(key)
	i := strings.Index(source, key)
	temp := source
	for ;i != -1;{
		eol := strings.Index(temp[i:], "\r\n") + i
		begin := i + keyOffset + 1
		sType := temp[begin : eol]
		temp = strings.TrimLeft(temp[eol:], "\r\n")
		i = strings.Index(temp, key)
		if i == -1 {
			shaderSource[shaderTypeToGL(sType)] = temp
		}else {
			shaderSource[shaderTypeToGL(sType)] = temp[:i]
		}
	}
	return shaderSource
}
func readFile(path string) string {
	file, err := loader.ReadFile("/shaders/" + path)
	if err != nil {
		fmt.Println(err)
	}
	buffer, _ := io.ReadAll(file)
	return string(buffer)
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
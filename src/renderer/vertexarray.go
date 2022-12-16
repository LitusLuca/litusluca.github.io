package renderer

import "fmt"

type VertexArray struct {
	renderID          uint32
	vertexBufferIndex uint32
	vertexBuffers     []*VertexBuffer
	indexBuffer       *IndexBuffer
}

func NewVAO() *VertexArray {
	vao := new(VertexArray)
	vao.renderID = sRenderer.gapi.CreateVertexArray()
	return vao
}

func (vao *VertexArray) Bind() {
	sRenderer.gapi.BindVertexArray(vao.renderID)
}

func (vao *VertexArray) AddVertexBuffer(vbo *VertexBuffer) {
	layout := vbo.GetLayout()
	if len(layout.GetElements()) == 0 {
		fmt.Println("! ERROR: Must specifie layout!")
		return
	}
	vao.Bind()
	vbo.Bind()
	for _, v := range layout.elements {
		switch v.Type {
		case TypeFloat, TypeFloat2, TypeFloat3, TypeFloat4:
			sRenderer.gapi.EnableVertexAttribArray(vao.vertexBufferIndex)
			sRenderer.gapi.VertexAttribPointer(vao.vertexBufferIndex, v.GetComponentCount(),
			 	shaderDataTypeToGLBase(v.Type), boolToGLBool(v.Normalize), layout.stride, v.Offset)
			vao.vertexBufferIndex++
		case TypeInt, TypeInt2,TypeInt3,TypeInt4, TypeBool:
			sRenderer.gapi.EnableVertexAttribArray(vao.vertexBufferIndex)
			sRenderer.gapi.VertexAttribIPointer(vao.vertexBufferIndex, v.GetComponentCount(),
				shaderDataTypeToGLBase(v.Type), layout.stride, v.Offset)
			vao.vertexBufferIndex++
		default:
			return
		}
	}
	vao.vertexBuffers = append(vao.vertexBuffers, vbo)
}

func (vao *VertexArray) AddIndexBuffer(ibo *IndexBuffer) {
	sRenderer.gapi.BindVertexArray(vao.renderID)
	ibo.Bind()
	vao.indexBuffer = ibo
}

func (vao *VertexArray) GetIndexBuffer() *IndexBuffer {
	return vao.indexBuffer
}
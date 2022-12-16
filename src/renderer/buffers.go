package renderer

import "github.com/litusluca/litusluca.github.io/src/glapi"

type VertexBuffer struct {
	renderID uint32
	layout *BufferLayout
}

func NewVBO(vertices []float32) *VertexBuffer{
	vbo := new(VertexBuffer)
	vbo.renderID = sRenderer.gapi.CreateBuffer()
	sRenderer.gapi.BindBuffer(glapi.ARRAY_BUFFER, vbo.renderID)
	sRenderer.gapi.BufferData(glapi.ARRAY_BUFFER, vertices, glapi.STATIC_DRAW)
	return vbo
}

func (vbo *VertexBuffer) Delete()  {
	sRenderer.gapi.DeleteBuffer(vbo.renderID)
}

func (vbo *VertexBuffer) Bind()  {
	sRenderer.gapi.BindBuffer(glapi.ARRAY_BUFFER, vbo.renderID)
}

func (vbo *VertexBuffer) SetLayout(layout *BufferLayout)  {
	vbo.layout=layout
}

func (vbo *VertexBuffer) GetLayout() *BufferLayout {
	return vbo.layout
}

type IndexBuffer struct {
	renderID uint32
	count uint32
}

func NewIBO(indexes []uint32) *IndexBuffer{
	ibo := new(IndexBuffer)	
	ibo.count = uint32(len(indexes))
	ibo.renderID = sRenderer.gapi.CreateBuffer()
	sRenderer.gapi.BindBuffer(glapi.ELEMENT_ARRAY_BUFFER, ibo.renderID)
	sRenderer.gapi.BufferData(glapi.ELEMENT_ARRAY_BUFFER, indexes, glapi.STATIC_DRAW)
	return ibo
}

func (ibo *IndexBuffer) Delete()  {
	sRenderer.gapi.DeleteBuffer(ibo.renderID)
}

func (ibo *IndexBuffer) Bind()  {
	sRenderer.gapi.BindBuffer(glapi.ELEMENT_ARRAY_BUFFER, ibo.renderID)
}

func (ibo *IndexBuffer) GetCount() uint32 {
	return ibo.count
}
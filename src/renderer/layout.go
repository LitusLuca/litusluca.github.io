package renderer

import "github.com/litusluca/litusluca.github.io/src/glapi"

type ShaderDataType uint32

const (
	TypeNone ShaderDataType = iota
	TypeFloat
	TypeFloat2
	TypeFloat3
	TypeFloat4
	TypeMat3
	TypeMat4
	TypeInt
	TypeInt2
	TypeInt3
	TypeInt4
	TypeBool
)

func boolToGLBool(trueFalse bool) uint32 {
	if trueFalse {
		return glapi.TRUE
	}else {
		return glapi.FALSE
	}
}

func shaderDataTypeSize(datatype ShaderDataType) uint32 {
	switch datatype {
	case TypeFloat:
		return 4
	case TypeFloat2:
		return 4 * 2
	case TypeFloat3:
		return 4 * 3
	case TypeFloat4:
		return 4 * 4
	case TypeMat3:
		return 4 * 3 * 3
	case TypeMat4:
		return 4 * 4 * 4
	case TypeInt:
		return 4
	case TypeInt2:
		return 4 * 2
	case TypeInt3:
		return 4 * 3
	case TypeInt4:
		return 4 * 4
	case TypeBool:
		return 1
	default:
		return 0
	}
}

func shaderDataTypeToGLBase(datatype ShaderDataType) uint32 {
	switch datatype {
	case TypeFloat, TypeFloat2, TypeFloat3, TypeFloat4, TypeMat3, TypeMat4:
		return glapi.FLOAT
	case TypeInt, TypeInt2, TypeInt3, TypeInt4:
		return glapi.INT
	case TypeBool:
		return glapi.BOOL
	default:
		return 0
	}
}

type BufferElement struct {
	Name string
	Type ShaderDataType
	Size uint32
	Offset uint32
	Normalize bool
}
func MakeBufferElement(datatype ShaderDataType, name string) BufferElement {
	return BufferElement{name, datatype, shaderDataTypeSize(datatype), 0, false}
}

func (element BufferElement) GetComponentCount() uint32 {
	switch element.Type {
	case TypeFloat, TypeInt, TypeBool:
		return 1
	case TypeFloat2, TypeInt2:
		return 2
	case TypeFloat3, TypeInt3, TypeMat3:
		return 3
	case TypeFloat4, TypeInt4, TypeMat4:
		return 4
	default:
		return 0
	}
}

type BufferLayout struct {
	elements []BufferElement
	stride uint32
}

func NewBufferLayout(elements []BufferElement) *BufferLayout {
	layout := new(BufferLayout)
	layout.elements = elements
	layout.calculateOffsetAndStride()
	return layout
}

func (layout *BufferLayout) calculateOffsetAndStride() {
	var offset uint32 = 0 
	layout.stride = 0

	for i,v := range layout.elements {
		layout.elements[i].Offset = offset
		offset += v.Size
		layout.stride += v.Size
	}
}

func (layout *BufferLayout) GetElements() []BufferElement {
	return layout.elements
}
func (layout *BufferLayout) GetStride() uint32 {
	return layout.stride
}
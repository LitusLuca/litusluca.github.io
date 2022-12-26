package renderer

import (
	"fmt"

	"github.com/litusluca/litusluca.github.io/src/glapi"
	"github.com/litusluca/litusluca.github.io/src/utils/loader"
)

type ITexture interface {
	GetWidth() uint32
	GetHeight() uint32

	//SetData(data []byte)
	Bind(slot uint32)
	IsLoaded() bool
}

type ColorFormat byte

const (
	GrayScale ColorFormat = 0b00
	TrueColor ColorFormat = 0b10
	GrayScaleAlpha ColorFormat = 0b01
	TrueColorAlpha ColorFormat = 0b11
)

type Texture2D struct {
	renderID                   uint32
	width, height              uint32
	formatInternal, formatData uint32
	isLoaded                   bool
}

func NewTexture2D(path string, colorFormat ColorFormat) *Texture2D {
	tex := new(Texture2D)
	fmt.Println(path)
	var width, height uint32
	path = "textures/" + path
	data, err := loader.LoadImage(path, &width, &height)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	tex.width, tex.height = width, height
	tex.isLoaded = true
	tex.formatData = glapi.RGBA
	
	switch colorFormat {
	case GrayScale:
		tex.formatInternal = glapi.R8
		tex.formatData = glapi.RED
		data = optimizeData(data, 1, 0b0001)
	case GrayScaleAlpha:
		tex.formatInternal = glapi.RG8
		tex.formatData = glapi.RG
		data = optimizeData(data, 2, 0b1001)
	case TrueColor:
		tex.formatInternal = glapi.RGB8
		tex.formatData = glapi.RGB
		data = optimizeData(data, 3, 0b0111)
	case TrueColorAlpha:
		tex.formatInternal = glapi.RGBA8
		tex.formatData = glapi.RED
		data = optimizeData(data, 4, 0b1111)
	default:
		fmt.Println("Error: unknown colorformat")
		return nil
	}

	fmt.Println(data)

	tex.renderID = sRenderer.gapi.CreateTexture()
	sRenderer.gapi.BindTexture(glapi.TEXTURE_2D, tex.renderID)
	sRenderer.gapi.TexStorage2D(glapi.TEXTURE_2D, 1, tex.formatInternal, tex.width, tex.height)
	sRenderer.gapi.TexSubImage2D(glapi.TEXTURE_2D, 0, 0, 0, tex.width, tex.height, tex.formatData, glapi.UNSIGNED_BYTE, data)
	sRenderer.gapi.GenerateMipmap(glapi.TEXTURE_2D)
	return tex
}

func optimizeData(data []byte,channels, channelBits byte) []byte {
	newCap := len(data)/4*int(channels)
	newData := make([]byte, 0,newCap)

	for i,v := range data {
		bit := byte(i) % 4
		if checkIfBitisSet(channelBits, bit) != 0 {
			newData = append(newData, v)
		}
	}
	return newData
}

func checkIfBitisSet(bits byte, nbit byte) byte {
	return bits & (1 << nbit)
}

func (tex *Texture2D) IsLoaded() bool {
	return tex.isLoaded
}

func (tex *Texture2D) GetWidth() uint32 {
	return tex.width
}

func (tex *Texture2D) GetHeight() uint32 {
	return tex.height
}

func (tex *Texture2D) Bind(slot uint32)  {
	sRenderer.gapi.ActiveTexture(glapi.TEXTURE0+slot)
	sRenderer.gapi.BindTexture(glapi.TEXTURE_2D, tex.renderID)
}
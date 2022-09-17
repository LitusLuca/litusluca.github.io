package layers

//TODO
import "time"

type ILayer interface {
	OnAttach()
	OnUpdate(dt time.Duration)
	OnDetach()
}

type LayerStack struct {
	layers []ILayer
	layerInsertIndex uint
}

func NewLayerStack() *LayerStack {
	ls := new(LayerStack)
	ls.layers = make([]ILayer, 0)
	ls.layerInsertIndex = 0
	return ls
}

func (ls *LayerStack) PushLayer(layer ILayer)  {
	
}

func (ls *LayerStack) PushOveray(layer ILayer)  {
	
}

func (ls *LayerStack) PopLayer(layer ILayer)  {
	
}

func (ls *LayerStack) PopOverlay(layer ILayer)  {
	
}

func (ls *LayerStack) GetLayerByIndex(index uint) ILayer{
	return ls.layers[index]
}
func (ls *LayerStack) GetLayerCount()  uint{
	return uint(len(ls.layers))
}
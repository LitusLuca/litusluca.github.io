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
	layerInsertIndex int
}

func NewLayerStack() *LayerStack {
	ls := new(LayerStack)
	ls.layers = make([]ILayer, 0)
	ls.layerInsertIndex = 0
	return ls
}

func (ls *LayerStack) PushLayer(layer ILayer)  {
	layer.OnAttach()
	temp := append([]ILayer{layer}, ls.layers[ls.layerInsertIndex:]...)
	ls.layers = append(ls.layers[:ls.layerInsertIndex], temp...)
	ls.layerInsertIndex++
}

func (ls *LayerStack) PushOveray(overlay ILayer)  {
	overlay.OnAttach()
	ls.layers = append(ls.layers, overlay)
}

func (ls *LayerStack) PopLayer(layer ILayer)  {
	for p,v := range ls.layers[:ls.layerInsertIndex] {
		if v == layer {
			layer.OnDetach()
			ls.layers = append(ls.layers[:p], ls.layers[p+1:]...)
			ls.layerInsertIndex--
		}
	}
}

func (ls *LayerStack) PopOverlay(overlay ILayer)  {
	for p,v := range ls.layers[ls.layerInsertIndex:]{
		if v== overlay{
			overlay.OnDetach()
			ls.layers = append(ls.layers[:ls.layerInsertIndex+p], ls.layers[ls.layerInsertIndex+p+1:]...)
		}
	}
}

func (ls *LayerStack) PopAll()  {
	for _, layer :=range ls.layers{
		layer.OnDetach()
	}
	ls.layers = ls.layers[:0]
	ls.layerInsertIndex = 0
}

func (ls *LayerStack) GetLayerByIndex(index int) ILayer{
	return ls.layers[index]
}
func (ls *LayerStack) GetLayerCount()  int{
	return len(ls.layers)
}
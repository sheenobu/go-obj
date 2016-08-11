package drag

//go:generate rxgen -name RxDraggable -type Draggable drag.go

// Draggable defines an object that can be dragged
type Draggable interface {
	Position() (int32, int32)
	Move(int32, int32)
}

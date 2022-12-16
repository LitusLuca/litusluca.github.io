//go:build wasm
// +build wasm

package input

var keyStateMap = map[KeyCode]bool{
	KEY_UNKNOWN:       false,
	KEY_SPACE:         false,
	KEY_APOSTROPHE:    false,
	KEY_COMMA:         false,
	KEY_MINUS:         false,
	KEY_PERIOD:        false,
	KEY_SLASH:         false,
	KEY_0:             false,
	KEY_1:             false,
	KEY_2:             false,
	KEY_3:             false,
	KEY_4:             false,
	KEY_5:             false,
	KEY_6:             false,
	KEY_7:             false,
	KEY_8:             false,
	KEY_9:             false,
	KEY_SEMICOLON:     false,
	KEY_EQUAL:         false,
	KEY_A:             false,
	KEY_B:             false,
	KEY_C:             false,
	KEY_D:             false,
	KEY_E:             false,
	KEY_F:             false,
	KEY_G:             false,
	KEY_H:             false,
	KEY_I:             false,
	KEY_J:             false,
	KEY_K:             false,
	KEY_L:             false,
	KEY_M:             false,
	KEY_N:             false,
	KEY_O:             false,
	KEY_P:             false,
	KEY_Q:             false,
	KEY_R:             false,
	KEY_S:             false,
	KEY_T:             false,
	KEY_U:             false,
	KEY_V:             false,
	KEY_W:             false,
	KEY_X:             false,
	KEY_Y:             false,
	KEY_Z:             false,
	KEY_LEFT_BRACKET:  false,
	KEY_BACKSLASH:     false,
	KEY_RIGHT_BRACKET: false,
	KEY_GRAVE_ACCENT:  false,
	KEY_WORLD_1:       false,
	KEY_WORLD_2:       false,
	KEY_ESCAPE:        false,
	KEY_ENTER:         false,
	KEY_TAB:           false,
	KEY_BACKSPACE:     false,
	KEY_INSERT:        false,
	KEY_DELETE:        false,
	KEY_RIGHT:         false,
	KEY_LEFT:          false,
	KEY_DOWN:          false,
	KEY_UP:            false,
	KEY_PAGE_UP:       false,
	KEY_PAGE_DOWN:     false,
	KEY_HOME:          false,
	KEY_END:           false,
	KEY_CAPS_LOCK:     false,
	KEY_SCROLL_LOCK:   false,
	KEY_NUM_LOCK:      false,
	KEY_PRINT_SCREEN:  false,
	KEY_PAUSE:         false,
	KEY_F1:            false,
	KEY_F2:            false,
	KEY_F3:            false,
	KEY_F4:            false,
	KEY_F5:            false,
	KEY_F6:            false,
	KEY_F7:            false,
	KEY_F8:            false,
	KEY_F9:            false,
	KEY_F10:           false,
	KEY_F11:           false,
	KEY_F12:           false,
	KEY_F13:           false,
	KEY_F14:           false,
	KEY_F15:           false,
	KEY_F16:           false,
	KEY_F17:           false,
	KEY_F18:           false,
	KEY_F19:           false,
	KEY_F20:           false,
	KEY_F21:           false,
	KEY_F22:           false,
	KEY_F23:           false,
	KEY_F24:           false,
	KEY_F25:           false,
	KEY_KP_0:          false,
	KEY_KP_1:          false,
	KEY_KP_2:          false,
	KEY_KP_3:          false,
	KEY_KP_4:          false,
	KEY_KP_5:          false,
	KEY_KP_6:          false,
	KEY_KP_7:          false,
	KEY_KP_8:          false,
	KEY_KP_9:          false,
	KEY_KP_DECIMAL:    false,
	KEY_KP_DIVIDE:     false,
	KEY_KP_MULTIPLY:   false,
	KEY_KP_SUBTRACT:   false,
	KEY_KP_ADD:        false,
	KEY_KP_ENTER:      false,
	KEY_KP_EQUAL:      false,
	KEY_LEFT_SHIFT:    false,
	KEY_LEFT_CONTROL:  false,
	KEY_LEFT_ALT:      false,
	KEY_LEFT_SUPER:    false,
	KEY_RIGHT_SHIFT:   false,
	KEY_RIGHT_CONTROL: false,
	KEY_RIGHT_ALT:     false,
	KEY_RIGHT_SUPER:   false,
	KEY_MENU:          false,
}

func IsKeyPressed(key KeyCode) bool {
	return keyStateMap[key]
}

func OnKeyEvent(key KeyCode, state bool){
	keyStateMap[key] = state
}

var mouseButtonStateMap = map[MouseButton]bool{
	MOUSE_BUTTON_1: false,
	MOUSE_BUTTON_2: false,
	MOUSE_BUTTON_3: false,
	MOUSE_BUTTON_4: false,
	MOUSE_BUTTON_5: false,
	MOUSE_BUTTON_6: false,
	MOUSE_BUTTON_7: false,
	MOUSE_BUTTON_8: false,
}

func IsMouseButtonPressed(button MouseButton) bool {
	return mouseButtonStateMap[button]
}

func OnMouseButtonEvent(button MouseButton, state bool)  {
	mouseButtonStateMap[button] = state
}

var mouseX,mouseY int32

func GetMousePos() (float32,float32) {
	return float32(mouseX), float32(mouseY)
}

func OnMouseMove(x, y int32)  {
	mouseX, mouseY = x, y
}
package utils

import (
	"Elevator/driver-go-master/elevio"
)

func Init(addr string, numFloors int) {
	elevio.Init(addr, numFloors)
}

func WrapRequestButton(floor int, button elevio.ButtonType) bool {
	return elevio.GetButton(button, floor)
}

func WrapRequestButtonLight(floor int, button elevio.ButtonType, value bool) {
	elevio.SetButtonLamp(button, floor, value)
}

func WrapMotorDirection(direction elevio.MotorDirection) {
	elevio.SetMotorDirection(direction)
}

// GetInputDevice returns the elevator's input device.
func GetInputDevice() ElevInputDevice {
	return ElevInputDevice{
		FloorSensor:   elevio.PollFloorSensor,
		RequestButton: WrapRequestButton,
		StopButton:    elevio.PollStopButton,
		Obstruction:   elevio.PollObstructionSwitch,
	}
}

// GetOutputDevice returns the elevator's output device.
func GetOutputDevice() ElevOutputDevice {
	return ElevOutputDevice{
		FloorIndicator:     elevio.SetFloorIndicator,
		RequestButtonLight: WrapRequestButtonLight,
		DoorLight:          elevio.SetDoorOpenLamp,
		StopButtonLight:    elevio.SetStopLamp,
		MotorDirection:     WrapMotorDirection,
	}
}

// DirnToString converts Direction to a string.
func DirnToString(d Dirn) string {
	switch d {
	case D_Up:
		return "D_Up"
	case D_Down:
		return "D_Down"
	case D_Stop:
		return "D_Stop"
	default:
		return "D_UNDEFINED"
	}
}

// ButtonToString converts Button to a string.
func ButtonToString(b Button) string {
	switch b {
	case B_HallUp:
		return "B_HallUp"
	case B_HallDown:
		return "B_HallDown"
	case B_Cab:
		return "B_Cab"
	default:
		return "B_UNDEFINED"
	}
}
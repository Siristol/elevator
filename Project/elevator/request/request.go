package request

import (
	"Elevator/driver_go_master/elevio"
	"Elevator/elevator/initial"
)


type DirnBehaviourPair struct {
	Dirn      elevio.MotorDirection
	Behaviour initial.ElevatorBehaviour
}

func RequestsAbove(e initial.Elevator) bool {
	for f := e.Floor + 1; f < initial.N_FLOORS; f++ {
		for btn := 0; btn < initial.N_BUTTONS; btn++ {
			if e.Requests[f][btn] {
				return true
			}
		}
	}
	return false
}

func RequestsBelow(e initial.Elevator) bool {
	for f := 0; f < e.Floor; f++ {
		for btn := 0; btn < initial.N_BUTTONS; btn++ {
			if e.Requests[f][btn] {
				return true
			}
		}
	}
	return false
}

func RequestsHere(e initial.Elevator) bool {
	for btn := 0; btn < initial.N_BUTTONS; btn++ {
		if e.Requests[e.Floor][btn] {
			return true
		}
	}
	return false
}

func RequestsChooseDirection(e initial.Elevator) DirnBehaviourPair {
	switch e.Dirn {
	case elevio.MDUp:
		if RequestsAbove(e) {
			return DirnBehaviourPair{elevio.MDUp, initial.EBMoving}
		} else if RequestsHere(e) {
			return DirnBehaviourPair{elevio.MDDown, initial.EBDoorOpen}
		} else if RequestsBelow(e) {
			return DirnBehaviourPair{elevio.MDDown, initial.EBMoving}
		} else {
			return DirnBehaviourPair{elevio.MDStop, initial.EBIdle}
		}
	case elevio.MDDown:
		if RequestsBelow(e) {
			return DirnBehaviourPair{elevio.MDDown, initial.EBMoving}
		} else if RequestsHere(e) {
			return DirnBehaviourPair{elevio.MDUp, initial.EBDoorOpen}
		} else if RequestsAbove(e) {
			return DirnBehaviourPair{elevio.MDUp, initial.EBMoving}
		} else {
			return DirnBehaviourPair{elevio.MDStop, initial.EBIdle}
		}
	case elevio.MDStop:
		if RequestsHere(e) {
			return DirnBehaviourPair{elevio.MDStop, initial.EBDoorOpen}
		} else if RequestsAbove(e) {
			return DirnBehaviourPair{elevio.MDUp, initial.EBMoving}
		} else if RequestsBelow(e) {
			return DirnBehaviourPair{elevio.MDDown, initial.EBMoving}
		} else {
			return DirnBehaviourPair{elevio.MDStop, initial.EBIdle}
		}
	default:
		return DirnBehaviourPair{elevio.MDStop, initial.EBIdle}
	}
}

func RequestsShouldStop(e initial.Elevator) bool {
	switch e.Dirn {
	case elevio.MDDown:
		return e.Requests[e.Floor][elevio.BTHallDown] ||
			e.Requests[e.Floor][elevio.BTCab] ||
			!RequestsBelow(e)
	case elevio.MDUp:
		return e.Requests[e.Floor][elevio.BTHallUp] ||
			e.Requests[e.Floor][elevio.BTCab] ||
			!RequestsAbove(e)
	case elevio.MDStop:
	default:
		return true
	}
	return true
}

func RequestsShouldClearImmediately(e initial.Elevator, btnFloor int, btnType elevio.ButtonType) bool {
	switch e.ClearRequestVariant {
	case initial.CV_All:
		return e.Floor == btnFloor
	case initial.CV_InDirn:
		return e.Floor == btnFloor &&
			((e.Dirn == elevio.MDUp && btnType == elevio.BTHallUp) ||
				(e.Dirn == elevio.MDDown && btnType == elevio.BTHallDown) ||
				e.Dirn == elevio.MDStop ||
				btnType == elevio.BTCab)
	default:
		return false
	}
}

func RequestsClearAtCurrentFloor(e initial.Elevator) initial.Elevator {
	switch e.ClearRequestVariant {
	case initial.CV_All:
		for btn := 0; btn < initial.N_BUTTONS; btn++ {
			e.Requests[e.Floor][btn] = false
		}
	case initial.CV_InDirn:
		e.Requests[e.Floor][elevio.BTCab] = false
		switch e.Dirn {
		case elevio.MDUp:
			if !RequestsAbove(e) && !e.Requests[e.Floor][elevio.BTHallUp] {
				e.Requests[e.Floor][elevio.BTHallDown] = false
			}
			e.Requests[e.Floor][elevio.BTHallUp] = false
		case elevio.MDDown:
			if !RequestsBelow(e) && !e.Requests[e.Floor][elevio.BTHallDown] {
				e.Requests[e.Floor][elevio.BTHallUp] = false
			}
			e.Requests[e.Floor][elevio.BTHallDown] = false
		case elevio.MDStop:
		default:
			e.Requests[e.Floor][elevio.BTHallUp] = false
			e.Requests[e.Floor][elevio.BTHallDown] = false
		}
	default:
	}
	return e
}

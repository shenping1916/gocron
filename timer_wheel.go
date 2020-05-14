package gocron

type timeWheel struct {
	nextTw *timeWheel
	prevTw *timeWheel
}

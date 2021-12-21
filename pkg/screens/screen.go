package screens

type Screen interface {
	Render()
	OnDial(int) Screen
}

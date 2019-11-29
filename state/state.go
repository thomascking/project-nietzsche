package state

//State is type enum for tracking state of game
type State int

//State enum
const (
	GS State = iota
	MS
	PS
)

//CurrState is the current state
var CurrState = MS

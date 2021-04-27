package d1typ

const (
	GameServerStateOffline GameServerState = iota
	GameServerStateOnline
	GameServerStateStarting
)

type GameServerState int

var GameServerStates = map[GameServerState]string{
	GameServerStateOffline:  "Offline",
	GameServerStateOnline:   "Online",
	GameServerStateStarting: "Starting",
}

func (g GameServerState) Validate() error {
	_, ok := GameServerStates[g]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (g GameServerState) String() string {
	name, ok := GameServerStates[g]
	if ok {
		return name
	}

	return unknownStr
}

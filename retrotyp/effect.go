package retrotyp

type Effect struct {
	Id        int
	ZoneShape EffectZoneShape
	ZoneSize  int
	DiceNum   int
	DiceSide  int
	Value     int
	Param     string
	Random    int
	Duration  int
	// Delay      int
	TargetId int
	// TargetMask string
	// Group      int
	// Triggers   string
	Hidden bool
}

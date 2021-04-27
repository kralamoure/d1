package d1typ

type SpellLevel struct {
	Grade                      int
	Effects                    []Effect
	EffectsCritical            []Effect
	APCost                     int
	Range                      int
	RangeMax                   int
	CriticalHitProbability     int
	CriticalFailureProbability int
	Linear                     bool
	RequiresLineOfSight        bool
	RequiresFreeCell           bool
	AdjustableRange            bool
	ClassId                    ClassId
	MaxCastsPerTurn            int
	MaxCastsPerTarget          int
	MinCastInterval            int
	StatesRequired             []int
	StatesForbidden            []int
	MinPlayerLevel             int
	CriticalFailureEndsTurn    bool
}

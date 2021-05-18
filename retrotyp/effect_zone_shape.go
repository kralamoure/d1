package retrotyp

const (
	EffectZoneShapePoint  EffectZoneShape = "P"
	EffectZoneShapeCircle EffectZoneShape = "C"
	EffectZoneShapeCross  EffectZoneShape = "X"
	EffectZoneShapeLine   EffectZoneShape = "L"
	EffectZoneShapeRing   EffectZoneShape = "O"
	EffectZoneShapeT      EffectZoneShape = "T"
)

type EffectZoneShape string

var EffectZoneShapes = map[EffectZoneShape]string{
	EffectZoneShapePoint:  "Point",
	EffectZoneShapeCircle: "Circle",
	EffectZoneShapeCross:  "Cross",
	EffectZoneShapeLine:   "Line",
	EffectZoneShapeRing:   "Ring",
	EffectZoneShapeT:      "T",
}

func (z EffectZoneShape) Validate() error {
	_, ok := EffectZoneShapes[z]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (z EffectZoneShape) String() string {
	name, ok := EffectZoneShapes[z]
	if ok {
		return name
	}

	return unknownStr
}

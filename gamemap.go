package d1

import "github.com/kralamoure/d1util"

type GameMap struct {
	Id            int
	Name          string
	Width         int
	Height        int
	Background    int
	Ambiance      int
	Music         int
	Outdoor       bool
	Capabilities  int
	Data          string
	EncryptedData string
	Key           string
}

type GameMapPath struct {
	CellId int
	Dir    int
}

func (m GameMap) Cells() ([]d1util.Cell, error) {
	utilCells, err := d1util.DecompressCells(m.Data, false)
	if err != nil {
		return nil, err
	}

	builtCells := d1util.BuiltCells(nil, false, m.Width, utilCells)

	cells := make([]d1util.Cell, len(builtCells))
	for i, v := range builtCells {
		cells[i] = d1util.Cell{
			Id:                             v.Id,
			Active:                         v.Active,
			LineOfSight:                    v.LineOfSight,
			LayerGroundRot:                 v.LayerGroundRot,
			GroundLevel:                    v.GroundLevel,
			Movement:                       v.Movement,
			LayerGroundNum:                 v.LayerGroundNum,
			GroundSlope:                    v.GroundSlope,
			LayerGroundFlip:                v.LayerGroundFlip,
			LayerObject1Num:                v.LayerObject1Num,
			LayerObject1Rot:                v.LayerObject1Rot,
			LayerObject1Flip:               v.LayerObject1Flip,
			LayerObject2Flip:               v.LayerObject2Flip,
			LayerObject2Interactive:        v.LayerObject2Interactive,
			LayerObject2Num:                v.LayerObject2Num,
			PermanentLevel:                 v.PermanentLevel,
			LayerObjectExternal:            v.LayerObjectExternal,
			LayerObjectExternalInteractive: v.LayerObjectExternalInteractive,
			X:                              v.X,
			Y:                              v.Y,
			SpriteOnId:                     v.SpriteOnId,
		}
	}

	return cells, nil
}

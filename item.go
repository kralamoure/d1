package retro

import (
	"time"

	"github.com/kralamoure/retro/retrotyp"
)

type Item struct {
	Id         int
	TemplateId int
	Quantity   int
	Effects    []retrotyp.Effect
}

func (item Item) DisplayEffects() []retrotyp.Effect {
	effects := make([]retrotyp.Effect, len(item.Effects))
	copy(effects, item.Effects)

	for _, v := range effects {
		if v.Id == 995 {
			validity := time.Unix(int64(v.DiceSide/1000), 0)

			if !validity.Before(time.Now()) {
				dur := validity.Sub(time.Now())

				days := int(dur.Hours()) / 24
				hours := int(dur.Hours()) % 24
				minutes := int(dur.Minutes()) % 60

				effects = append(effects, retrotyp.Effect{
					Id:       998,
					DiceNum:  days,
					DiceSide: hours,
					Value:    minutes,
				})
			} else {
				effects = append(effects, retrotyp.Effect{
					Id: 994,
				})
			}

			break
		}
	}

	return effects
}

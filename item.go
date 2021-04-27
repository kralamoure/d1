package d1

import (
	"time"

	"github.com/kralamoure/d1/d1typ"
)

type Item struct {
	Id         int
	TemplateId int
	Quantity   int
	Effects    []d1typ.Effect
}

func (item Item) DisplayEffects() []d1typ.Effect {
	effects := make([]d1typ.Effect, len(item.Effects))
	copy(effects, item.Effects)

	for _, v := range effects {
		if v.Id == 995 {
			validity := time.Unix(int64(v.DiceSide/1000), 0)

			if !validity.Before(time.Now()) {
				dur := validity.Sub(time.Now())

				days := int(dur.Hours()) / 24
				hours := int(dur.Hours()) % 24
				minutes := int(dur.Minutes()) % 60

				effects = append(effects, d1typ.Effect{
					Id:       998,
					DiceNum:  days,
					DiceSide: hours,
					Value:    minutes,
				})
			} else {
				effects = append(effects, d1typ.Effect{
					Id: 994,
				})
			}

			break
		}
	}

	return effects
}

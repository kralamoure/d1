// Package d1 is a core library for Dofus 1.
package d1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1/d1typ"
)

func DecodeItemEffects(sli []string) (effects []d1typ.Effect, err error) {
	effects = make([]d1typ.Effect, len(sli))
	for i, v := range sli {
		effect, err2 := DecodeItemEffect(v)
		if err2 != nil {
			err = err2
			return
		}
		effects[i] = effect
	}

	return
}

func DecodeItemEffect(s string) (effect d1typ.Effect, err error) {
	if s == "" {
		err = errors.New("string is empty")
		return
	}

	effect = d1typ.Effect{
		ZoneShape: d1typ.EffectZoneShapeCircle,
		Hidden:    true,
	}

	sli := strings.Split(s, "#")
	for i, v := range sli {
		switch i {
		case 0:
			n, err2 := strconv.ParseInt(v, 16, 64)
			if err2 != nil {
				err = err2
				return
			}
			effect.Id = int(n)
		case 1:
			n, err2 := strconv.ParseInt(v, 16, 64)
			if err2 != nil {
				err = err2
				return
			}
			effect.DiceNum = int(n)
		case 2:
			n, err2 := strconv.ParseInt(v, 16, 64)
			if err2 != nil {
				err = err2
				return
			}
			effect.DiceSide = int(n)
		case 3:
			n, err2 := strconv.ParseInt(v, 16, 64)
			if err2 != nil {
				err = err2
				return
			}
			effect.Value = int(n)
		case 4:
			effect.Param = v
		}
	}

	return
}

func EncodeItemEffects(effects []d1typ.Effect) []string {
	sli := make([]string, len(effects))
	for i, v := range effects {
		sli[i] = EncodeItemEffect(v)
	}
	return sli
}

func EncodeItemEffect(effect d1typ.Effect) string {
	param := ""
	if effect.Param != "" {
		param = fmt.Sprintf("#%s", effect.Param)
	}
	return fmt.Sprintf("%x#%x#%x#%x%s", effect.Id, effect.DiceNum, effect.DiceSide, effect.Value, param)
}

func EffectDiceParam(effect d1typ.Effect) string {
	if effect.DiceNum == 0 {
		return ""
	}

	diceNum := effect.DiceNum
	diceSide := effect.DiceSide

	if diceSide < diceNum {
		return fmt.Sprintf("0d0+%d", diceNum)
	}

	return fmt.Sprintf("1d%d+%d", diceSide-diceNum+1, diceNum-1)
}

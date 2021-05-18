package retro

import (
	"time"

	"github.com/kralamoure/retro/retrotyp"
)

var MountXPFloors = []int{600, 1_750, 2_750, 4_000, 5_500, 7_250, 9_250, 11_500, 14_000, 16_750, 19_750, 23_000, 26_500,
	30_250, 34_250, 38_500, 43_000, 47_750, 52_750, 58_000, 63_500, 69_250, 75_250, 81_500, 88_000, 94_750, 101_750,
	109_000, 116_500, 124_250, 132_250, 140_500, 149_000, 157_750, 166_750, 176_000, 185_500, 195_250, 205_250, 215_500,
	226_000, 236_750, 247_750, 259_000, 270_500, 282_250, 294_250, 306_500, 319_000, 331_750, 344_750, 358_000, 371_500,
	385_250, 399_250, 413_500, 427_000, 441_750, 456_750, 473_000, 488_500, 504_250, 520_250, 536_500, 553_000, 569_750,
	586_750, 604_000, 621_500, 639_250, 657_250, 675_500, 694_000, 712_750, 731_750, 751_000, 770_000, 790_250, 810_250,
	830_500, 851_000, 871_750, 892_750, 914_000, 935_500, 957_250, 979_250, 1_001_500, 1_025_000, 1_046_750, 1_069_750,
	1_093_000, 1_116_500, 1_140_250, 1_164_250, 1_188_500, 1_213_000, 1_237_750, 1_262_750,
}

type Mount struct {
	Id          int
	TemplateId  int
	CharacterId int
	Name        string
	Sex         retrotyp.Sex
	XP          int
	Capacities  []retrotyp.MountCapacityId
	Validity    time.Time
}

func (c Mount) Level() int {
	level := 1
	for _, v := range MountXPFloors {
		if c.XP >= v {
			level++
		}
	}
	return level
}

func (c Mount) XPLow() int {
	if c.Level() < 2 {
		return 0
	}
	return MountXPFloors[c.Level()-2]
}

func (c Mount) XPHigh() int {
	if c.Level() > len(MountXPFloors) {
		// return math.MaxInt64
		return -1
	}
	return MountXPFloors[c.Level()-1]
}

package retro

import (
	"math"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro/retrotyp"
)

var CharacterXPFloors = []int{110, 650, 1_500, 2_800, 4_800, 7_300, 10_500, 14_500, 19_200, 25_200, 32_600, 41_000,
	50_500, 61_000, 75_000, 91_000, 115_000, 142_000, 171_000, 202_000, 235_000, 270_000, 310_000, 353_000, 398_500,
	448_000, 503_000, 561_000, 621_600, 687_000, 755_000, 829_000, 910_000, 1_000_000, 1_100_000, 1_240_000,
	1_400_000, 1_580_000, 1_780_000, 2_000_000, 2_250_000, 2_530_000, 2_850_000, 3_200_000, 3_570_000, 3_960_000,
	4_400_000, 4_860_000, 5_350_000, 5_860_000, 6_390_000, 6_950_000, 7_530_000, 8_130_000, 8_765_100, 9_420_000,
	10_150_000, 10_894_000, 11_655_000, 12_450_000, 13_278_000, 14_138_000, 15_171_000, 16_251_000, 17_377_000,
	18_553_000, 19_778_000, 21_055_000, 22_385_000, 23_769_000, 25_209_000, 26_707_000, 28_264_000, 29_882_000,
	31_563_000, 33_307_000, 35_118_000, 36_997_000, 38_945_000, 40_965_000, 43_059_000, 45_229_000, 47_476_000,
	49_803_000, 52_211_000, 54_704_000, 57_284_000, 59_952_000, 62_712_000, 65_565_000, 68_514_000, 71_561_000,
	74_710_000, 77_963_000, 81_323_000, 84_792_000, 88_374_000, 92_071_000, 95_886_000, 99_823_000, 103_885_000,
	108_075_000, 112_396_000, 116_853_000, 121_447_000, 126_184_000, 131_066_000, 136_098_000, 141_283_000,
	146_626_000, 152_130_000, 157_800_000, 163_640_000, 169_655_000, 175_848_000, 182_225_000, 188_791_000,
	195_550_000, 202_507_000, 209_667_000, 217_037_000, 224_620_000, 232_424_000, 240_452_000, 248_712_000,
	257_209_000, 265_949_000, 274_939_000, 284_186_000, 293_694_000, 303_473_000, 313_527_000, 323_866_000,
	334_495_000, 345_423_000, 356_657_000, 368_206_000, 380_076_000, 392_278_000, 404_818_000, 417_706_000,
	430_952_000, 444_564_000, 458_551_000, 472_924_000, 487_693_000, 502_867_000, 518_458_000, 534_476_000,
	550_933_000, 567_839_000, 585_206_000, 603_047_000, 621_374_000, 640_199_000, 659_536_000, 679_398_000,
	699_798_000, 720_751_000, 742_272_000, 764_374_000, 787_074_000, 810_387_000, 834_329_000, 858_917_000,
	884_167_000, 910_098_000, 936_727_000, 964_073_000, 992_154_000, 1_020_991_000, 1_050_603_000, 1_081_010_000,
	1_112_235_000, 1_144_298_000, 1_177_222_000, 1_211_030_000, 1_245_745_000, 1_281_393_000, 1_317_997_000,
	1_355_584_000, 1_404_179_000, 1_463_811_000, 1_534_506_000, 1_616_294_000, 1_709_205_000, 1_813_267_000,
	1_928_513_000, 2_054_975_000, 2_192_686_000, 2_341_679_000, 2_501_990_000, 2_673_655_000, 2_856_710_000,
	3_051_194_000, 3_257_146_000, 3_474_606_000, 3_703_616_000, 7_407_232_000,
}

var CharacterHonorFloors = []int{500, 1500, 3000, 5000, 7500, 10000, 12500, 15000, 17500, 18000}

type Character struct {
	Id               int
	AccountId        string
	GameServerId     int
	Name             retrotyp.CharacterName
	Sex              retrotyp.Sex
	ClassId          retrotyp.ClassId
	Color1           retrotyp.Color
	Color2           retrotyp.Color
	Color3           retrotyp.Color
	Alignment        retrotyp.Alignment
	AlignmentEnabled bool
	XP               int
	Kamas            int
	BonusPoints      int
	BonusPointsSpell int
	Honor            int
	Disgrace         int
	Stats            CharacterStats // TODO: validation (non-negative)
	GameMapId        int
	Cell             int
	Direction        int // TODO: maybe a custom type
	Spells           []CharacterSpell
	MountId          int
	Mounting         bool
}

type CharacterStats struct {
	Vitality     int
	Wisdom       int
	Strength     int
	Intelligence int
	Chance       int
	Agility      int
}

func (c Character) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required,
		),
		validation.Field(&c.Sex),
		validation.Field(&c.ClassId,
			validation.Required,
		),
		validation.Field(&c.Color1),
		validation.Field(&c.Color2),
		validation.Field(&c.Color3),
		validation.Field(&c.Alignment),
		validation.Field(&c.XP, validation.Min(0)),
		validation.Field(&c.Kamas, validation.Min(0)),
		validation.Field(&c.BonusPoints, validation.Min(0)),
		validation.Field(&c.BonusPointsSpell, validation.Min(0)),
		validation.Field(&c.Honor, validation.Min(0)),
		validation.Field(&c.Disgrace, validation.Min(0)),
	)
}

func (c Character) Level() int {
	level := 1
	for _, v := range CharacterXPFloors {
		if c.XP >= v {
			level++
		}
	}
	return level
}

func (c Character) Grade() retrotyp.AlignmentGrade {
	if c.Alignment == retrotyp.AlignmentNeutral {
		return retrotyp.AlignmentGradeNeutral
	}
	grade := 1
	for _, v := range CharacterHonorFloors {
		if c.Honor >= v {
			grade++
		}
	}
	return retrotyp.AlignmentGrade(grade)
}

func (c Character) XPLow() int {
	if c.Level() < 2 {
		return 0
	}
	return CharacterXPFloors[c.Level()-2]
}

func (c Character) XPHigh() int {
	if c.Level() > len(CharacterXPFloors) {
		return math.MaxInt64
	}
	return CharacterXPFloors[c.Level()-1]
}

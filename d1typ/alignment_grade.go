package d1typ

const (
	AlignmentGradeNeutral AlignmentGrade = iota
	AlignmentGradeRecruit
	AlignmentGradeCandidate
	AlignmentGradeSentry
	AlignmentGradeDefender
	AlignmentGradeKnight
	AlignmentGradeChampion
	AlignmentGradeConqueror
	AlignmentGradeStrategist
	AlignmentGradeCommander
	AlignmentGradeHero
)

type AlignmentGrade int

var AlignmentGrades = map[AlignmentGrade]string{
	AlignmentGradeNeutral:    "Neutral",
	AlignmentGradeRecruit:    "Recruit",
	AlignmentGradeCandidate:  "Candidate",
	AlignmentGradeSentry:     "Sentry",
	AlignmentGradeDefender:   "Defender",
	AlignmentGradeKnight:     "Knight",
	AlignmentGradeChampion:   "Champion",
	AlignmentGradeConqueror:  "Conqueror",
	AlignmentGradeStrategist: "Strategist",
	AlignmentGradeCommander:  "Commander",
	AlignmentGradeHero:       "Hero",
}

func (g AlignmentGrade) Validate() error {
	_, ok := AlignmentGrades[g]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (g AlignmentGrade) String() string {
	name, ok := AlignmentGrades[g]
	if ok {
		return name
	}

	return unknownStr
}

package character

type PC struct {
	FirstName string
	LastName  string

	HP    int
	HPMax int

	SpellSlots    int
	SpellSlotsMax int

	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int

	Proficiency int
	Saves       []string

	Skills    []string
	Expertise []string

	Misty int
}

func (p *PC) Mod(skill string) int {
	switch skill {
	case "Strength":
		return p.Strength
	case "Dexterity":
		return p.Dexterity
	case "Constitution":
		return p.Constitution
	case "Intelligence":
		return p.Intelligence
	case "Wisdom":
		return p.Wisdom
	case "Charisma":
		return p.Charisma
	}
	panic("invalid skill")
}

func (p *PC) SaveProficient(skill string) int {
	for _, s := range p.Saves {
		if s == skill {
			return p.Proficiency
		}
	}
	return 0
}

func (p *PC) SkillProficient(skill string) int {
	for _, s := range p.Skills {
		if s == skill {
			return p.Proficiency
		}
	}
	return 0
}

func (p *PC) ExpertiseMod(skill string) int {
	for _, s := range p.Expertise {
		if s == skill {
			return p.Proficiency
		}
	}
	return 0
}

const (
	Strength     = "Strength"
	Dexterity    = "Dexterity"
	Constitution = "Constitution"
	Intelligence = "Intelligence"
	Wisdom       = "Wisdom"
	Charisma     = "Charisma"
)

var Skills = map[string]string{
	"Acrobatics":      Dexterity,
	"Animal Handling": Wisdom,
	"Arcana":          Intelligence,
	"Athletics":       Strength,
	"Deception":       Charisma,
	"History":         Intelligence,
	"Insight":         Wisdom,
	"Intimidation":    Charisma,
	"Investigation":   Intelligence,
	"Medicine":        Wisdom,
	"Nature":          Intelligence,
	"Perception":      Wisdom,
	"Performance":     Charisma,
	"Persuasion":      Charisma,
	"Religion":        Intelligence,
	"Sleight of Hand": Dexterity,
	"Stealth":         Dexterity,
	"Survival":        Wisdom,
}

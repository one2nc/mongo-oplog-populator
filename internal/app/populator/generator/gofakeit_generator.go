package generator

type GoFakeItGenerator struct {
}

func (g *GoFakeItGenerator) NewGoFakeItGenerator() CustomGenerator {
	return &GoFakeItGenerator{}
}

// generateAge implements CustomGenerator
func (*GoFakeItGenerator) generateAge() int {
	return 0
}

// generateFirstName implements CustomGenerator
func (*GoFakeItGenerator) generateFirstName() string {
	return ""

}

// generateLastName implements CustomGenerator
func (*GoFakeItGenerator) generateLastName() string {
	return ""
}

// generatePhone implements CustomGenerator
func (*GoFakeItGenerator) generatePhone() string {
	return ""
}

// generatePositions implements CustomGenerator
func (*GoFakeItGenerator) generatePositions() []string {
	return nil
}

// generateSalary implements CustomGenerator
func (*GoFakeItGenerator) generateSalary() float64 {
	return 0
}

// generateStreetAddress implements CustomGenerator
func (*GoFakeItGenerator) generateStreetAddress() string {
	return ""
}

// generateSubjects implements CustomGenerator
func (*GoFakeItGenerator) generateSubjects() []string {
	return nil
}

// generateWorkHours implements CustomGenerator
func (*GoFakeItGenerator) generateWorkHours() int {
	return 0
}

// generateZip implements CustomGenerator
func (*GoFakeItGenerator) generateZip() string {
	return ""
}

// Package PowerTools
// Layer: Adapter
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package PowerTools

// Store SQL Adapters

func fromDao(dao powerToolsDao) *PowerTool {
	entity := NewPowerTool(Code(dao.Code), dao.Type, dao.Rate)
	return entity
}

func toDao(domain PowerTool) *powerToolsDao {
	return &powerToolsDao{
		Code: string(domain.Code),
		Rate: domain.Rate,
		Type: domain.Type,
	}
}

func toPresentation(v *PowerTool) *PowerToolDto {
	return &PowerToolDto{
		Code: string(v.Code),
		Type: v.Type,
		Rate: v.Rate,
	}
}

func fromPresentation(v PowerToolDto) *PowerTool {
	return NewPowerTool(Code(v.Code), v.Type, v.Rate)
}

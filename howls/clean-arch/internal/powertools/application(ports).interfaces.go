// Package PowerTools
// Layer: Application Business
// Interfaces define abstractions so layers can be referred
// 1) Inner layers define a set of interface
// 2) Outer layers implement interfaces
// 3) Dependency injection patter is used to provide instances to inner layers
// 4) Inner layers act on outer layers thought this interfaces abstractions
package PowerTools

// INPUT Ports

// Reader interface.
type Reader interface {
	Read(code Code) (*PowerTool, error)
	List() ([]*PowerTool, error)
}

// Writer interface.
type Writer interface {
	Create(ent *PowerTool) error
	Update(ent *PowerTool) error
	Delete(code Code) error
}

// Repository interface.
type Repository interface {
	Reader
	Writer
}

// OUTPUT Ports

// Service UseCase interface.
type Service interface {
	Get(code Code) (*PowerTool, error)
	List() ([]*PowerTool, error)
	Add(id Code, toolType string, rate int) (*PowerTool, error)
	Update(ent *PowerTool) error
	Delete(code Code) error
}

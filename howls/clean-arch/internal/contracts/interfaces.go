// Package contracts
// Layer: Application Business
// Interfaces define abstractions so layers can be referred
// 1) Inner layers define a set of interface
// 2) Outer layers implement interfaces
// 3) Dependency injection patter is used to provide instances to inner layers
// 4) Inner layers act on outer layers thought this interfaces abstractions
package contracts

//Repository interface
type Repository interface {
	Read(id ContractType) (*Contract, error)
	List() ([]*Contract, error)
}

//Service interface
type Service interface {
	Get(id ContractType) (*Contract, error)
	List() ([]*Contract, error)
}

package api

// #SOLID: S
// The only possible reason to change this file is to change interfaces related to the business layer
// This indicates that the single responsibility principle has been applied

// #SOLID: D
// The API code depends on this interface, the interface would be implemented by lower level layers
// This means that API code (higher-level code) is not depended on the lower-level layer, instead the lower-level needs to satisfy higher-level api
// The terminology used in the interface is based on the current layer's terminology, so the current layer would not be affected by the changes in lower-layer code
// And this is an example of applying Dependency Inversion principle
// This pattern has been repeated throughout the project
type balanceManager interface {
	Create(accountsNum int) error
	GetAll() (int64, error)
	Get(accId int) (int, error)
	// TODO rename add with another name
	AddToAll(increment int) error
	Add(increment int, accId int) error
}

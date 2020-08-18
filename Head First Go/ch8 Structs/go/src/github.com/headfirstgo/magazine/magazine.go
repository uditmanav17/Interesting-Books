// Package magazine defines structue Subscriber
package magazine

// Subscriber defines subscriber structure
type Subscriber struct { // name and attributes must be Title
	Name   string
	Rate   float64
	Active bool
}

// Address defines Address structure
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Employee defines Employee structure
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

// Employee2 defines employee address as Anonymous struct
type Employee2 struct {
	Name   string
	Salary float64
	Address
}

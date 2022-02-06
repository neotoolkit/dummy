package model

type Query struct {
	SelectionSet SelectionSet
}

type SelectionSet []SelectionField

type SelectionField struct {
	Name         string
	SelectionSet SelectionSet
}

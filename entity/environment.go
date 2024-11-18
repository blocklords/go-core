package entity

import "strings"

type Environment string

const (
	EDemo Environment = `demo`
	EBeta Environment = `beta`
	EPre  Environment = `pre`
	EMain Environment = `main`
)

func (e Environment) String() string {
	return strings.ToLower(string(e))
}

func (e Environment) AuthHostPrefix() string {
	switch e {
	case EBeta:
		return "beta-"
	case EDemo:
		return "tomato-"
	case EPre:
		return "pre-"
	}
	return ""
}

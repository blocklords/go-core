package entity

import "strings"

type Environment string

const (
	EDemo Environment = `demo`
	EBeta Environment = `beta`
	EPre  Environment = `pre`
	ESteam Environment = `steam`
	ESteamMain Environment = `steam-main`
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
	case ESteam:
		return "steam-"
	case ESteamMain:
		return "steam-main-"
	}
	return ""
}

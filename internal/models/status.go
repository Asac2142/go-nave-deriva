// Package models contains data custom type definitions for "Nave a la deriva"
package models

// DamageSystem damage system type.
type DamageSystem string

// IssueKey damag system key.
type IssueKey string

const (
	NAV01  DamageSystem = "NAV-01"  // NAV01 specifies a damage in the navigation system.
	COM02  DamageSystem = "COM-02"  // COM02 specifies a damage in the communications system.
	LIFE03 DamageSystem = "LIFE-03" // LIFE03 specifies a damage in the life support system.
	ENG04  DamageSystem = "ENG-04"  // ENG04 specifies a damage in the engines.
	SHLD05 DamageSystem = "SHLD-05" // SHLD05 specifies a damage in a shield.
)

const (
	NAVIGATION      IssueKey = "navigation"       // NAVIGATION key.
	COMMUNICATIONS  IssueKey = "communications"   // COMMUNICATIONS key
	LIFESUPPORT     IssueKey = "life_support"     // LIFESUPPORT key
	ENGINES         IssueKey = "engines"          // ENGINES key
	DEFLECTORSHIELD IssueKey = "deflector_shield" // DEFLECTORSHIELD key
)

// DamageKeys slice holding keys.
var DamageKeys = [...]IssueKey{NAVIGATION, COMMUNICATIONS, LIFESUPPORT, ENGINES, DEFLECTORSHIELD}

// DamageSchema holds ship issues key-value pair.
var DamageSchema = map[IssueKey]DamageSystem{
	NAVIGATION:      NAV01,
	COMMUNICATIONS:  COM02,
	LIFESUPPORT:     LIFE03,
	ENGINES:         ENG04,
	DEFLECTORSHIELD: SHLD05,
}

// StatusInfo defines the store schema.
type StatusInfo struct {
	Status IssueKey
}

// ResponseStatus defines the response schema.
type ResponseStatus struct {
	DamagedSystem IssueKey `json:"damaged_system"`
}

// NewResponse returns an instance of ResponseStatus.
func NewResponse(dmg *IssueKey) *ResponseStatus {
	return &ResponseStatus{DamagedSystem: *dmg}
}

package votesmarttypes

type Election struct {
	ElectionID   string `json:"electionId"`
	Name         string `json:"name"`
	StateID      string `json:"stateId"`
	OfficeTypeID string `json:"officeTypeId"`
	Special      string `json:"special"`
	ElectionYear string `json:"electionYear"`
	Stage        []struct {
		StageID                 string `json:"stageId"`
		ElectionElectionstageID string `json:"electionElectionstageId"`
		Name                    string `json:"name"`
		StateID                 string `json:"stateId"`
		ElectionDate            string `json:"electionDate"`
		FilingDeadline          string `json:"filingDeadline"`
		NpatMailed              string `json:"npatMailed"`
	} `json:"stage"`
}

// ElectionGetElection is the response message for Election.getElection.
// See http://api.votesmart.org/docs/Election.html for usage details.
type ElectionGetElection struct {
	Elections struct {
		Election Election `json:"election"`
	} `json:"elections"`
}

func (ElectionGetElection) Method() string {
	return "Election.getElection"
}

// ElectionGetElectionByYearState is the response message for Election.getElectionByYearState.
// See http://api.votesmart.org/docs/Election.html for usage details.
type ElectionGetElectionByYearState struct {
	Elections struct {
		Election []Election `json:"election"`
	} `json:"elections"`
}

func (ElectionGetElectionByYearState) Method() string {
	return "Election.getElectionByYearState"
}

// ElectionGetElectionByZip is the response message for Election.getElectionByZip.
// See http://api.votesmart.org/docs/Election.html for usage details.
type ElectionGetElectionByZip struct {
	Elections struct {
		Election []Election `json:"election"`
	} `json:"elections"`
}

func (ElectionGetElectionByZip) Method() string {
	return "Election.getElectionByZip"
}

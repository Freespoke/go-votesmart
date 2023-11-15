package votesmarttypes

import (
	"encoding/json"
	"errors"
)

type Election struct {
	ElectionID   string         `json:"electionId"`
	Name         string         `json:"name"`
	StateID      string         `json:"stateId"`
	OfficeTypeID string         `json:"officeTypeId"`
	Special      string         `json:"special"`
	ElectionYear string         `json:"electionYear"`
	Stage        MaybeStageList `json:"stage"`
}

type MaybeElectionList []Election

func (m *MaybeElectionList) UnmarshalJSON(in []byte) error {
	var single Election
	if err := json.Unmarshal(in, &single); err == nil {
		*m = []Election{single}
		return nil
	}

	var list []Election
	if err := json.Unmarshal(in, &list); err == nil {
		*m = list
		return nil
	}

	return errors.New("unexpected election response format")
}

type Stage struct {
	StageID                 string `json:"stageId"`
	ElectionElectionstageID string `json:"electionElectionstageId"`
	Name                    string `json:"name"`
	StateID                 string `json:"stateId"`
	ElectionDate            string `json:"electionDate"`
	FilingDeadline          string `json:"filingDeadline"`
	NpatMailed              string `json:"npatMailed"`
}

type MaybeStageList []Stage

func (m *MaybeStageList) UnmarshalJSON(in []byte) error {
	var single Stage
	if err := json.Unmarshal(in, &single); err == nil {
		*m = []Stage{single}
		return nil
	}

	var list []Stage
	if err := json.Unmarshal(in, &list); err == nil {
		*m = list
		return nil
	}

	return errors.New("unexpected stage response format")
}

// ElectionGetElection is the response message for Election.getElection.
// See http://api.votesmart.org/docs/Election.html for usage details.
type ElectionGetElection struct {
	Elections struct {
		Election MaybeElectionList `json:"election"`
	} `json:"elections"`
}

func (ElectionGetElection) Method() string {
	return "Election.getElection"
}

// ElectionGetElectionByYearState is the response message for Election.getElectionByYearState.
// See http://api.votesmart.org/docs/Election.html for usage details.
type ElectionGetElectionByYearState struct {
	Elections struct {
		Election MaybeElectionList `json:"election"`
	} `json:"elections"`
}

func (ElectionGetElectionByYearState) Method() string {
	return "Election.getElectionByYearState"
}

// ElectionGetElectionByZip is the response message for Election.getElectionByZip.
// See http://api.votesmart.org/docs/Election.html for usage details.
type ElectionGetElectionByZip struct {
	Elections struct {
		Election MaybeElectionList `json:"election"`
	} `json:"elections"`
}

func (ElectionGetElectionByZip) Method() string {
	return "Election.getElectionByZip"
}

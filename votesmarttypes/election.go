package votesmarttypes

type ElectionGetElection struct {
	Elections struct {
		Election struct {
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
		} `json:"election"`
	} `json:"elections"`
}

type ElectionGetElectionByYearState struct {
	Elections struct {
		Election []struct {
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
		} `json:"election"`
	} `json:"elections"`
}

type ElectionGetElectionByZip struct {
	Elections struct {
		Election []struct {
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
		} `json:"election"`
	} `json:"elections"`
}

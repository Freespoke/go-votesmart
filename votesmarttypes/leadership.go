package votesmarttypes

// LeadershipGetPositions is the response message for Leadership.getPositions.
// See http://api.votesmart.org/docs/Leadership.html for usage details.
type LeadershipGetPositions struct {
	Leadership struct {
		Position []struct {
			LeadershipID string `json:"leadershipId"`
			Name         string `json:"name"`
			OfficeID     string `json:"officeId"`
			OfficeName   string `json:"officeName"`
		} `json:"position"`
	} `json:"leadership"`
}

func (LeadershipGetPositions) Method() string {
	return "Leadership.getPositions"
}

// LeadershipGetOfficials is the response message for Leadership.getOfficials.
// See http://api.votesmart.org/docs/Leadership.html for usage details.
type LeadershipGetOfficials struct {
	Leaders struct {
		Leader struct {
			CandidateID string `json:"candidateId"`
			FirstName   string `json:"firstName"`
			MiddleName  string `json:"middleName"`
			LastName    string `json:"lastName"`
			Suffix      string `json:"suffix"`
			Position    string `json:"position"`
			OfficeID    string `json:"officeId"`
			Title       string `json:"title"`
		} `json:"leader"`
	} `json:"leaders"`
}

func (LeadershipGetOfficials) Method() string {
	return "Leadership.getOfficials"
}

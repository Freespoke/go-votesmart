package votesmarttypes

// LocalGetCounties is the response message for Local.getCounties.
// See http://api.votesmart.org/docs/Local.html for usage details.
type LocalGetCounties struct {
	Counties struct {
		County []struct {
			LocalID string `json:"localId"`
			Name    string `json:"name"`
			URL     string `json:"url"`
		} `json:"county"`
	} `json:"counties"`
}

func (LocalGetCounties) Method() string {
	return "Local.getCounties"
}

// LocalGetCities is the response message for Local.getCities.
// See http://api.votesmart.org/docs/Local.html for usage details.
type LocalGetCities struct {
	Cities struct {
		City []struct {
			LocalID string `json:"localId"`
			Name    string `json:"name"`
			URL     string `json:"url"`
		} `json:"city"`
	} `json:"cities"`
}

func (LocalGetCities) Method() string {
	return "Local.getCities"
}

// LocalGetOfficials is the response message for Local.getOfficials.
// See http://api.votesmart.org/docs/Local.html for usage details.
type LocalGetOfficials struct {
	CandidateList struct {
		Candidate MaybeList[Candidate] `json:"candidate"`
	} `json:"candidateList"`
}

func (LocalGetOfficials) Method() string {
	return "Local.getOfficials"
}

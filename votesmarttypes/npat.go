package votesmarttypes

// NpatGetNpat is the response message for Npat.getNpat.
// See http://api.votesmart.org/docs/Npat.html for usage details.
type NpatGetNpat struct {
	Npat struct {
		CandidateID   string `json:"candidateId"`
		Passed        string `json:"passed"`
		NpatName      string `json:"npatName"`
		ElectionName  string `json:"electionName"`
		ElectionYear  string `json:"electionYear"`
		ElectionDate  string `json:"electionDate"`
		ElectionStage string `json:"electionStage"`
		SurveyMessage string `json:"surveyMessage"`
		Candidate     string `json:"candidate"`
		// TODO: Fill this in with a real data representation.
		Section any `json:"section"`
	} `json:"npat"`
}

func (NpatGetNpat) Method() string {
	return "Npat.getNpat"
}

package votesmarttypes

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

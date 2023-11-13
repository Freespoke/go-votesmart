package votesmarttypes

type Candidate struct {
	CandidateID          string `json:"candidateId"`
	FirstName            string `json:"firstName"`
	NickName             string `json:"nickName"`
	MiddleName           string `json:"middleName"`
	PreferredName        string `json:"preferredName"`
	LastName             string `json:"lastName"`
	Suffix               string `json:"suffix"`
	Title                string `json:"title"`
	BallotName           string `json:"ballotName"`
	ElectionParties      string `json:"electionParties"`
	ElectionStatus       string `json:"electionStatus"`
	ElectionStage        string `json:"electionStage"`
	ElectionDistrictID   string `json:"electionDistrictId"`
	ElectionDistrictName string `json:"electionDistrictName"`
	ElectionOffice       string `json:"electionOffice"`
	ElectionOfficeID     string `json:"electionOfficeId"`
	ElectionStateID      string `json:"electionStateId"`
	ElectionOfficeTypeID string `json:"electionOfficeTypeId"`
	ElectionYear         string `json:"electionYear"`
	ElectionSpecial      string `json:"electionSpecial"`
	ElectionDate         string `json:"electionDate"`
	OfficeParties        string `json:"officeParties"`
	OfficeStatus         string `json:"officeStatus"`
	OfficeDistrictID     string `json:"officeDistrictId"`
	OfficeDistrictName   string `json:"officeDistrictName"`
	OfficeStateID        string `json:"officeStateId"`
	OfficeID             string `json:"officeId"`
	OfficeName           string `json:"officeName"`
	OfficeTypeID         string `json:"officeTypeId"`
	RunningMateID        string `json:"runningMateId"`
	RunningMateName      string `json:"runningMateName"`
}

// CandidatesGetByOfficeState is the response message for Candidates.getByOfficeState.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByOfficeState struct {
	CandidateList struct {
		Candidate []Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByOfficeState) Method() string {
	return "Candidates.getByOfficeState"
}

// CandidatesGetByOfficeTypeState is the response message for Candidates.getByOfficeTypeState.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByOfficeTypeState struct {
	CandidateList struct {
		Candidate []Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByOfficeTypeState) Method() string {
	return "Candidates.getByOfficeTypeState"
}

// CandidatesGetByLastname is the response message for Candidates.getByLastname.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByLastname struct {
	CandidateList struct {
		Candidate []Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByLastname) Method() string {
	return "Candidates.getByLastname"
}

// CandidatesGetByLevenshtein is the response message for Candidates.getByLevenshtein.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByLevenshtein struct {
	CandidateList struct {
		Candidate Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByLevenshtein) Method() string {
	return "Candidates.getByLevenshtein"
}

// CandidatesGetByElection is the response message for Candidates.getByElection.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByElection struct {
	CandidateList struct {
		Candidate []Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByElection) Method() string {
	return "Candidates.getByElection"
}

// CandidatesGetByDistrict is the response message for Candidates.getByDistrict.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByDistrict struct {
	CandidateList struct {
		Candidate []Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByDistrict) Method() string {
	return "Candidates.getByDistrict"
}

// CandidatesGetByZip is the response message for Candidates.getByZip.
// See http://api.votesmart.org/docs/Candidates.html for usage details.
type CandidatesGetByZip struct {
	CandidateList struct {
		ZipMessage string      `json:"zipMessage"`
		Candidate  []Candidate `json:"candidate"`
	} `json:"candidateList"`
}

func (CandidatesGetByZip) Method() string {
	return "Candidates.getByZip"
}

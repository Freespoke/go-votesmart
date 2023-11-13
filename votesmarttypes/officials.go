package votesmarttypes

// OfficialsGetStatewide is the response message for Officials.getStatewide.
// See http://api.votesmart.org/docs/Officials.html for usage details.
type OfficialsGetStatewide struct {
	CandidateList struct {
		Candidate []struct {
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
		} `json:"candidate"`
	} `json:"candidateList"`
}

func (OfficialsGetStatewide) Method() string {
	return "Officials.getStatewide"
}

// OfficialsGetByOfficeState is the response message for Officials.getByOfficeState.
// See http://api.votesmart.org/docs/Officials.html for usage details.
type OfficialsGetByOfficeState struct {
	CandidateList struct {
		Candidate struct {
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
		} `json:"candidate"`
	} `json:"candidateList"`
}

func (OfficialsGetByOfficeState) Method() string {
	return "Officials.getByOfficeState"
}

// OfficialsGetByOfficeTypeState is the response message for Officials.getByOfficeTypeState.
// See http://api.votesmart.org/docs/Officials.html for usage details.
type OfficialsGetByOfficeTypeState struct {
	CandidateList struct {
		Candidate []struct {
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
		} `json:"candidate"`
	} `json:"candidateList"`
}

func (OfficialsGetByOfficeTypeState) Method() string {
	return "Officials.getByOfficeTypeState"
}

// OfficialsGetByLastname is the response message for Officials.getByLastname.
// See http://api.votesmart.org/docs/Officials.html for usage details.
type OfficialsGetByLastname struct {
	CandidateList struct {
		Candidate struct {
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
		} `json:"candidate"`
	} `json:"candidateList"`
}

func (OfficialsGetByLastname) Method() string {
	return "Officials.getByLastname"
}

// OfficialsGetByLevenshtein is the response message for Officials.getByLevenshtein.
// See http://api.votesmart.org/docs/Officials.html for usage details.
type OfficialsGetByLevenshtein struct {
	CandidateList struct {
		Candidate struct {
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
		} `json:"candidate"`
	} `json:"candidateList"`
}

func (OfficialsGetByLevenshtein) Method() string {
	return "Officials.getByLevenshtein"
}

// OfficialsGetByZip is the response message for Officials.getByZip.
// See http://api.votesmart.org/docs/Officials.html for usage details.
type OfficialsGetByZip struct {
	CandidateList struct {
		ZipMessage string `json:"zipMessage"`
		Candidate  []struct {
			CandidateID        string `json:"candidateId"`
			FirstName          string `json:"firstName"`
			NickName           string `json:"nickName"`
			MiddleName         string `json:"middleName"`
			PreferredName      string `json:"preferredName"`
			LastName           string `json:"lastName"`
			Suffix             string `json:"suffix"`
			Title              string `json:"title"`
			OfficeParties      string `json:"officeParties"`
			OfficeStatus       string `json:"officeStatus"`
			OfficeDistrictID   string `json:"officeDistrictId"`
			OfficeDistrictName string `json:"officeDistrictName"`
			OfficeStateID      string `json:"officeStateId"`
			OfficeID           string `json:"officeId"`
			OfficeName         string `json:"officeName"`
			OfficeTypeID       string `json:"officeTypeId"`
		} `json:"candidate"`
	} `json:"candidateList"`
}

func (OfficialsGetByZip) Method() string {
	return "Officials.getByZip"
}

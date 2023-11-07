package votesmarttypes

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

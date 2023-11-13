package votesmarttypes

// CandidateBioGetBio is the response message for CandidateBio.getBio.
// See http://api.votesmart.org/docs/CandidateBio for usage details.
type CandidateBioGetBio struct {
	Bio struct {
		Candidate struct {
			CandidateID   string `json:"candidateId"`
			CrpID         string `json:"crpId"`
			Photo         string `json:"photo"`
			FirstName     string `json:"firstName"`
			NickName      string `json:"nickName"`
			MiddleName    string `json:"middleName"`
			PreferredName string `json:"preferredName"`
			LastName      string `json:"lastName"`
			Suffix        string `json:"suffix"`
			BirthDate     string `json:"birthDate"`
			BirthPlace    string `json:"birthPlace"`
			Pronunciation string `json:"pronunciation"`
			Gender        string `json:"gender"`
			Family        string `json:"family"`
			HomeCity      string `json:"homeCity"`
			HomeState     string `json:"homeState"`
			Religion      string `json:"religion"`
			SpecialMsg    string `json:"specialMsg"`
		} `json:"candidate"`
		Office struct {
			Name       []string `json:"name"`
			Parties    string   `json:"parties"`
			Title      string   `json:"title"`
			ShortTitle string   `json:"shortTitle"`
			Type       string   `json:"type"`
			Status     string   `json:"status"`
			FirstElect string   `json:"firstElect"`
			LastElect  string   `json:"lastElect"`
			NextElect  string   `json:"nextElect"`
			TermStart  string   `json:"termStart"`
			TermEnd    string   `json:"termEnd"`
			District   string   `json:"district"`
			DistrictID string   `json:"districtId"`
			StateID    string   `json:"stateId"`
		} `json:"office"`
		Election struct {
			Office     string `json:"office"`
			OfficeType string `json:"officeType"`
			Parties    string `json:"parties"`
			District   string `json:"district"`
			DistrictID string `json:"districtId"`
			StateID    string `json:"stateId"`
			Status     string `json:"status"`
			BallotName string `json:"ballotName"`
		} `json:"election"`
	} `json:"bio"`
}

func (CandidateBioGetBio) Method() string {
	return "CandidateBio.getBio"
}

// CandidateBioGetDetailedBio is the response message for CandidateBio.getDetailedBio.
// See http://api.votesmart.org/docs/CandidateBio for usage details.
type CandidateBioGetDetailedBio struct {
	Bio struct {
		Candidate struct {
			CandidateID   string `json:"candidateId"`
			CrpID         string `json:"crpId"`
			Photo         string `json:"photo"`
			FirstName     string `json:"firstName"`
			NickName      string `json:"nickName"`
			MiddleName    string `json:"middleName"`
			PreferredName string `json:"preferredName"`
			LastName      string `json:"lastName"`
			Suffix        string `json:"suffix"`
			BirthDate     string `json:"birthDate"`
			BirthPlace    string `json:"birthPlace"`
			Pronunciation string `json:"pronunciation"`
			Gender        string `json:"gender"`
			Family        string `json:"family"`
			HomeCity      string `json:"homeCity"`
			HomeState     string `json:"homeState"`
			Education     struct {
				Institution []struct {
					Degree   string `json:"degree"`
					Field    string `json:"field"`
					School   string `json:"school"`
					Span     string `json:"span"`
					Gpa      string `json:"gpa"`
					FullText string `json:"fullText"`
				} `json:"institution"`
			} `json:"education"`
			Profession struct {
				Experience []struct {
					Title        string `json:"title"`
					Organization string `json:"organization"`
					Span         string `json:"span"`
					Special      string `json:"special"`
					District     string `json:"district"`
					FullText     string `json:"fullText"`
				} `json:"experience"`
			} `json:"profession"`
			Political struct {
				Experience []struct {
					Title        string `json:"title"`
					Organization string `json:"organization"`
					Span         string `json:"span"`
					Special      string `json:"special"`
					District     string `json:"district"`
					FullText     string `json:"fullText"`
				} `json:"experience"`
			} `json:"political"`
			CongMembership struct {
				Experience []struct {
					Title        string `json:"title"`
					Organization string `json:"organization"`
					Span         string `json:"span"`
					Special      string `json:"special"`
					District     string `json:"district"`
					FullText     string `json:"fullText"`
				} `json:"experience"`
			} `json:"congMembership"`
			OrgMembership struct {
				Experience []struct {
					Title        string `json:"title"`
					Organization string `json:"organization"`
					Span         string `json:"span"`
					Special      string `json:"special"`
					District     string `json:"district"`
					FullText     string `json:"fullText"`
				} `json:"experience"`
			} `json:"orgMembership"`
			Religion   string `json:"religion"`
			SpecialMsg string `json:"specialMsg"`
		} `json:"candidate"`
		Office struct {
			Name       []string `json:"name"`
			Parties    string   `json:"parties"`
			Title      string   `json:"title"`
			ShortTitle string   `json:"shortTitle"`
			Type       string   `json:"type"`
			Status     string   `json:"status"`
			FirstElect string   `json:"firstElect"`
			LastElect  string   `json:"lastElect"`
			NextElect  string   `json:"nextElect"`
			TermStart  string   `json:"termStart"`
			TermEnd    string   `json:"termEnd"`
			District   string   `json:"district"`
			DistrictID string   `json:"districtId"`
			StateID    string   `json:"stateId"`
		} `json:"office"`
		Election struct {
			Office     string `json:"office"`
			OfficeType string `json:"officeType"`
			Parties    string `json:"parties"`
			District   string `json:"district"`
			DistrictID string `json:"districtId"`
			StateID    string `json:"stateId"`
			Status     string `json:"status"`
			BallotName string `json:"ballotName"`
		} `json:"election"`
	} `json:"bio"`
}

func (CandidateBioGetDetailedBio) Method() string {
	return "CandidateBio.getDetailedBio"
}

// CandidateBioGetAdditionalBio is the response message for CandidateBio.getAdditionalBio.
// See http://api.votesmart.org/docs/CandidateBio for usage details.
type CandidateBioGetAdditionalBio struct {
	AddlBio struct {
		Candidate struct {
			ShortTitle string `json:"shortTitle"`
			FirstName  string `json:"firstName"`
			NickName   string `json:"nickName"`
			MiddleName string `json:"middleName"`
			LastName   string `json:"lastName"`
			Suffix     string `json:"suffix"`
		} `json:"candidate"`
		Additional struct {
			Item []struct {
				Name string `json:"name"`
				Data string `json:"data"`
			} `json:"item"`
		} `json:"additional"`
	} `json:"addlBio"`
}

func (CandidateBioGetAdditionalBio) Method() string {
	return "CandidateBio.getAdditionalBio"
}

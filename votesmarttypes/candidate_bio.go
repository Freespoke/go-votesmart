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
		Office   MaybeList[BioOffice]   `json:"office"`
		Election MaybeList[BioElection] `json:"election"`
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
			CandidateID    string                  `json:"candidateId"`
			CrpID          string                  `json:"crpId"`
			Photo          string                  `json:"photo"`
			FirstName      string                  `json:"firstName"`
			NickName       string                  `json:"nickName"`
			MiddleName     string                  `json:"middleName"`
			PreferredName  string                  `json:"preferredName"`
			LastName       string                  `json:"lastName"`
			Suffix         string                  `json:"suffix"`
			BirthDate      string                  `json:"birthDate"`
			BirthPlace     string                  `json:"birthPlace"`
			Pronunciation  string                  `json:"pronunciation"`
			Gender         string                  `json:"gender"`
			Family         string                  `json:"family"`
			HomeCity       string                  `json:"homeCity"`
			HomeState      string                  `json:"homeState"`
			Education      MaybeList[BioEducation] `json:"education"`
			Profession     MaybeList[BioItem]      `json:"profession"`
			Political      MaybeList[BioItem]      `json:"political"`
			CongMembership MaybeList[BioItem]      `json:"congMembership"`
			OrgMembership  MaybeList[BioItem]      `json:"orgMembership"`
			Religion       string                  `json:"religion"`
			SpecialMsg     string                  `json:"specialMsg"`
		} `json:"candidate"`
		Office   MaybeList[BioOffice]   `json:"office"`
		Election MaybeList[BioElection] `json:"election"`
	} `json:"bio"`
}

func (CandidateBioGetDetailedBio) Method() string {
	return "CandidateBio.getDetailedBio"
}

// CandidateBioGetAdditionalBio is the response message for CandidateBio.getAddlBio.
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
		Additional MaybeList[BioAdditional] `json:"additional"`
	} `json:"addlBio"`
}

func (CandidateBioGetAdditionalBio) Method() string {
	return "CandidateBio.getAddlBio"
}

type BioElection struct {
	Office     string `json:"office"`
	OfficeType string `json:"officeType"`
	Parties    string `json:"parties"`
	District   string `json:"district"`
	DistrictID string `json:"districtId"`
	StateID    string `json:"stateId"`
	Status     string `json:"status"`
	BallotName string `json:"ballotName"`
}

type BioOffice struct {
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
}

type BioAdditional struct {
	Item MaybeList[BioAdditionalItem] `json:"item"`
}

type BioAdditionalItem struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type BioEducation struct {
	Institution MaybeList[BioEducationInstitution] `json:"institution"`
}

type BioEducationInstitution struct {
	Degree   string `json:"degree"`
	Field    string `json:"field"`
	School   string `json:"school"`
	Span     string `json:"span"`
	Gpa      string `json:"gpa"`
	FullText string `json:"fullText"`
}

type BioItem struct {
	Experience MaybeList[BioItemExperience] `json:"experience"`
}

type BioItemExperience struct {
	Title        string `json:"title"`
	Organization string `json:"organization"`
	Span         string `json:"span"`
	Special      string `json:"special"`
	District     string `json:"district"`
	FullText     string `json:"fullText"`
}

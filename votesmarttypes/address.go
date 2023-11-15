package votesmarttypes

type WebAddress struct {
	WebAddressTypeID string `json:"webAddressTypeId"`
	WebAddressType   string `json:"webAddressType"`
	WebAddress       string `json:"webAddress"`
}

type OfficeAddress struct {
	Candidate struct {
		CandidateID string `json:"candidateId"`
		Title       string `json:"title"`
		FirstName   string `json:"firstName"`
		MiddleName  string `json:"middleName"`
		NickName    string `json:"nickName"`
		LastName    string `json:"lastName"`
		Suffix      string `json:"suffix"`
	} `json:"candidate"`
	Address struct {
		Type   string `json:"type"`
		TypeID string `json:"typeId"`
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		Zip    string `json:"zip"`
	} `json:"address"`
	Phone struct {
		Phone1    string `json:"phone1"`
		Phone2    string `json:"phone2"`
		Fax1      string `json:"fax1"`
		Fax2      string `json:"fax2"`
		TollFree  string `json:"tollFree"`
		Ttyd      string `json:"ttyd"`
		Cellphone string `json:"cellphone"`
	} `json:"phone"`
	Notes struct {
		ContactName  string `json:"contactName"`
		ContactTitle string `json:"contactTitle"`
	} `json:"notes"`
}

// AddressGetCampaign is the response message for Address.getCampaign.
// See http://api.votesmart.org/docs/Address.html for usage details.
type AddressGetCampaign struct {
	Address struct {
		Candidate struct {
			Title      string `json:"title"`
			FirstName  string `json:"firstName"`
			MiddleName string `json:"middleName"`
			NickName   string `json:"nickName"`
			LastName   string `json:"lastName"`
			Suffix     string `json:"suffix"`
		} `json:"candidate"`
		Office MaybeList[OfficeAddress] `json:"office"`
	} `json:"address"`
}

func (AddressGetCampaign) Method() string {
	return "Address.getCampaign"
}

// AddressGetCampaignWebAddress is the response message for Address.getCampaignWebAddress.
// See http://api.votesmart.org/docs/Address.html for usage details.
type AddressGetCampaignWebAddress struct {
	Webaddress struct {
		Candidate struct {
			Title      string `json:"title"`
			FirstName  string `json:"firstName"`
			MiddleName string `json:"middleName"`
			NickName   string `json:"nickName"`
			LastName   string `json:"lastName"`
			Suffix     string `json:"suffix"`
		} `json:"candidate"`
		Address MaybeList[WebAddress] `json:"address"`
	} `json:"webaddress"`
}

func (AddressGetCampaignWebAddress) Method() string {
	return "Address.getCampaignWebAddress"
}

// AddressGetCampaignByElection is the response message for Address.getCampaignByElection.
// See http://api.votesmart.org/docs/Address.html for usage details.
type AddressGetCampaignByElection struct {
	Address struct {
		Office MaybeList[OfficeAddress] `json:"office"`
	} `json:"address"`
}

func (AddressGetCampaignByElection) Method() string {
	return "Address.getCampaignByElection"
}

package votesmarttypes

type District struct {
	DistrictID string `json:"districtId"`
	Name       string `json:"name"`
	OfficeID   string `json:"officeId"`
	StateID    string `json:"stateId"`
}

// DistrictGetByOfficeState is the response message for District.getByOfficeState.
// See http://api.votesmart.org/docs/District.html for usage details.
type DistrictGetByOfficeState struct {
	DistrictList struct {
		District []District `json:"district"`
	} `json:"districtList"`
}

func (DistrictGetByOfficeState) Method() string {
	return "District.getByOfficeState"
}

// DistrictGetByZip is the response message for District.getByZip.
// See http://api.votesmart.org/docs/District.html for usage details.
type DistrictGetByZip struct {
	DistrictList struct {
		ZipMessage        string     `json:"zipMessage"`
		District          []District `json:"district"`
		ElectionDistricts struct {
			District MaybeList[District] `json:"district"`
		} `json:"electionDistricts"`
	} `json:"districtList"`
}

func (DistrictGetByZip) Method() string {
	return "District.getByZip"
}

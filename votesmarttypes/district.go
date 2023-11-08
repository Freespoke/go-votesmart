package votesmarttypes

type District struct {
	DistrictID string `json:"districtId"`
	Name       string `json:"name"`
	OfficeID   string `json:"officeId"`
	StateID    string `json:"stateId"`
}

type DistrictGetByOfficeState struct {
	DistrictList struct {
		District []District `json:"district"`
	} `json:"districtList"`
}

type DistrictGetByZip struct {
	DistrictList struct {
		ZipMessage        string     `json:"zipMessage"`
		District          []District `json:"district"`
		ElectionDistricts struct {
			District []District `json:"district"`
		} `json:"electionDistricts"`
	} `json:"districtList"`
}

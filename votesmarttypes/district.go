package votesmarttypes

type DistrictGetByOfficeState struct {
	DistrictList struct {
		District []struct {
			DistrictID string `json:"districtId"`
			Name       string `json:"name"`
			OfficeID   string `json:"officeId"`
			StateID    string `json:"stateId"`
		} `json:"district"`
	} `json:"districtList"`
}

type DistrictGetByZip struct {
	DistrictList struct {
		ZipMessage string `json:"zipMessage"`
		District   []struct {
			DistrictID string `json:"districtId"`
			Name       string `json:"name"`
			OfficeID   string `json:"officeId"`
			StateID    string `json:"stateId"`
		} `json:"district"`
		ElectionDistricts struct {
			District []struct {
				DistrictID string `json:"districtId"`
				Name       string `json:"name"`
				OfficeID   string `json:"officeId"`
				StateID    string `json:"stateId"`
			} `json:"district"`
		} `json:"electionDistricts"`
	} `json:"districtList"`
}

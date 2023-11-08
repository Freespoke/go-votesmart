package votesmarttypes

type Office struct {
	OfficeID       string `json:"officeId"`
	OfficeTypeID   string `json:"officeTypeId"`
	OfficeLevelID  string `json:"officeLevelId"`
	OfficeBranchID string `json:"officeBranchId"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	ShortTitle     string `json:"shortTitle"`
}

type OfficeGetTypes struct {
	OfficeTypes struct {
		Type []struct {
			OfficeTypeID   string `json:"officeTypeId"`
			OfficeLevelID  string `json:"officeLevelId"`
			OfficeBranchID string `json:"officeBranchId"`
			Name           string `json:"name"`
		} `json:"type"`
	} `json:"officeTypes"`
}

type OfficeGetBranches struct {
	Branches struct {
		Branch []struct {
			OfficeBranchID string `json:"officeBranchId"`
			Name           string `json:"name"`
		} `json:"branch"`
	} `json:"branches"`
}

type OfficeGetLevels struct {
	Levels struct {
		Level []struct {
			OfficeLevelID string `json:"officeLevelId"`
			Name          string `json:"name"`
		} `json:"level"`
	} `json:"levels"`
}

type OfficeGetOfficesByType struct {
	Offices struct {
		Office []Office `json:"office"`
	} `json:"offices"`
}

type OfficeGetOfficesByLevel struct {
	Offices struct {
		Office []Office `json:"office"`
	} `json:"offices"`
}

type OfficeGetOfficesByTypeLevel struct {
	Offices struct {
		Office []Office `json:"office"`
	} `json:"offices"`
}

type OfficeGetOfficesByBranchLevel struct {
	Offices struct {
		Office []Office `json:"office"`
	} `json:"offices"`
}

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

// OfficeGetTypes is the response message for Office.getTypes.
// See http://api.votesmart.org/docs/Office.html for usage details.
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

func (OfficeGetTypes) Method() string {
	return "Office.getTypes"
}

// OfficeGetBranches is the response message for Office.getBranches.
// See http://api.votesmart.org/docs/Office.html for usage details.
type OfficeGetBranches struct {
	Branches struct {
		Branch []struct {
			OfficeBranchID string `json:"officeBranchId"`
			Name           string `json:"name"`
		} `json:"branch"`
	} `json:"branches"`
}

func (OfficeGetBranches) Method() string {
	return "Office.getBranches"
}

// OfficeGetLevels is the response message for Office.getLevels.
// See http://api.votesmart.org/docs/Office.html for usage details.
type OfficeGetLevels struct {
	Levels struct {
		Level []struct {
			OfficeLevelID string `json:"officeLevelId"`
			Name          string `json:"name"`
		} `json:"level"`
	} `json:"levels"`
}

func (OfficeGetLevels) Method() string {
	return "Office.getLevels"
}

// OfficeGetOfficesByType is the response message for Office.getOfficesByType.
// See http://api.votesmart.org/docs/Office.html for usage details.
type OfficeGetOfficesByType struct {
	Offices struct {
		Office MaybeList[Office] `json:"office"`
	} `json:"offices"`
}

func (OfficeGetOfficesByType) Method() string {
	return "Office.getOfficesByType"
}

// OfficeGetOfficesByLevel is the response message for Office.getOfficesByLevel.
// See http://api.votesmart.org/docs/Office.html for usage details.
type OfficeGetOfficesByLevel struct {
	Offices struct {
		Office MaybeList[Office] `json:"office"`
	} `json:"offices"`
}

func (OfficeGetOfficesByLevel) Method() string {
	return "Office.getOfficesByLevel"
}

// OfficeGetOfficesByTypeLevel is the response message for Office.getOfficesByTypeLevel.
// See http://api.votesmart.org/docs/Office.html for usage details.
type OfficeGetOfficesByTypeLevel struct {
	Offices struct {
		Office MaybeList[Office] `json:"office"`
	} `json:"offices"`
}

func (OfficeGetOfficesByTypeLevel) Method() string {
	return "Office.getOfficesByTypeLevel"
}

// OfficeGetOfficesByBranchLevel is the response message for Office.getOfficesByBranchLevel.
// See http://api.votesmart.org/docs/Office.html for usage details.
type OfficeGetOfficesByBranchLevel struct {
	Offices struct {
		Office MaybeList[Office] `json:"office"`
	} `json:"offices"`
}

func (OfficeGetOfficesByBranchLevel) Method() string {
	return "Office.getOfficesByBranchLevel"
}

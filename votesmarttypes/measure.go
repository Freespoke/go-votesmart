package votesmarttypes

// MeasureGetMeasuresByYearState is the response message for Measure.getMeasuresByYearState.
// See http://api.votesmart.org/docs/Measure.html for usage details.
type MeasureGetMeasuresByYearState struct {
	Measures struct {
		Measure []struct {
			MeasureID   string `json:"measureId"`
			MeasureCode string `json:"measureCode"`
			Title       string `json:"title"`
			Outcome     string `json:"outcome"`
		} `json:"measure"`
	} `json:"measures"`
}

func (MeasureGetMeasuresByYearState) Method() string {
	return "Measure.getMeasuresByYearState"
}

// MeasureGetMeasure is the response message for Measure.getMeasure.
// See http://api.votesmart.org/docs/Measure.html for usage details.
type MeasureGetMeasure struct {
	Measure struct {
		MeasureID    string `json:"measureId"`
		MeasureCode  string `json:"measureCode"`
		Title        string `json:"title"`
		ElectionDate string `json:"electionDate"`
		ElectionType string `json:"electionType"`
		Source       string `json:"source"`
		URL          string `json:"url"`
		Summary      string `json:"summary"`
		SummaryURL   string `json:"summaryUrl"`
		MeasureText  string `json:"measureText"`
		TextURL      string `json:"textUrl"`
		ProURL       string `json:"proUrl"`
		ConURL       string `json:"conUrl"`
		Yes          string `json:"yes"`
		No           string `json:"no"`
		Outcome      string `json:"outcome"`
	} `json:"measure"`
}

func (MeasureGetMeasure) Method() string {
	return "Measure.getMeasure"
}

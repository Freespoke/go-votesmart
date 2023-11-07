package votesmarttypes

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

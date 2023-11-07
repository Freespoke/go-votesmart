package votesmarttypes

type StateGetStateIDs struct {
	StateList struct {
		List struct {
			State []struct {
				StateID string `json:"stateId"`
				Name    string `json:"name"`
			} `json:"state"`
		} `json:"list"`
	} `json:"stateList"`
}

type StateGetState struct {
	State struct {
		Details struct {
			StateID     string `json:"stateId"`
			StateType   string `json:"stateType"`
			Name        string `json:"name"`
			NickName    string `json:"nickName"`
			Capital     string `json:"capital"`
			Area        string `json:"area"`
			Population  string `json:"population"`
			Statehood   string `json:"statehood"`
			Motto       string `json:"motto"`
			Flower      string `json:"flower"`
			Tree        string `json:"tree"`
			Bird        string `json:"bird"`
			HighPoint   string `json:"highPoint"`
			LowPoint    string `json:"lowPoint"`
			Bicameral   string `json:"bicameral"`
			UpperLegis  string `json:"upperLegis"`
			LowerLegis  string `json:"lowerLegis"`
			LtGov       string `json:"ltGov"`
			Senators    string `json:"senators"`
			Reps        string `json:"reps"`
			TermLimit   string `json:"termLimit"`
			TermLength  string `json:"termLength"`
			BillURL     string `json:"billUrl"`
			VoteURL     string `json:"voteUrl"`
			VoterReg    string `json:"voterReg"`
			PrimaryDate string `json:"primaryDate"`
			GeneralDate string `json:"generalDate"`
			LargestCity string `json:"largestCity"`
			RollUpper   string `json:"rollUpper"`
			RollLower   string `json:"rollLower"`
			UsCircuit   string `json:"usCircuit"`
		} `json:"details"`
	} `json:"state"`
}

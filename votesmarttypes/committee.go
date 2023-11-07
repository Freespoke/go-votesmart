package votesmarttypes

type CommitteeGetTypes struct {
	CommitteeTypes struct {
		Type []struct {
			CommitteeTypeID string `json:"committeeTypeId"`
			Name            string `json:"name"`
		} `json:"type"`
	} `json:"committeeTypes"`
}

type CommitteeGetCommitteesByTypeState struct {
	Committees struct {
		Committee []struct {
			CommitteeID     string `json:"committeeId"`
			ParentID        string `json:"parentId"`
			StateID         string `json:"stateId"`
			CommitteetypeID string `json:"committeetypeId"`
			Name            string `json:"name"`
		} `json:"committee"`
	} `json:"committees"`
}

type CommitteeGetCommittee struct {
	Committee struct {
		CommitteeID     string `json:"committeeId"`
		ParentID        string `json:"parentId"`
		StateID         string `json:"stateId"`
		CommitteeTypeID string `json:"committeeTypeId"`
		Name            string `json:"name"`
		Jurisdiction    string `json:"jurisdiction"`
		Contact         struct {
			Address1     string `json:"address1"`
			Address2     string `json:"address2"`
			City         string `json:"city"`
			State        string `json:"state"`
			Zip          string `json:"zip"`
			Phone        string `json:"phone"`
			Fax          string `json:"fax"`
			Email        string `json:"email"`
			URL          string `json:"url"`
			StaffContact string `json:"staffContact"`
		} `json:"contact"`
	} `json:"committee"`
}

type CommitteeGetCommitteeMembers struct {
	CommitteeMembers struct {
		Committee struct {
			CommitteeID string `json:"committeeId"`
			ParentID    string `json:"parentId"`
			Name        string `json:"name"`
		} `json:"committee"`
		Member []struct {
			CandidateID string `json:"candidateId"`
			Title       string `json:"title"`
			FirstName   string `json:"firstName"`
			MiddleName  string `json:"middleName"`
			LastName    string `json:"lastName"`
			Suffix      string `json:"suffix"`
			Party       string `json:"party"`
			Position    string `json:"position"`
		} `json:"member"`
	} `json:"committeeMembers"`
}

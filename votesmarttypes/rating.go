package votesmarttypes

type RatingGetCategories struct {
	Categories struct {
		Category []struct {
			CategoryID string `json:"categoryId"`
			Name       string `json:"name"`
		} `json:"category"`
	} `json:"categories"`
}

type RatingGetSigList struct {
	Sigs struct {
		Sig []struct {
			SigID    string `json:"sigId"`
			ParentID string `json:"parentId"`
			Name     string `json:"name"`
		} `json:"sig"`
	} `json:"sigs"`
}

type RatingGetSig struct {
	Sig struct {
		SigID       string `json:"sigId"`
		ParentID    string `json:"parentId"`
		StateID     string `json:"stateId"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
		City        string `json:"city"`
		State       string `json:"state"`
		Zip         string `json:"zip"`
		Phone1      string `json:"phone1"`
		Phone2      string `json:"phone2"`
		Fax         string `json:"fax"`
		Email       string `json:"email"`
		URL         string `json:"url"`
		ContactName string `json:"contactName"`
	} `json:"sig"`
}

type RatingGetSigRatings struct {
	SigRatings struct {
		Sig struct {
			SigID string `json:"sigId"`
			Name  string `json:"name"`
		} `json:"sig"`
		Rating []struct {
			RatingID   string `json:"ratingId"`
			Timespan   string `json:"timespan"`
			RatingName string `json:"ratingName"`
			RatingText string `json:"ratingText"`
		} `json:"rating"`
	} `json:"sigRatings"`
}

type RatingGetCandidateRating struct {
	CandidateRating struct {
		Candidate struct {
			Title      string `json:"title"`
			FirstName  string `json:"firstName"`
			MiddleName string `json:"middleName"`
			LastName   string `json:"lastName"`
			Suffix     string `json:"suffix"`
			Office     string `json:"office"`
		} `json:"candidate"`
		Rating []struct {
			SigID      string `json:"sigId"`
			RatingID   string `json:"ratingId"`
			Categories struct {
				Category []struct {
					CategoryID string `json:"categoryId"`
					Name       string `json:"name"`
				} `json:"category"`
			} `json:"categories"`
			Timespan   string `json:"timespan"`
			Rating     string `json:"rating"`
			RatingName string `json:"ratingName"`
			RatingText string `json:"ratingText"`
		} `json:"rating"`
	} `json:"candidateRating"`
}

type RatingGetRating struct {
	Rating struct {
		CandidateRating []struct {
			CandidateID string `json:"candidateId"`
			Rating      string `json:"rating"`
		} `json:"candidateRating"`
	} `json:"rating"`
}

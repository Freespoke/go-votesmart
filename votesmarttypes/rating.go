package votesmarttypes

// RatingGetCategories is the response message for Rating.getCategories.
// See http://api.votesmart.org/docs/Rating.html for usage details.
type RatingGetCategories struct {
	Categories struct {
		Category []struct {
			CategoryID string `json:"categoryId"`
			Name       string `json:"name"`
		} `json:"category"`
	} `json:"categories"`
}

func (RatingGetCategories) Method() string {
	return "Rating.getCategories"
}

// RatingGetSigList is the response message for Rating.getSigList.
// See http://api.votesmart.org/docs/Rating.html for usage details.
type RatingGetSigList struct {
	Sigs struct {
		Sig []struct {
			SigID    string `json:"sigId"`
			ParentID string `json:"parentId"`
			Name     string `json:"name"`
		} `json:"sig"`
	} `json:"sigs"`
}

func (RatingGetSigList) Method() string {
	return "Rating.getSigList"
}

// RatingGetSig is the response message for Rating.getSig.
// See http://api.votesmart.org/docs/Rating.html for usage details.
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

func (RatingGetSig) Method() string {
	return "Rating.getSig"
}

// RatingGetSigRatings is the response message for Rating.getSigRatings.
// See http://api.votesmart.org/docs/Rating.html for usage details.
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

func (RatingGetSigRatings) Method() string {
	return "Rating.getSigRatings"
}

// RatingGetCandidateRating is the response message for Rating.getCandidateRating.
// See http://api.votesmart.org/docs/Rating.html for usage details.
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

func (RatingGetCandidateRating) Method() string {
	return "Rating.getCandidateRating"
}

// RatingGetRating is the response message for Rating.getRating.
// See http://api.votesmart.org/docs/Rating.html for usage details.
type RatingGetRating struct {
	Rating struct {
		CandidateRating []struct {
			CandidateID string `json:"candidateId"`
			Rating      string `json:"rating"`
		} `json:"candidateRating"`
	} `json:"rating"`
}

func (RatingGetRating) Method() string {
	return "Rating.getRating"
}

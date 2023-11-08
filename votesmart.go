package votesmart

import (
	"context"
	"net/http"
	"net/url"

	"dev.freespoke.com/go-votesmart/internal"
	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

var _ VoteSmart = &internal.Client{}

const defaultBaseURL = "https://api.votesmart.org/"

func New(apiKey string, options ...Option) (VoteSmart, error) {
	var opts clientOptions
	for _, opt := range options {
		opts = opt(opts)
	}

	if opts.baseURL == "" {
		opts.baseURL = defaultBaseURL
	}

	if opts.client == nil {
		opts.client = http.DefaultClient
	}

	u, err := url.Parse(opts.baseURL)
	if err != nil {
		return nil, err
	}

	return &internal.Client{
		HTTP:    opts.client,
		BaseURL: u,
		APIKey:  apiKey,
	}, nil
}

type clientOptions struct {
	client  *http.Client
	baseURL string
}

type Option func(clientOptions) clientOptions

func WithClient(v *http.Client) Option {
	return func(o clientOptions) clientOptions {
		o.client = v
		return o
	}
}

func WithBaseURL(v string) Option {
	return func(o clientOptions) clientOptions {
		o.baseURL = v
		return o
	}
}

type VoteSmart interface {
	Address
	BallotMeasure
	CandidateBio
	Candidates
	Committee
	District
	Election
	Leadership
	Local
	Npat
	Office
	Officials
	Rating
	State
	Votes
}

type Address interface {
	// AddressGetCampaign returns campaign office(s) and basic candidate
	// information for the specified candidate.
	// See http://api.votesmart.org/docs/Address.html for details.
	AddressGetCampaign(ctx context.Context, candidateID string) (any, error)
	// AddressGetCampaignWebAddress returns the campaign office's Web
	// address(es) and basic candidate information for the specified candidate.
	// See http://api.votesmart.org/docs/Address.html for details.
	AddressGetCampaignWebAddress(ctx context.Context, candidateID string) (any, error)
	// AddressGetCampaignByElection returns campaign office(s) and basic
	// candidate information for the specified election.
	// See http://api.votesmart.org/docs/Address.html for details.
	AddressGetCampaignByElection(ctx context.Context, electionID string) (any, error)
	// AddressGetOffice returns office(s) and basic candidate information
	// for the specified candidate.
	// See http://api.votesmart.org/docs/Address.html for details.
	AddressGetOffice(ctx context.Context, candidateID string) (any, error)
	// AddressGetOfficeWebAddress returns office's Web address(es) and basic
	// candidate information for the specified candidate.
	// See http://api.votesmart.org/docs/Address.html for details.
	AddressGetOfficeWebAddress(ctx context.Context, candidateID string) (any, error)
	// AddressGetOfficeByOfficeState returns office address and basic candidate
	// information according to the officeId and state.
	// Leave `stateID` blank to omit.
	// See http://api.votesmart.org/docs/Address.html for details.
	AddressGetOfficeByOfficeState(ctx context.Context, officeID string, stateID string) (any, error)
}

type BallotMeasure interface {
	// MeasureGetMeasuresByYearState returns a list of state ballot measures in
	// a given year.
	// See http://api.votesmart.org/docs/Measure.html for details.
	MeasureGetMeasuresByYearState(ctx context.Context, year string, stateID string) (*votesmarttypes.MeasureGetMeasuresByYearState, error)
	// MeasureGetMeasure returns a single Ballot Measure detail.
	// See http://api.votesmart.org/docs/Measure.html for details.
	MeasureGetMeasure(ctx context.Context, measureID string) (*votesmarttypes.MeasureGetMeasure, error)
}

type CandidateBio interface {
	// CandidateBioGetBio returns the main bio for each candidate.
	// See http://api.votesmart.org/docs/CandidateBio.html for details.
	CandidateBioGetBio(ctx context.Context, candidateID string) (any, error)
	// CandidateBioGetDetailedBio returns the education, profession, political,
	// orgMembership, and congMembership elements.
	// See http://api.votesmart.org/docs/CandidateBio.html for details.
	CandidateBioGetDetailedBio(ctx context.Context, candidateID string) (any, error)
	// CandidatedBioGetAdditionalBio returns the extended bio for each
	// candidate that has one.
	// See http://api.votesmart.org/docs/CandidateBio.html for details.
	CandidatedBioGetAdditionalBio(ctx context.Context, candidateID string) (any, error)
}

type Candidates interface {
	// CandidatesGetByOfficeState returns a list of candidates according to
	// office and state representation.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByOfficeState(ctx context.Context, officeID, stateID, electionYear, stageID string) (*votesmarttypes.CandidatesGetByOfficeState, error)
	// CandidatesGetByOfficeTypeState returns a list of candidates according to
	// office type and state representation.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByOfficeTypeState(ctx context.Context, officeTypeID, stateID, electionYear, stageID string) (*votesmarttypes.CandidatesGetByOfficeTypeState, error)
	// CandidatesGetByLastname returns a list of candidates according to a
	// lastname match.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByLastname(ctx context.Context, lastName, electionYear, stageID string) (*votesmarttypes.CandidatesGetByLastname, error)
	// CandidatesGetByLevenshtein returns a list of candidates according to a
	// fuzzy lastname match.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByLevenshtein(ctx context.Context, lastName, electionYear, stageID string) (*votesmarttypes.CandidatesGetByLastname, error)
	// CandidatesGetByElection returns a list of candidates according to the
	// election they are running in.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByElection(ctx context.Context, electionID, stageID string) (*votesmarttypes.CandidatesGetByElection, error)
	// CandidatesGetByDistrict returns a list of candidates according to the
	// district they represent.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByDistrict(ctx context.Context, districtID, electionYear, stageID string) (*votesmarttypes.CandidatesGetByDistrict, error)
	// CandidatesGetByZip returns actx returns according to the zip code they
	// represent.
	// See http://api.votesmart.org/docs/Candidates.html for details.
	CandidatesGetByZip(ctx context.Context, zip5, zip4, electionYear, stageID string) (*votesmarttypes.CandidatesGetByZip, error)
}

type Committee interface {
	// CommitteeGetTypes returns the committee types (house, senate, joint, etc).
	// See http://api.votesmart.org/docs/Committee.html for details.
	CommitteeGetTypes(ctx context.Context) (*votesmarttypes.CommitteeGetTypes, error)
	// CommitteeGetCommitteesByTypeState returns a list of committees that fit
	// the criteria.
	// See http://api.votesmart.org/docs/Committee.html for details.
	CommitteeGetCommitteesByTypeState(ctx context.Context, typeID, stateID string) (*votesmarttypes.CommitteeGetCommitteesByTypeState, error)
	// CommitteeGetCommittee returns detailed committee data.
	// See http://api.votesmart.org/docs/Committee.html for details.
	CommitteeGetCommittee(ctx context.Context, committeeID string) (*votesmarttypes.CommitteeGetCommittee, error)
	// CommitteeGetCommitteeMembers returns members of the committee.
	// See http://api.votesmart.org/docs/Committee.html for details.
	CommitteeGetCommitteeMembers(ctx context.Context, committeeID string) (*votesmarttypes.CommitteeGetCommitteeMembers, error)
}

type District interface {
	// DistrictGetByOfficeState returns district IDs according to the office
	// and state.
	// See http://api.votesmart.org/docs/District.html for details.
	DistrictGetByOfficeState(ctx context.Context, officeID, stateID, districtName string) (*votesmarttypes.DistrictGetByOfficeState, error)
	// DistrictGetByZip returns district IDs according to the zip code.
	// See http://api.votesmart.org/docs/District.html for details.
	DistrictGetByZip(ctx context.Context, zip5, zip4 string) (*votesmarttypes.DistrictGetByZip, error)
}

type Election interface {
	// ElectionGetElection returns district basic election data according to
	// electionId
	// See http://api.votesmart.org/docs/Election.html for details.
	ElectionGetElection(ctx context.Context, electionID string) (*votesmarttypes.ElectionGetElection, error)
	// ElectionGetElectionByYearState returns district basic election data
	// according to year and stateid.
	// See http://api.votesmart.org/docs/Election.html for details.
	ElectionGetElectionByYearState(ctx context.Context, year, stateID string) (*votesmarttypes.ElectionGetElectionByYearState, error)
	// ElectionGetElectionByZip returns district basic election data according
	// to zip code.
	// See http://api.votesmart.org/docs/Election.html for details.
	ElectionGetElectionByZip(ctx context.Context, zip5, zip4, year string) (*votesmarttypes.ElectionGetElectionByZip, error)
	// ElectionGetStageCandidates returns district basic election data
	// according to electionId and stageId. Per state lists of a Presidential
	// election are available by specifying the stateId.
	// See http://api.votesmart.org/docs/Election.html for details.
	ElectionGetStageCandidates(ctx context.Context, electionID, stageID, major, party, districtID, stateID string) (any, error)
}

type Leadership interface {
	// LeadershipGetPositions returns leadership positions by state and office.
	// See http://api.votesmart.org/docs/Local.html for details.
	LeadershipGetPositions(ctx context.Context, stateID, officeID string) (*votesmarttypes.LeadershipGetPositions, error)
	// LeadershipGetOfficials returns officials that hold the leadership role
	// in certain states.
	// See http://api.votesmart.org/docs/Local.html for details.
	LeadershipGetOfficials(ctx context.Context, leadershipID, stateID string) (*votesmarttypes.LeadershipGetOfficials, error)
}

type Local interface {
	// LocalGetCounties returns counties in a state
	LocalGetCounties(ctx context.Context, stateID string) (*votesmarttypes.LocalGetCounties, error)
	// LocalGetCities returns cities in a state
	LocalGetCities(ctx context.Context, stateID string) (*votesmarttypes.LocalGetCities, error)
	// LocalGetOfficials returns officials for a locality
	LocalGetOfficials(ctx context.Context, localID string) (*votesmarttypes.LocalGetOfficials, error)
}

type Npat interface {
	// NpatGetNpat returns the candidates most recently filled out NPAT/PCT.
	// See http://api.votesmart.org/docs/Npat.html for details.
	NpatGetNpat(ctx context.Context, candidateID string) (any, error)
}

type Office interface {
	// OfficeGetTypes returns all office types we keep track of.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetTypes(ctx context.Context) (*votesmarttypes.OfficeGetTypes, error)
	// OfficeGetBranches returns the branches of government and their IDs.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetBranches(ctx context.Context) (*votesmarttypes.OfficeGetBranches, error)
	// OfficeGetLevels returns the levels of government and their IDs.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetLevels(ctx context.Context) (*votesmarttypes.OfficeGetLevels, error)
	// OfficeGetOfficesByType returns offices we keep track of according to
	// type.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetOfficesByType(ctx context.Context, officeTypeID string) (*votesmarttypes.OfficeGetOfficesByType, error)
	// OfficeGetOfficesByLevel returns offices we keep track of according to
	// level.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetOfficesByLevel(ctx context.Context, levelID string) (*votesmarttypes.OfficeGetOfficesByLevel, error)
	// OfficeGetOfficesByTypeLevel returns offices we keep track of according
	// to type and level.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetOfficesByTypeLevel(ctx context.Context, officeTypeID, officeLevelID string) (*votesmarttypes.OfficeGetOfficesByTypeLevel, error)
	// OfficeGetOfficesByBranchLevel returns offices we keep track of
	// according to branch and level.
	// See http://api.votesmart.org/docs/Office.html for details.
	OfficeGetOfficesByBranchLevel(ctx context.Context, branchID, levelID string) (*votesmarttypes.OfficeGetOfficesByBranchLevel, error)
}

type Officials interface {
	// OfficialsGetStatewide returns a list of officials according to state
	// representation.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetStatewide(ctx context.Context, stateID string) (*votesmarttypes.OfficialsGetStatewide, error)
	// OfficialsGetByOfficeState returns a list of officials according to
	// office and state representation.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetByOfficeState(ctx context.Context, officeID, stateID string) (*votesmarttypes.OfficialsGetByOfficeState, error)
	// OfficialsGetByOfficeTypeState returns a list of officials according to
	// office type and state representation.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetByOfficeTypeState(ctx context.Context, officeTypeID, stateID string) (*votesmarttypes.OfficialsGetByOfficeTypeState, error)
	// OfficialsGetByLastname returns a list of officials according to a
	// lastName match.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetByLastname(ctx context.Context, lastName string) (*votesmarttypes.OfficialsGetByLastname, error)
	// OfficialsGetByLevenshtein returns a list of officials according to a
	// fuzzy lastName match.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetByLevenshtein(ctx context.Context, lastName string) (*votesmarttypes.OfficialsGetByLevenshtein, error)
	// OfficialsGetByDistrict returns a list of officials according to the
	// district they are running for.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetByDistrict(ctx context.Context, districtID string) (any, error)
	// OfficialsGetByZip returns a list of officials according to the zip code
	// they represent.
	// See http://api.votesmart.org/docs/Officials.html for details.
	OfficialsGetByZip(ctx context.Context, zip5, zip4 string) (*votesmarttypes.OfficialsGetByZip, error)
}

type Rating interface {
	// RatingGetCategories returns categories that contain released ratings
	// according to state.
	// See http://api.votesmart.org/docs/Rating.html for details.
	RatingGetCategories(ctx context.Context, stateID string) (any, error)
	// RatingGetSigList returns Special Interest Groups according to category
	// and state.
	// See http://api.votesmart.org/docs/Rating.html for details.
	RatingGetSigList(ctx context.Context, categoryID, stateID string) (any, error)
	// RatingGetSig returns detailed information an a Special Interest Group.
	// See http://api.votesmart.org/docs/Rating.html for details.
	RatingGetSig(ctx context.Context, sigID string) (any, error)
	// RatingGetSigRatings returns all ratings(scorecards) by a Special
	// Interest Group.
	// See http://api.votesmart.org/docs/Rating.html for details.
	RatingGetSigRatings(ctx context.Context, sigID string) (any, error)
	// RatingGetCandidateRating returns a candidate's rating by an SIG.
	// See http://api.votesmart.org/docs/Rating.html for details.
	RatingGetCandidateRating(ctx context.Context, candidateID, sigID string) (any, error)
	// RatingGetRating returns all candidate ratings from a scorecard by an SIG.
	// See http://api.votesmart.org/docs/Rating.html for details.
	RatingGetRating(ctx context.Context, ratingID string) (any, error)
}

type State interface {
	// StateGetStateIDs returns a simple state ID and name list for mapping
	// ID to state names.
	// See http://api.votesmart.org/docs/State.html for details.
	StateGetStateIDs(ctx context.Context) (*votesmarttypes.StateGetStateIDs, error)
	// StateGetState returns a various data we keep on a state
	// See http://api.votesmart.org/docs/State.html for details.
	StateGetState(ctx context.Context, stateID string) (*votesmarttypes.StateGetState, error)
}

type Votes interface {
	// VotesGetCategories returns categories that contain released bills
	// according to year and state.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetCategories(ctx context.Context, year, stateID string) (any, error)
	// VotesGetBill returns general information on a bill.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBill(ctx context.Context, billID string) (any, error)
	// VotesGetBillAction returns detailed action information on a certain
	// stage of the bill.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillAction(ctx context.Context, actionID string) (any, error)
	// VotesGetBillActionVotes provides votes listed by candidate on a
	// certain bill action.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillActionVotes(ctx context.Context, actionID string) (any, error)
	// VotesGetBillActionVoteByOfficial returns a single vote according to
	// official and action.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillActionVoteByOfficial(ctx context.Context, actionID, candidateID string) (any, error)
	// VotesGetByBillNumber returns a list of bills that are like the
	// billNumber input.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetByBillNumber(ctx context.Context, billNumber string) (any, error)
	// VotesGetBillsByCategoryYearState returns a list of bills that fit the
	// category, year, and state input.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillsByCategoryYearState(ctx context.Context, categoryID, yearID, stateID string) (any, error)
	// VotesGetBillByYearState returns a list of bills that fit the year and
	// state input.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillByYearState(ctx context.Context, year, stateID string) (any, error)
	// VotesGetBillsByOfficialYearOffice returns a list of bills that fit the
	// candidate and year.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillsByOfficialYearOffice(ctx context.Context, candidateID, year, officeID string) (any, error)
	// VotesGetBillsByOfficialCategoryOffice returns a list of bills that fit
	// the candidate and category.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillsByOfficialCategoryOffice(ctx context.Context, candidateID, categoryID, officeID string) (any, error)
	// VotesGetByOfficial returns all the bills an official has voted on based
	// on the candidateId, officeId, categoryId, and year
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetByOfficial(ctx context.Context, officeID, categoryID, year string) (any, error)
	// VotesGetBillsBySponsorYear returns a list of bills that fit the
	// sponsor's candidateId and year.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillsBySponsorYear(ctx context.Context, candidateID, year string) (any, error)
	// VotesGetBillsBySponsorCategory returns a list of bills that fit the
	// sponsor's candidateId and category.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillsBySponsorCategory(ctx context.Context, candidateID, categoryID string) (any, error)
	// VotesGetBillsByStateRecent returns a list of recent bills according to
	// the state. Max returned is 100 or however much less you want.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetBillsByStateRecent(ctx context.Context, amount, stateID string) (any, error)
	// VotesGetVetoes returns a list of vetoes according to candidate.
	// See http://api.votesmart.org/docs/Votes.html for details.
	VotesGetVetoes(ctx context.Context, candidateID string) (any, error)
}

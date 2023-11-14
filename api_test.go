package votesmart_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"dev.freespoke.com/go-votesmart"
	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func TestAPI(t *testing.T) {
	dummyKey := "apiKey"
	b, err := os.ReadFile("testdata/state_get_state_ids.json")
	if err != nil {
		t.Fatal(err)
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") != dummyKey {
			t.Fatal("missing or incorrect api key")
			return
		}

		if r.URL.Query().Get("o") != "JSON" {
			t.Fatal("missing output set to json")
			return
		}

		_, err = w.Write(b)
		if err != nil {
			t.Fatal(err)
		}
	}))

	c, err := votesmart.New(dummyKey, votesmart.WithBaseURL(svr.URL), votesmart.WithClient(svr.Client()))
	if err != nil {
		t.Fatal(err)
	}

	var a votesmarttypes.StateGetStateIDs
	if err := c.Invoke(context.Background(), nil, &a); err != nil {
		t.Fatal("err not nil", err)
	}

	if len(a.StateList.List.State) == 0 {
		t.Fatal("response was empty")
	}

	if a.StateList.List.State[0].StateID != "NA" {
		t.Fatal("unexpected value in first state")
	}
}

func TestWithParams(t *testing.T) {
	dummyKey := "apiKey"
	b, err := os.ReadFile("testdata/candidates_get_by_levenshtein.json")
	if err != nil {
		t.Fatal(err)
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") == "" {
			t.Fatal("missing api key")
			return
		}

		if r.URL.Query().Get("o") != "JSON" {
			t.Fatal("missing output set to json")
			return
		}

		if r.URL.Query().Get("lastName") != "Trump" {
			t.Fatal("missing or incorrect lastName param")
		}

		_, err = w.Write(b)
		if err != nil {
			t.Fatal(err)
		}
	}))

	c, err := votesmart.New(dummyKey, votesmart.WithBaseURL(svr.URL), votesmart.WithClient(svr.Client()))
	if err != nil {
		t.Fatal(err)
	}

	var a votesmarttypes.CandidatesGetByLevenshtein
	v := url.Values{}
	v.Add("lastName", "Trump")
	if err := c.Invoke(context.Background(), &v, &a); err != nil {
		t.Fatal("err not nil", err)
	}

	if len(a.CandidateList.Candidate) == 0 {
		t.Fatal("response was empty")
	}

	if a.CandidateList.Candidate[0].FirstName != "Donald" {
		t.Fatal("unexpected value in first candidate")
	}
}

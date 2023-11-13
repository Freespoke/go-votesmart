package votesmart_test

import (
	"context"
	"log"
	"net/url"
	"testing"

	"dev.freespoke.com/go-votesmart"
	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func TestAPI(t *testing.T) {
	c, err := votesmart.New("")
	if err != nil {
		t.Fatal(err)
	}

	var a votesmarttypes.AddressGetCampaign
	v := url.Values{}
	v.Add("candidateId", "15723")
	log.Println(c.Invoke(context.Background(), &v, &a))
	log.Println(a)
}

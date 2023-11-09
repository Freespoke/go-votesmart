package votesmart_test

import (
	"context"
	"log"
	"testing"

	"dev.freespoke.com/go-votesmart"
)

func TestAPI(t *testing.T) {
	c, err := votesmart.New("")
	if err != nil {
		t.Fatal(err)
	}

	res, err := c.RatingGetSig(context.Background(), "2709")
	log.Println(res, err)
}

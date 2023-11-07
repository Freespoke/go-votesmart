package votesmart_test

import (
	"context"
	"log"
	"testing"

	"dev.freespoke.com/go-votesmart"
)

func TestAPI(t *testing.T) {
	c, err := votesmart.New("api_key_here")
	if err != nil {
		t.Fatal(err)
	}

	res, err := c.StateGetState(context.Background(), "MI")
	log.Println(res, err)
}

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

	res, err := c.MeasureGetMeasure(context.Background(), "2528")
	log.Println(res, err)
}

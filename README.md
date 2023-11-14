# go-votesmart

A golang client for the [Votesmart](http://api.votesmart.org/docs/index.html) API.

[![Go Reference](https://pkg.go.dev/badge/dev.freespoke.com/go-votesmart.svg)](https://pkg.go.dev/dev.freespoke.com/go-votesmart)
[![Actions Status](https://github.com/freespoke/go-votesmart/workflows/test/badge.svg)](https://github.com/freespoke/go-votesmart/actions)

## Installation

```sh
$ go get dev.freespoke.com/go-votesmart
```

## Usage

Create an API client instance using your API Key.

```go
client, err := votesmart.New(os.Getenv("VOTESMART_API_KEY"))
if err != nil {
    t.Fatal(err)
}
```

Make requests by providing a VoteSmart response object, and a `*url.Values`
containing desired URL params:

```go
var resp votesmarttypes.CandidatesGetByLastname
vals := url.Values{}
vals.Add("lastName", "desantis")
vals.Add("electionYear", "2022")

if err := client.Invoke(context.Background(), &vals, &resp); err != nil {
    log.Fatal(err)
}

log.Println(resp)
```

Refer to the [API Documentation](http://api.votesmart.org/docs/index.html) for
required and optional parameters. Do not set `o` or `key`, as the library
handles these for you.

### License

MIT
package graphqlfuzz

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

// go-fuzz-build github.com/titanous/graphql-fuzz
// go-fuzz -bin ./graphqlfuzz-fuzz.zip -workdir workdir
func Fuzz(b []byte) int {
	r := graphql.Do(graphql.Params{
		Schema:        testutil.StarWarsSchema,
		RequestString: string(b),
	})
	if len(r.Errors) == 0 {
		return 1
	}
	return 0
}

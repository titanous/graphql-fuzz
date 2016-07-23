package graphqlfuzz

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

// To get a coverage profile: `go test -coverpkg github.com/graphql-go/graphql/... -coverprofile cover.out && go tool cover -html cover.out`
func TestCoverage(t *testing.T) {
	var valid int
	var invalid int
	filepath.Walk("workdir/corpus", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Error(err)
		}

		r := graphql.Do(graphql.Params{
			Schema:        testutil.StarWarsSchema,
			RequestString: string(data),
		})
		if len(r.Errors) == 0 {
			valid++
		} else {
			invalid++
		}
		return nil
	})
	fmt.Println(valid, "valid", invalid, "invalid")
}

package funcs

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestJsonnetFile(t *testing.T) {
	tests := []struct {
		Path cty.Value
		Opts cty.Value
		Err  string
	}{
		{
			cty.StringVal("testdata/jsonnet/hello.jsonnet"),
			cty.ObjectVal(nil),
			``,
		},

		{
			cty.StringVal("testdata/jsonnet/imports.jsonnet"),
			cty.ObjectVal(nil),
			``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Path.AsString(), func(t *testing.T) {
			require := require.New(t)

			abs, err := filepath.Abs(tt.Path.AsString())
			require.NoError(err)
			tt.Path = cty.StringVal(abs)

			got, err := JsonnetFileFunc.Call([]cty.Value{
				tt.Path,
				tt.Opts,
			})
			if tt.Err != "" {
				require.Error(err)
				require.Contains(err.Error(), tt.Err)
				return
			}
			require.NoError(err)

			// Ensure that our file ends in ".json"
			path := got.AsString()
			require.Equal(filepath.Ext(path), ".json")

			data, err := ioutil.ReadFile(path)
			require.NoError(err)

			const outSuffix = ".out"
			g := goldie.New(t,
				goldie.WithFixtureDir(filepath.Join("testdata", "jsonnet")),
				goldie.WithNameSuffix(outSuffix),
			)
			g.Assert(t, filepath.Base(tt.Path.AsString()), data)
		})
	}
}

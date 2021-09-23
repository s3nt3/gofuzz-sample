package offbyone_test

import (
	"testing"

	"github.com/s3nt3/gofuzz-sample/offbyone"
)

func FuzzOffByOne(f *testing.F) {
	f.Fuzz(func(t *testing.T, payload []byte) {
		t.Parallel()

		offbyone.OffByOne(payload)
	})
}

package tls_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/s3nt3/gofuzz-sample/tls"
)

func read(filename string) ([]byte, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	out, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func corpus(dir string, f *testing.F) {
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, seed := range d {
		if seed.IsDir() {
			corpus(path.Join(dir, seed.Name()), f)
		} else {
			s, err := read(path.Join(dir, seed.Name()))
			if err != nil {
				continue
			}

			f.Add(s)
		}
	}

}

func FuzzTLS(f *testing.F) {
	corpus("./corpus", f)

	f.Fuzz(func(t *testing.T, payload []byte) {
		t.Parallel()

		tls.TLS(payload)
	})
}

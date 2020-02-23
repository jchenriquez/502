package compareversions

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	V1     string `json:"v1"`
	V2     string `json:"v2"`
	Output int    `json:"output"`
}

func TestCompareVersion(t *testing.T) {
	f, err := os.Open("tests.json")

	if err != nil {
		t.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				t.Run(name, func(k *testing.T) {
					result := CompareVersion(test.V1, test.V2)

					if result != test.Output {
						k.Errorf("result is %d", result)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			t.Error(err)
			break
		}
	}
}

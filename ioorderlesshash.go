package orderlesshash

import (
	"bufio"
	"fmt"
	"io"
)

//IoUnorderedSha ...
func (hashes *Hashes) IoUnorderedSha(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)
	var tmp []string

	for scanner.Scan() {
		tmp = append(tmp, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return hashes.SliceUnorderedSha(tmp)
}

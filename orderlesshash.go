package orderlesshash

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"sort"
)

//Hashes ...
type Hashes []string

//GetUnorderedSha ...
func (hashes *Hashes) GetUnorderedSha() []byte {
	sort.Strings(*hashes)

	hash := md5.New()
	//resSha := sha1.New()
	for _, s := range *hashes {
		io.WriteString(hash, s)
	}

	return hash.Sum(nil)
	//return resSha.Sum(nil)
}

//SliceUnorderedSha ...
func (hashes *Hashes) SliceUnorderedSha(slice []string) ([]byte, error) {
	for _, s := range slice {
		hash := md5.New()
		//sha := sha1.New()
		io.WriteString(hash, s)
		*hashes = append(*hashes, string(hash.Sum(nil)))
	}
	return hashes.GetUnorderedSha(), nil
}

func (hashes *Hashes) anonymousUnorderedSha(anon map[string]interface{}) ([]byte, error) {
	for key, value := range anon {

		//sha := sha1.New()

		hash := md5.New()
		switch value.(type) {
		case string:
			v := value.(string)
			io.WriteString(hash, key+v)
			*hashes = append(*hashes, string(hash.Sum(nil)))
		case []interface{}:
			var tmp []string
			var tmpHashes Hashes
			for _, v := range value.([]interface{}) {
				switch v.(type) {
				case string:
					tmp = append(tmp, v.(string))
				case int, float64:
					tmp = append(tmp, fmt.Sprint(v))
				default:
					log.Printf("%T ====> %s", v, v)
					tmp = append(tmp, fmt.Sprint(v))
				}
			}
			h, _ := tmpHashes.SliceUnorderedSha(tmp)
			*hashes = append(*hashes, key+string(h))
		case map[string]interface{}:
			log.Printf("{} %T ====> %s", value, value)
			var tmpHashes Hashes
			h, _ := tmpHashes.anonymousUnorderedSha(value.(map[string]interface{}))
			*hashes = append(*hashes, key+string(h))
		}
	}
	return hashes.GetUnorderedSha(), nil
}

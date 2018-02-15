package thumb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Supported sizes
const (
	SMALL  imageSizeType = "small"
	MEDIUM imageSizeType = "medium"
	LARGE  imageSizeType = "large"
	MAX    imageSizeType = "max"
)

type imageSizeType string

func (s imageSizeType) GetProviderSizeName(provider string) string {
	if provider == "youtube" {
		switch s {
		case SMALL:
			return "default"
		case MEDIUM:
			return "mqdefault"
		case LARGE:
			return "sddefault"
		case MAX:
			return "maxresdefault"
		default:
			return "sddefault"
		}
	} else if provider == "vimeo" {
		switch s {
		case SMALL:
			return "thumbnail_small"
		case MEDIUM:
			return "thumbnail_medium"
		case LARGE:
			return "thumbnail_large"
		case MAX:
			return "thumbnail_large"
		default:
			return "thumbnail_large"
		}
	} else {
		return ""
	}

}

func Get(url string, size imageSizeType) string {
	if strings.Contains(url, "youtu.be") || strings.Contains(url, "youtube") {
		re := regexp.MustCompile(`(https?:\/\/)?(www.)?(youtube\.com\/watch\?v=|youtu\.be\/|youtube\.com\/watch\?[A-Za-z0-9_=-]+&v=)([A-Za-z0-9_-]*)(\&\S+)?(\?\S+)?`)
		return fmt.Sprintf("https://img.youtube.com/vi/%s/%s.jpg", re.FindStringSubmatch(url)[4], size.GetProviderSizeName("youtube"))
	} else if strings.Contains(url, "vimeo") {
		re := regexp.MustCompile(`^https?:\/\/(?:.*?)\.?(vimeo)\.com\/(\d+).*$`)
		match := re.FindStringSubmatch(url)
		if !(len(match) >= 3) {
			return ""
		}

		res, err := http.Get(fmt.Sprintf("http://vimeo.com/api/v2/video/%s.json", match[2]))
		if err != nil {
			return ""
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return ""
		}

		var a []interface{}
		err = json.Unmarshal(body, &a)
		if err != nil {
			return ""
		}

		if len(a) >= 1 {
			b := a[0].(map[string]interface{})
			return b[size.GetProviderSizeName("vimeo")].(string)
		}

		return ""

	} else {
		return ""
	}

}

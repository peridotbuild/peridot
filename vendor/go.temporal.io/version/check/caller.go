package check

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Caller interface {
	Call(r *VersionCheckRequest) (*VersionCheckResponse, error)
}

type callerImpl struct {
	scheme string
	host   string
}

func NewCaller() Caller {
	return &callerImpl{"https", "version-info.temporal.io"}
}

func (vc *callerImpl) Call(r *VersionCheckRequest) (*VersionCheckResponse, error) {
	err := validateRequest(r)
	if err != nil {
		return nil, err
	}
	u := vc.getUrl(r)
	tr := &http.Transport{
		DisableKeepAlives:   true,
		MaxIdleConnsPerHost: -1,
	}
	if vc.scheme == "https" {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	client := &http.Client{Transport: tr}
	reqBody, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("bad response code %v", resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	versionCheckResponse := &VersionCheckResponse{}
	err = json.Unmarshal(body, versionCheckResponse)
	if err != nil {
		return nil, err
	}
	err = validateResponse(versionCheckResponse)
	if err != nil {
		return nil, err
	}
	return versionCheckResponse, nil
}

func validateResponse(r *VersionCheckResponse) error {
	if len(r.Products) == 0 {
		return errors.New("invalid response: missing product list")
	}
	firstProduct := r.Products[0]
	if firstProduct.Product == "" || firstProduct.Current.Version == "" || firstProduct.Recommended.Version == "" {
		return errors.New("invalid response: missing product name, current or recommended version")
	}
	return nil
}
func validateRequest(r *VersionCheckRequest) error {
	if r.Product == "" || r.Version == "" || r.ClusterID == "" || r.DB == "" || r.OS == "" || r.Arch == "" || r.Timestamp == 0 {
		return errors.New("invalid request: missing required fields")
	}
	for _, info := range r.SDKInfo {
		if info.Name == "" || info.Version == "" {
			return errors.New("invalid request: missing required fields")
		}
	}
	return nil
}

func (vc *callerImpl) getUrl(r *VersionCheckRequest) *url.URL {
	var u url.URL
	u.Scheme = vc.scheme
	u.Host = vc.host
	u.Path = fmt.Sprintf("check")
	return &u
}

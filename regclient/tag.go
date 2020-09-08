package regclient

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/sirupsen/logrus"
)

func (rc *regClient) TagsList(ctx context.Context, ref Ref) (TagList, error) {
	tl := TagList{}
	host := rc.getHost(ref.Registry)
	repoURL := url.URL{
		Scheme: host.Scheme,
		Host:   host.DNS[0],
		Path:   "/v2/" + ref.Repository + "/tags/list",
	}

	rty := rc.getRetryable(host)
	resp, err := rty.DoRequest(ctx, "GET", repoURL)
	if err != nil {
		return tl, err
	}
	respBody, err := ioutil.ReadAll(resp)
	if err != nil {
		rc.log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Failed to read tag list")
		return tl, err
	}
	err = json.Unmarshal(respBody, &tl)
	if err != nil {
		rc.log.WithFields(logrus.Fields{
			"err":  err,
			"body": respBody,
		}).Warn("Failed to unmarshal tag list")
		return tl, err
	}

	return tl, nil
}
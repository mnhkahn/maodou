// Package analytics
// https://stackoverflow.com/questions/12837748/analytics-google-api-error-403-user-does-not-have-any-google-analytics-account
package analytics

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/mnhkahn/gogogo/logger"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
)

func Analytics(link string, conf *jwt.Config) (map[string]int, error) {
	// Initiate an http.Client, the following GET request will be
	// authorized and authenticated on the behalf of user@example.com.
	client := conf.Client(oauth2.NoContext)
	resp, err := client.Get(link)
	if err != nil {
		return nil, err
	}

	anaByts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := new(analyticsResult)
	err = json.Unmarshal(anaByts, res)
	if err != nil {
		return nil, err
	}

	analytics := make(map[string]int)

	for _, ana := range res.Rows {
		if len(ana) != 2 {
			continue
		}
		path := ana[0]
		pv, err := strconv.Atoi(ana[1])
		if err != nil {
			logger.Warn(err)
			continue
		}
		analytics[path] = pv
	}

	return analytics, nil
}

type analyticsResult struct {
	ColumnHeaders []struct {
		ColumnType string `json:"columnType"`
		DataType   string `json:"dataType"`
		Name       string `json:"name"`
	} `json:"columnHeaders"`
	ContainsSampledData bool   `json:"containsSampledData"`
	ID                  string `json:"id"`
	ItemsPerPage        int64  `json:"itemsPerPage"`
	Kind                string `json:"kind"`
	ProfileInfo         struct {
		AccountID             string `json:"accountId"`
		InternalWebPropertyID string `json:"internalWebPropertyId"`
		ProfileID             string `json:"profileId"`
		ProfileName           string `json:"profileName"`
		TableID               string `json:"tableId"`
		WebPropertyID         string `json:"webPropertyId"`
	} `json:"profileInfo"`
	Query struct {
		Dimensions  string   `json:"dimensions"`
		End_date    string   `json:"end-date"`
		Ids         string   `json:"ids"`
		Max_results int64    `json:"max-results"`
		Metrics     []string `json:"metrics"`
		Start_date  string   `json:"start-date"`
		Start_index int64    `json:"start-index"`
	} `json:"query"`
	Rows                [][]string `json:"rows"`
	SelfLink            string     `json:"selfLink"`
	TotalResults        int64      `json:"totalResults"`
	TotalsForAllResults struct {
		Ga_pageviews string `json:"ga:pageviews"`
	} `json:"totalsForAllResults"`
}

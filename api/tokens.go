package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/xchange/core"
	"io/ioutil"
	"net/http"
)

func FetchTokens(web *http.Client, host string) (records.Page, error) {
	url := fmt.Sprintf("%s/tokens", host)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.Currency{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

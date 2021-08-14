package generalApiReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func JsonRequest(req *http.Request, result interface{}) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("[" + strconv.Itoa(res.StatusCode) + ":] " + req.URL.String())

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("respond with %d http code", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}

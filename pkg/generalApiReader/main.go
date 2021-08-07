package generalApiReader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//TODO: server log with requests and code of response

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	//fmt.Println(string(body))

	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}

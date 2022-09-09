package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func CheckDebeziumConnector() error {
	response, err := http.Get("http://localhost:8083/connectors/default_connector")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		plan, _ := ioutil.ReadFile("./debezium-plan.json")
		_, err := http.Post("http://localhost:8083/connectors/", "application/json", bytes.NewBuffer(plan))

		if err != nil {
			panic(err)
		}

		return nil
	}

	return err
}

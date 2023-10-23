package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var data entities.RecordData
		decodeErr := decoder.Decode(&data)
		if decodeErr != nil {
			panic(decodeErr)
		}

		file, err := filepath.Abs("data/pre_approved_phone_no.csv")
		if err != nil {
			log.Fatal(err)
		}

		var result = engine.DecideEngineRules(data, file)
		jsonResp := entities.JSONResponse{
			Status: result,
		}

		encodeErr := json.NewEncoder(resp).Encode(jsonResp)
		if encodeErr != nil {
			http.Error(resp, encodeErr.Error(), 500)
			return
		}

	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}

}

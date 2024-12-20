package application

import (
	"encoding/json"
	"net/http"
	"os"

	calc "github.com/StepanShel/YaCalc/pkg/calculation"
)

type Request struct {
	Expression string `json:"expression"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseRes struct {
	Result float64 `json:"result"`
}

func respJson(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-type", "application/json")
	switch data := data.(type) {
	case string:
		resp := ResponseError{Error: data}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			return err
		}
	case float64:
		resp := ResponseRes{Result: data}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respJson(w, calc.InvalidReqBody)
		return
	}
	defer r.Body.Close()
	res, err := calc.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respJson(w, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		respJson(w, res)
		return
	}

}

func StartServer() {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	http.ListenAndServe(":"+GetAddr(), nil)
}

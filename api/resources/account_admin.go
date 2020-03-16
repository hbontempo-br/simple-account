package resources

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"simple-account/api/controllers"
	"simple-account/middleware/db_middleware"
	"strconv"
)


type accountRequest struct {
DocumentNumber string `json:"document_number"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request){
tx := db_middleware.RetrieveContext(r.Context())

var acc accountRequest

err := json.NewDecoder(r.Body).Decode(&acc)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

account, err := controllers.CreateAccount(tx, acc.DocumentNumber)
if err != nil {
_ = err
}
json.NewEncoder(w).Encode(account)
tx.Commit()
}

func GetAccountList(w http.ResponseWriter, r *http.Request) {
	tx := db_middleware.RetrieveContext(r.Context())

	accountList := controllers.GetAccountList(tx)

	json.NewEncoder(w).Encode(&accountList)
}

func GetSingleAccount(w http.ResponseWriter, r *http.Request) {
	tx := db_middleware.RetrieveContext(r.Context())

	vars := mux.Vars(r)

	idString, found := vars["id"]
	if !found{
		w.WriteHeader(http.StatusBadRequest)
		str := `{"description":"Missing id (GET /account/{id})"}`
		fmt.Fprint(w, str)
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		str := `{"description":"Unable to convert provided id ot integer (GET /account/{id})"}`
		fmt.Fprint(w, str)
	}

	// Validate errors
	account := controllers.GetAccount(tx, id, "")

	json.NewEncoder(w).Encode(&account)



	//account := controllers.GetAccount(tx, vars["id"], )
	print(vars)
}
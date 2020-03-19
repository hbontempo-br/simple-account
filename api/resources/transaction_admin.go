package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-account/api/controllers"
	"simple-account/middleware/db_middleware"
)

type transactionRequest struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	tx := db_middleware.RetrieveContext(r.Context())

	var trans transactionRequest

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	operationType := controllers.GetOperationType(tx, trans.OperationTypeId)
	if operationType.Behavior == "negative" {
		trans.Amount = -trans.Amount
	}

	err = controllers.UpdateBalance(tx, trans.AccountId, trans.Amount)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		str := `{"description":"Insufficient funds"}`
		fmt.Fprint(w, str)
		return
	}


	transaction, err := controllers.CreateTransaction(tx, trans.AccountId, trans.OperationTypeId, trans.Amount)
	if err != nil {
		_ = err
	}
	json.NewEncoder(w).Encode(transaction)
	tx.Commit()
}

//func GetAccountList(w http.ResponseWriter, r *http.Request) {
//	tx := db_middleware.RetrieveContext(r.Context())
//
//	parameters := r.URL.Query()
//	documentNumbers, ok := parameters["documentNumbers"]
//
//	var ids []int
//	idsString, ok := parameters["id"]
//	for _, idString := range(idsString) {
//		id , err := strconv.Atoi(idString)
//		if err != nil {
//			print(err)
//		}
//		ids = append(ids, id)
//	}
//	print(ok)
//	print(documentNumbers)
//	print(parameters)
//
//	accountList := controllers.GetAccountList(tx, ids, documentNumbers)
//
//	json.NewEncoder(w).Encode(&accountList)
//}

func GetSingleTransaction(w http.ResponseWriter, r *http.Request) {
	tx := db_middleware.RetrieveContext(r.Context())

	//vars := mux.Vars(r)
	//
	//idString, found := vars["id"]
	//if !found{
	//	w.WriteHeader(http.StatusBadRequest)
	//	str := `{"description":"Missing id (GET /account/{id})"}`
	//	fmt.Fprint(w, str)
	//}
	//id, err := strconv.Atoi(idString)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	str := `{"description":"Unable to convert provided id ot integer (GET /account/{id})"}`
	//	fmt.Fprint(w, str)
	//}

	// Validate errors
	transaction := controllers.GetTransaction(tx, 1)

	json.NewEncoder(w).Encode(&transaction)

}

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

func GetOperationTypeList(w http.ResponseWriter, r *http.Request) {
	tx := db_middleware.RetrieveContext(r.Context())

	parameters := r.URL.Query()
	operationTypes, ok := parameters["operation_type"]

	var ids []int
	idsString, ok := parameters["id"]
	for _, idString := range idsString {
		id, err := strconv.Atoi(idString)
		if err != nil {
			print(err)
		}
		ids = append(ids, id)
	}

	print(ok)
	print(operationTypes)
	print(parameters)

	operationTypeList := controllers.GetOperationTypeList(tx, ids, operationTypes)

	json.NewEncoder(w).Encode(&operationTypeList)
}

func GetSingleOperationType(w http.ResponseWriter, r *http.Request) {
	tx := db_middleware.RetrieveContext(r.Context())

	vars := mux.Vars(r)

	idString, found := vars["id"]
	if !found {
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
	operationType := controllers.GetOperationType(tx, id)

	json.NewEncoder(w).Encode(&operationType)

}

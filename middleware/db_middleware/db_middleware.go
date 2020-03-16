package db_middleware

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"simple-account/utils/mysql_connector"
)

const transactionKey = "transaction"

func DBMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Creating DB transaction")
		db := mysql_connector.GetDB()
		tx := db.Begin()
		ctx := r.Context()
		//ctx = context.WithValue(ctx, "transaction", tx)
		NewContext(&ctx, tx)
		h.ServeHTTP(w, r.WithContext(ctx))
		if tx != nil {
			tx.Rollback()
		}
		//db.Close()
	})
}

func NewContext(ctx *context.Context, tx *gorm.DB){
	*ctx = context.WithValue(*ctx, transactionKey, tx)
}

func RetrieveContext(ctx context.Context) *gorm.DB{
	txInterface := ctx.Value(transactionKey)
	tx := txInterface.(*gorm.DB)
	return tx
}
package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/91diego/bank-transactions/models"
	"github.com/91diego/bank-transactions/repositories"
	"github.com/91diego/bank-transactions/server"
	"github.com/91diego/bank-transactions/utils"
	"github.com/google/uuid"
)

func InsertTransactions(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer file.Close()

		csvRecords, err := utils.ReadCSV(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		balanceID := uuid.New()
		userID := uuid.New()

		var totalDebitTransactions float64 = 0
		var totalCreditTransactions float64 = 0
		var trxType string
		var numberTrxDebit int = 0
		var numberTrxCredit int = 0

		var csvTransactions models.Transaction
		for _, record := range csvRecords {
			csvTransactions.ID = record[0]
			csvTransactions.TransactionDate = record[1]
			csvTransactions.TransactionAmount = record[2]
			trxType = GetTransactionInfo(&numberTrxDebit, &numberTrxCredit, &totalDebitTransactions, &totalCreditTransactions, csvTransactions.TransactionAmount)
			csvTransactions.TransactionType = trxType
			csvTransactions.BlanceID = balanceID.String()
			err := repositories.InsertTransactions(context.Background(), &csvTransactions)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// upsert balance
		reqBbalance := &models.Balance{
			ID:            balanceID.String(),
			Total:         totalDebitTransactions + totalCreditTransactions,
			DebitAvarage:  totalDebitTransactions / float64(numberTrxCredit),
			CreditAvarage: totalCreditTransactions / float64(numberTrxCredit),
			UserID:        userID.String(),
		}
		err = repositories.InsertBalance(context.Background(), reqBbalance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		balance, err := repositories.GetBalanceByID(context.Background(), balanceID.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if balance.ID == "" {
			balance = reqBbalance
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.TransactionResponse{
			Message:        "Transactions saved successfully",
			SummaryBalance: balance,
		})
	}
}

// GetTransactionInfo returns type of transaction, number of transactions for each transaction (debit or credit)
// and total number of transactions
func GetTransactionInfo(numberTrxDebit, numberTrxCredit *int, totalDebitTransactions, totalCreditTransactions *float64, transaction string) (transactionType string) {

	if strings.Contains(transaction, "+") {
		amount := strings.Split(transaction, "+")
		floatAmount, _ := strconv.ParseFloat(amount[1], 64)
		*totalDebitTransactions = *totalDebitTransactions + floatAmount
		*numberTrxDebit++
		transactionType = "CREDIT"
	}
	if strings.Contains(transaction, "-") {
		amount := strings.Split(transaction, "-")
		floatAmount, _ := strconv.ParseFloat(amount[1], 64)
		*totalCreditTransactions = *totalCreditTransactions - floatAmount
		*numberTrxCredit++
		transactionType = "DEBIT"
	}
	return
}

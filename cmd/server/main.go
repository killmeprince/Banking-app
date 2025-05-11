package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/robfig/cron/v3"
	"banking-app/internal/handler"
	"banking-app/internal/middleware"
	"banking-app/internal/repository"
	"banking-app/internal/service"
	"banking-app/pkg/database"
	"banking-app/pkg/jwt"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	ur := repository.NewUserRepo(db)
	ar := repository.NewAccountRepo(db)
	cr := repository.NewCardRepo(db)
	tr := repository.NewTransactionRepo(db)
	crd := repository.NewCreditRepo(db)
	psr := repository.NewPSRepo(db)

	us := service.NewUserService(ur)
	as := service.NewAccountService(ar, tr)
	cs := service.NewCardService(cr)
	ts := service.NewTransactionService(tr)
	cds := service.NewCreditService(crd, psr)
	an := service.NewAnalyticsService(tr, crd, ar)

	uh := handler.NewUserHandler(us)
	ah := handler.NewAccountHandler(as)
	ch := handler.NewCardHandler(cs)
	th := handler.NewTransactionHandler(ts)
	cdh := handler.NewCreditHandler(cds)
	anh := handler.NewAnalyticsHandler(an)

	// cronJob := cron.New()
	// cronJob.AddFunc("@every 12h", func() {
	// 	// TODO: DebitScheduledPayments()
	// 	log.Println("Scheduler tick")
	// })
	// cronJob.Start()

	r := mux.NewRouter()
	r.Use(middleware.Logger)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/register", uh.Register).Methods("POST")
	api.HandleFunc("/login", uh.Login).Methods("POST")

	sec := api.NewRoute().Subrouter()
	sec.Use(jwt.AuthMiddleware)
	sec.HandleFunc("/accounts", ah.Create).Methods("POST")
	sec.HandleFunc("/accounts", ah.List).Methods("GET")
	sec.HandleFunc("/accounts/deposit", ah.Deposit).Methods("POST")
	sec.HandleFunc("/transfer", ah.Transfer).Methods("POST")

	sec.HandleFunc("/cards", ch.Create).Methods("POST")
	sec.HandleFunc("/transactions", th.List).Methods("GET")
	sec.HandleFunc("/credits", cdh.Create).Methods("POST")
	sec.HandleFunc("/credits/{id}/schedule", cdh.Schedule).Methods("GET")

	sec.HandleFunc("/analytics/month", anh.MonthStats).Methods("GET")
	sec.HandleFunc("/analytics/credit", anh.CreditLoad).Methods("GET")
	sec.HandleFunc("/analytics/predict", anh.Predict).Methods("GET")
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

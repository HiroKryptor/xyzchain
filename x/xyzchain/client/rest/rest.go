package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    // this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers xyzchain-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("/xyzchain/pools/{id}", getPoolHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/xyzchain/pools", listPoolHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/xyzchain/pools", createPoolHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/xyzchain/pools/{id}", updatePoolHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/xyzchain/pools/{id}", deletePoolHandler(clientCtx)).Methods("POST")

}


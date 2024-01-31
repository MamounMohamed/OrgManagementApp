package routes

import (
	handlers "orgmanagementapp/pkg/api/handlers"
	"orgmanagementapp/pkg/api/middleware"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/").Subrouter()

	userRouter.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	userRouter.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
	userRouter.HandleFunc("/refresh-token", handlers.RefreshTokenHandler).Methods("POST")
}

func SetOrganizationRoutes(router *mux.Router) {
	orgRouter := router.PathPrefix("/organization").Subrouter()
	orgRouter.Use(middleware.AuthMiddleware)
	orgRouter.HandleFunc("", handlers.CreateOrganizationHandler).Methods("POST", "OPTIONS")
	orgRouter.HandleFunc("/{organization_id}", handlers.ReadOrganizationHandler).Methods("GET", "OPTIONS")
	orgRouter.HandleFunc("", handlers.ReadAllOrganizationsHandler).Methods("GET", "OPTIONS")
	orgRouter.HandleFunc("/{organization_id}", handlers.UpdateOrganizationHandler).Methods("PUT", "OPTIONS")
	orgRouter.HandleFunc("/{organization_id}", handlers.DeleteOrganizationHandler).Methods("DELETE", "OPTIONS")
	orgRouter.HandleFunc("/{organization_id}/invite", handlers.InviteUserToOrganizationHandler).Methods("POST", "OPTIONS")
}

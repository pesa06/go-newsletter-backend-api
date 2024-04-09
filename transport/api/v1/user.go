package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"
	userStructs "newsletter_backend_api/transport/api/v1/structs/user"
)

func (h *Handler) ListUsers(writer http.ResponseWriter, request *http.Request) {
	slog.Info("list users triggered")
}

func (h *Handler) GetUser(writer http.ResponseWriter, request *http.Request) {
	slog.Info("get user triggered")
}

func (h *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	slog.Info("create user triggered")
	slog.Info("request", request.Body)
	var user userStructs.CreateUserStruct
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		slog.Info("error in decoding")
		slog.Error(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	slog.Info("user", user)
	err = h.service.CreateUser(request.Context(), user)
	if err != nil {
		slog.Error(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

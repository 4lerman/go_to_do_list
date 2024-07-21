package to_do

import (
	"fmt"
	"net/http"

	"github.com/4lerman/go_to_do_list/types"
	"github.com/4lerman/go_to_do_list/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Handler is the HTTP handler for to-do operations
type Handler struct {
	store types.ToDoStore
}

// NewHandler creates a new to-do handler
func NewHandler(store types.ToDoStore) *Handler {
	return &Handler{
		store: store,
	}
}

// RegisterRoutes registers to-do routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todo-list/tasks", h.handleGetToDos).Methods(http.MethodGet)
	router.HandleFunc("/todo-list/tasks", h.handleCreateToDo).Methods(http.MethodPost)
	router.HandleFunc("/todo-list/tasks/{id}", h.handleDeleteToDo).Methods(http.MethodDelete)
	router.HandleFunc("/todo-list/tasks/{id}", h.handleUpdateToDo).Methods(http.MethodPut)
	router.HandleFunc("/todo-list/tasks/{id}/done", h.handleChangeToDone).Methods(http.MethodPut)
}

// handleCreateToDo creates a new to-do task
// @Summary Create a new to-do task
// @Description Create a new to-do task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body types.CreateToDoPayload true "ToDo Task"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} string
// @Router /todo-list/tasks [post]
func (h *Handler) handleCreateToDo(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateToDoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	id, err := h.store.CreateToDo(types.ToDo{
		Title:    payload.Title,
		ActiveAt: payload.ActiveAt,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"todo_id": id,
	})
}

// handleGetToDos retrieves to-do tasks
// @Summary Get to-do tasks
// @Description Get to-do tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Param status query string false "Task status (active or done)"
// @Success 200 {array} types.ToDo
// @Failure 400 {object} string
// @Router /todo-list/tasks [get]
func (h *Handler) handleGetToDos(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	status := queryParams.Get("status")
	if status == "" {
		status = "active"
	}

	toDos, err := h.store.GetToDos(status)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, toDos)
}

// handleDeleteToDo deletes a to-do task
// @Summary Delete a to-do task
// @Description Delete a to-do task
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /todo-list/tasks/{id} [delete]
func (h *Handler) handleDeleteToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	if taskId == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	if err := h.store.DeleteToDo(taskId); err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

// handleUpdateToDo updates a to-do task
// @Summary Update a to-do task
// @Description Update a to-do task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body types.UpdateToDoPayload true "Updated Task"
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /todo-list/tasks/{id} [put]
func (h *Handler) handleUpdateToDo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	if taskId == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	var payload types.UpdateToDoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	if err := h.store.UpdateToDo(taskId, types.ToDo{
		Title:    payload.Title,
		ActiveAt: payload.ActiveAt,
	}); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

// handleChangeToDone marks a to-do task as done
// @Summary Mark a to-do task as done
// @Description Mark a to-do task as done
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /todo-list/tasks/{id}/done [put]
func (h *Handler) handleChangeToDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	if taskId == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("id is not indicated"))
		return
	}

	if err := h.store.ChangeToDone(taskId); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

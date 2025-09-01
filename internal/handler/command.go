package handler

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

// ExecuteCommand godoc
// @Summary Execute a command in BusyBox
// @Description Execute a shell command within the BusyBox container and return its output
// @Tags commands
// @Accept json
// @Produce json
// @Param command body CommandRequest true "Command to execute"
// @Success 200 {object} CommandResponse
// @Failure 400 {object} CommandResponse
// @Failure 500 {object} CommandResponse
// @Router /execute [post]
func ExecuteCommand(w http.ResponseWriter, r *http.Request) {
	var cmd CommandRequest
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	output, err := exec.Command("sh", "-c", cmd.Command).CombinedOutput()
	response := CommandResponse{
		Output: string(output),
	}
	
	if err != nil {
		response.Error = err.Error()
		respondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

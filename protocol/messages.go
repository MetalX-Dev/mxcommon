package protocol

import "github.com/bytedance/sonic"

type CommandExecutionRequest struct {
	Command string `json:"command"`
}

type CommandExecutionResponse struct {
	Code   int32  `json:"code"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type FileOperation string

const (
	Download FileOperation = "Download"
	Upload   FileOperation = "Upload"
)

type FileOperationRequest struct {
	URL       string        `json:"url"`
	Path      string        `json:"path"`
	Operation FileOperation `json:"operation"`
}

type FileOperationResponse struct {
	Success bool `json:"success"`
}

type ControllerRequestPayload struct {
	// Type                    string                   `json:"type"`
	CommandExecutionRequest *CommandExecutionRequest `json:",omitempty"`
	FileOperationRequest    *FileOperationRequest    `json:",omitempty"`
}

type ControllerRequest struct {
	Version uint32                   `json:"version"`
	ID      uint64                   `json:"id"`
	Payload ControllerRequestPayload `json:"payload"`
}

type AgentResponsePayload struct {
	// Type                     string                    `json:"type"`
	CommandExecutionResponse *CommandExecutionResponse `json:",omitempty"`
	FileOperationResponse    *FileOperationResponse    `json:",omitempty"`
}

type AgentResponse struct {
	ID      uint64               `json:"id"`
	OK      bool                 `json:"ok"`
	Payload AgentResponsePayload `json:"payload"`
}

func ParseAgentResponse(s string) (*AgentResponse, error) {
	var resp AgentResponse
	err := sonic.Unmarshal([]byte(s), &resp)
	return &resp, err
}

func (r *ControllerRequest) String() string {
	b, _ := sonic.Marshal(r)
	return string(b)
}

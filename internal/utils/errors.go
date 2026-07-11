package utils

import "fmt"

type ErrType int

const (
	ErrProjectNotFound    ErrType = iota + 1
	ErrLanguageUnsupported
	ErrToolNotFound
	ErrAnalysisFailed
	ErrConfigInvalid
	ErrDriverFailed
	ErrTimeout
)

type KittyError struct {
	Type    ErrType
	Message string
	Wrapped error
}

func (e *KittyError) Error() string {
	if e.Wrapped != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
	}
	return e.Message
}

func (e *KittyError) Unwrap() error { return e.Wrapped }

func NewProjectNotFound(path string) *KittyError {
	return &KittyError{Type: ErrProjectNotFound, Message: fmt.Sprintf("project not found: %s", path)}
}

func NewLanguageUnsupported(lang string) *KittyError {
	return &KittyError{Type: ErrLanguageUnsupported, Message: fmt.Sprintf("unsupported language: %s", lang)}
}

func NewToolNotFound(tool string) *KittyError {
	return &KittyError{Type: ErrToolNotFound, Message: fmt.Sprintf("tool not found: %s", tool)}
}

func NewAnalysisFailed(cause error) *KittyError {
	return &KittyError{Type: ErrAnalysisFailed, Message: "analysis failed", Wrapped: cause}
}

func NewConfigInvalid(detail string) *KittyError {
	return &KittyError{Type: ErrConfigInvalid, Message: fmt.Sprintf("invalid config: %s", detail)}
}

func NewDriverFailed(name string, cause error) *KittyError {
	return &KittyError{Type: ErrDriverFailed, Message: fmt.Sprintf("driver %s failed", name), Wrapped: cause}
}

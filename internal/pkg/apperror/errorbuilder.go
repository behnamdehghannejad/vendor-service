package apperror

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/log"
)

type BuilderError struct {
	args      []any
	pattern   string
	err       error
	isPrinted bool
	input     []any
	errorType ErrorType
	warning   bool
	debugging bool
}

func Wrap(err error) *BuilderError {
	return &BuilderError{
		isPrinted: true,
		args:      []any{},
		pattern:   "",
		err:       err,
		input:     []any{},
	}
}

func WithoutParentError() *BuilderError {
	return &BuilderError{
		isPrinted: true,
		args:      []any{},
		pattern:   "",
		err:       nil,
		input:     []any{},
	}
}

func (m *BuilderError) getErrorType() ErrorType {
	if m.errorType != 0 {
		return m.errorType
	}

	var appErr *AppError

	if errors.As(m.err, &appErr) {
		return appErr.GetErrorType()
	}

	return UnExpected
}

func (m *BuilderError) Message() string {
	message := m.matchPatternAndArgs()
	if len(message) > 0 {
		return message
	}

	if m.err != nil {
		return m.err.Error()
	}

	return ""
}

func (m *BuilderError) UnExpected() *BuilderError {
	m.errorType = UnExpected
	return m
}

func (m *BuilderError) NotFound() *BuilderError {
	m.errorType = NotFound
	return m
}

func (m *BuilderError) BadRequest() *BuilderError {
	m.errorType = BadRequest
	return m
}

func (m *BuilderError) Forbidden() *BuilderError {
	m.errorType = Forbidden
	return m
}

func (m *BuilderError) Duplicate() *BuilderError {
	m.errorType = Duplicate
	return m
}

func (m *BuilderError) DeactiveWrite() *BuilderError {
	m.isPrinted = false
	return m
}

func (m *BuilderError) ActiveWrite() *BuilderError {
	m.isPrinted = true
	return m
}

func (m *BuilderError) ErrorMessage() string {
	message := fmt.Sprintf(`the main error is "%s"`, m.mainError())

	messageInput := m.getInputMessage()

	if len(messageInput) > 0 {
		message += fmt.Sprintf(` also we got ("%s")`, messageInput)
	}

	additionalMessage := m.matchPatternAndArgs()

	if len(additionalMessage) > 0 {
		message += fmt.Sprintf(` the additional information is "%s"`, additionalMessage)
	}

	message += "\n\nstack trace:\n" + string(debug.Stack())

	return message
}

func (m *BuilderError) matchPatternAndArgs() string {
	if len(m.pattern) > 0 && len(m.args) > 0 {
		return fmt.Sprintf(m.pattern, m.args...)
	}

	if len(m.pattern) > 0 {
		return m.pattern
	}

	return ""
}

func (m *BuilderError) Input(data ...any) *BuilderError {
	m.input = data
	return m
}

func (m *BuilderError) mainError() string {
	if m.err == nil {
		return "nothing"
	}

	err := m.err

	for {
		unwrapped := errors.Unwrap(err)

		if unwrapped == nil {
			break
		}

		err = unwrapped
	}

	return err.Error()
}

func (m *BuilderError) getInputMessage() string {
	messages := make([]string, 0, len(m.input))

	for _, item := range m.input {
		messages = append(messages, m.translateInput(item))
	}

	return strings.Join(messages, `", "`)
}

func (m *BuilderError) translateInput(input any) string {
	return fmt.Sprintf("%v", input)
}

func (m *BuilderError) Warningf(pattern string, args ...any) *BuilderError {
	m.args = args
	m.pattern = pattern

	m.warning = true
	m.debugging = false

	return m
}

func (m *BuilderError) DebuggingErrorf(pattern string, args ...any) *BuilderError {
	m.args = args
	m.pattern = pattern

	m.warning = false
	m.debugging = true

	return m
}

func (m *BuilderError) DebuggingError() *BuilderError {
	m.warning = false
	m.debugging = true

	return m
}

func (m *BuilderError) Warning() *BuilderError {
	m.warning = true
	m.debugging = false

	return m
}

func (m *BuilderError) Log() *BuilderError {
	if m.warning {
		log.Warning(m.ErrorMessage())
	} else if m.debugging {
		log.Debug(m.ErrorMessage())
	}

	return m
}

func (m *BuilderError) Build() *AppError {
	return NewAppError(m.Message(), m.getErrorType())
}

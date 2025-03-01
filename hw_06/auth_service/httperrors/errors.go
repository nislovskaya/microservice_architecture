package httperrors

import "fmt"

type BadRequestError struct {
	Message string
	Err     error
}

func (e *BadRequestError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s, error: %s", e.Message, e.Err.Error())
	}
	return e.Message
}

type UnauthorizedError struct {
	Message string
	Err     error
}

func (e *UnauthorizedError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s, error: %s", e.Message, e.Err.Error())
	}
	return e.Message
}

type NotFoundError struct {
	Message string
	Err     error
}

func (e *NotFoundError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s, error: %s", e.Message, e.Err.Error())
	}
	return e.Message
}

type ConflictError struct {
	Message string
	Err     error
}

func (e *ConflictError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s, error: %s", e.Message, e.Err.Error())
	}
	return e.Message
}

type InternalServerError struct {
	Message string
	Err     error
}

func (e *InternalServerError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s, error: %s", e.Message, e.Err.Error())
	}
	return e.Message
}

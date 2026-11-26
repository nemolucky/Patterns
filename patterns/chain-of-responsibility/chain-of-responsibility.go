package chainofresponsibility

import (
	"fmt"
	"strings"
)

var (
	ErrInvalidUsername = fmt.Errorf("invalid username")
	ErrInvalidRole     = fmt.Errorf("invalid role")
	ErrNotFoundInStock = fmt.Errorf("failed not found in stock")
)

type Handler interface {
	Handle(req *Request) error
}

type Request struct {
	Username string
	Role     string
	InStock  bool
}

func NewRequest(user, role string, in bool) *Request {
	return &Request{Username: user, Role: role, InStock: in}
}

type AuthHandler struct {
	next Handler
}

func NewAuthHandler(next Handler) *AuthHandler {
	return &AuthHandler{next: next}
}

func (a *AuthHandler) Handle(req *Request) error {
	if strings.TrimSpace(req.Username) == "" {
		return ErrInvalidUsername
	}

	if a.next != nil {
		return a.next.Handle(req)
	}

	return nil
}

type RoleHandler struct {
	next Handler
}

func NewRoleHandler(next Handler) *RoleHandler {
	return &RoleHandler{next: next}
}

func (a *RoleHandler) Handle(req *Request) error {
	if req.Role != "customer" {
		return ErrInvalidRole
	}

	if a.next != nil {
		return a.next.Handle(req)
	}

	return nil
}

type InStockHandler struct {
	next Handler
}

func NewInStockHandler(next Handler) *InStockHandler {
	return &InStockHandler{next: next}
}

func (a *InStockHandler) Handle(req *Request) error {
	if !req.InStock {
		return ErrNotFoundInStock
	}

	if a.next != nil {
		return a.next.Handle(req)
	}

	return nil
}

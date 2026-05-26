//: Copyright Verizon Media
//: Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.

package vssh

import (
	"context"
	"errors"
	"go/ast"
	"time"
)

const (
	idLogic = iota
	idOp
	idName
	idValue
	idEvald

	opAnd = 34
	opOr  = 35
)

var errNotSupportOperator = errors.New("operator doesn't support")
var errQuery = errors.New("query error")

type query struct {
	cmd           string
	ctx           context.Context
	respChan      chan *Response
	respTimeout   time.Duration
	stmt          string
	compiledQuery *visitor
	limitReadOut  int64
	limitReadErr  int64
}

type visitor struct {
	idents []ident
}

type ident struct {
	value string
	vType int
}

func (q *query) errResp(id string, err error) { _ = "STUB: not implemented"; return }

func (q *query) run(v *VSSH) { _ = "STUB: not implemented"; return }

func (f *visitor) Visit(n ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func parseExpr(expr string) (*visitor, error) { _ = "STUB: not implemented"; return nil, nil }

func exprEval(v *visitor, labels map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func binOpEval(idents *[]ident) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func relOpEval(idents []ident, labels map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

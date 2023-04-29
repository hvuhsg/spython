package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type context struct {
	*ir.Block
	fn       *ir.Func
	mod      *ir.Module
	parent   *context
	vars     map[string]value.Value
	regStack []value.Value
}

func newContext(mod *ir.Module, fn *ir.Func, b *ir.Block) *context {
	return &context{
		Block:  b,
		fn:     fn,
		mod:    mod,
		parent: nil,
		vars:   make(map[string]value.Value),
	}
}

func (c *context) newContext(name string) *context {
	b := c.fn.NewBlock(name)
	ctx := newContext(c.mod, c.fn, b)
	ctx.parent = c
	return ctx
}

func (c *context) getVar(name string) value.Value {
	if v, ok := c.vars[name]; ok {
		return v
	} else if c.parent != nil {
		return c.parent.getVar(name)
	} else {
		return nil
	}
}

func (c *context) createVar(name string, val value.Value) {
	c.vars[name] = val
}

func (c *context) pushReg(val value.Value) {
	c.regStack = append(c.regStack, val)
}

func (c *context) popReg() value.Value {
	if len(c.regStack) == 0 {
		panic("register not found")
	}

	index := len(c.regStack) - 1
	reg := (c.regStack)[index]
	c.regStack = (c.regStack)[:index]
	return reg
}

package compiler

import (
	"github.com/hvuhsg/spython/token"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type state struct {
	variables  map[string]value.Value
	token      token.Token
	module     *ir.Module
	function   *ir.Func
	block      *ir.Block
	regStack   []value.Value
	blockStack []*ir.Block
}

func newState() *state {
	s := &state{}
	s.regStack = make([]value.Value, 0)
	s.blockStack = make([]*ir.Block, 0)
	s.variables = make(map[string]value.Value)

	return s
}

func (s *state) addVariable(name string, val value.Value) {
	s.variables[name] = val
}

func (s *state) getVariable(name string) value.Value {
	return s.variables[name]
}

func (s *state) pushBlock(block *ir.Block) {
	s.blockStack = append(s.blockStack, block)
}

func (s *state) popBlock() *ir.Block {
	if len(s.blockStack) == 0 {
		return nil
	}

	index := len(s.blockStack) - 1
	block := (s.blockStack)[index]
	s.blockStack = (s.blockStack)[:index]
	return block
}

func (s *state) pushReg(val value.Value) {
	s.regStack = append(s.regStack, val)
}

func (s *state) popReg() value.Value {
	if len(s.regStack) == 0 {
		return nil
	}

	index := len(s.regStack) - 1
	reg := (s.regStack)[index]
	s.regStack = (s.regStack)[:index]
	return reg
}

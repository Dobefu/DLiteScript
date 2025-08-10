package token

import (
	"fmt"
	"sync"
)

// Pool provides a thread-safe pool of tokens.
type Pool struct {
	pool map[string]*Token
	mu   sync.RWMutex
}

// getPoolKey returns a unique key for the pool based on atom and token type.
func getPoolKey(atom string, tokenType Type) string {
	return fmt.Sprintf("%s:%d", atom, int(tokenType))
}

// NewPool creates a new token pool with some pre-allocated common tokens.
func NewPool() *Pool {
	p := &Pool{
		pool: make(map[string]*Token),
		mu:   sync.RWMutex{},
	}

	commonTokens := []struct {
		atom      string
		tokenType Type
	}{
		{fmt.Sprintf("+:%d", int(TokenTypeOperationAdd)), TokenTypeOperationAdd},
		{fmt.Sprintf("-%d", int(TokenTypeOperationSub)), TokenTypeOperationSub},
		{fmt.Sprintf("*:%d", int(TokenTypeOperationMul)), TokenTypeOperationMul},
		{fmt.Sprintf("/:%d", int(TokenTypeOperationDiv)), TokenTypeOperationDiv},
		{fmt.Sprintf("%%:%d", int(TokenTypeOperationMod)), TokenTypeOperationMod},
		{fmt.Sprintf("**:%d", int(TokenTypeOperationPow)), TokenTypeOperationPow},
		{fmt.Sprintf("(:%d", int(TokenTypeLParen)), TokenTypeLParen},
		{fmt.Sprintf("):%d", int(TokenTypeRParen)), TokenTypeRParen},
		{fmt.Sprintf(",:%d", int(TokenTypeComma)), TokenTypeComma},
		{fmt.Sprintf("=%d", int(TokenTypeAssign)), TokenTypeAssign},
		{fmt.Sprintf("{%d", int(TokenTypeLBrace)), TokenTypeLBrace},
		{fmt.Sprintf("}:%d", int(TokenTypeRBrace)), TokenTypeRBrace},
		{fmt.Sprintf("\n:%d", int(TokenTypeNewline)), TokenTypeNewline},
	}

	for _, t := range commonTokens {
		p.pool[getPoolKey(t.atom, t.tokenType)] = NewToken(t.atom, t.tokenType)
	}

	return p
}

// GetToken gets an existing token if it exists, or creates a new one.
func (p *Pool) GetToken(atom string, tokenType Type) *Token {
	key := getPoolKey(atom, tokenType)

	if token, exists := p.getFromPool(key); exists {
		return token
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	token := NewToken(atom, tokenType)
	p.pool[key] = token

	return token
}

// GetPoolSize returns the current size of the pool.
func (p *Pool) GetPoolSize() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.pool)
}

func (p *Pool) getFromPool(key string) (*Token, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	token, exists := p.pool[key]

	return token, exists
}

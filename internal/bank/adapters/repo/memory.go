package repo

import (
    "context"
    "sync"

    "github.com/alireza-aliabadi/golang-hexagonal/internal/bank/core"
)

type MemAccounts struct {
    mu   sync.RWMutex
    data map[core.AccountID]*core.Account
}

type MemTas struct {
    mu        sync.RWMutex
    byAccount map[core.AccountID][]*core.Transaction // ❗ changed from *Transaction to slice
}

// builder functions
func NewMemAccounts() *MemAccounts {
    return &MemAccounts{data: make(map[core.AccountID]*core.Account)}
}

func NewMemTas() *MemTas {
    return &MemTas{byAccount: make(map[core.AccountID][]*core.Transaction)} // ❗ initialize slice map
}

// MemAccounts methods
func (ma *MemAccounts) Save(ctx context.Context, a *core.Account) error {
    ma.mu.Lock()
    defer ma.mu.Unlock()
    acc := *a
    ma.data[a.ID] = &acc
    return nil
}

func (ma *MemAccounts) List(ctx context.Context) ([]*core.Account, error) {
    ma.mu.RLock()
    defer ma.mu.RUnlock()
    result := make([]*core.Account, 0, len(ma.data))
    for _, a := range ma.data {
        acc := *a
        result = append(result, &acc)
    }
    return result, nil
}

func (ma *MemAccounts) ByID(ctx context.Context, accID core.AccountID) (*core.Account, error) {
    ma.mu.RLock()
    defer ma.mu.RUnlock()
    a, ok := ma.data[accID]
    if !ok {
        return nil, core.ErrAccNotFound(accID)
    }
    acc := *a
    return &acc, nil
}

// MemTas methods
func (mt *MemTas) Add(ctx context.Context, t *core.Transaction) error {
    mt.mu.Lock()
    defer mt.mu.Unlock()
    ta := *t

    if ta.From != nil {
        mt.byAccount[*ta.From] = append(mt.byAccount[*ta.From], &ta) // ❗ dereference pointer
    }
    if ta.To != nil {
        mt.byAccount[*ta.To] = append(mt.byAccount[*ta.To], &ta) // ❗ dereference pointer
    }
    return nil
}

func (mt *MemTas) ByAccount(ctx context.Context, accID core.AccountID) ([]*core.Transaction, error) {
    mt.mu.RLock()
    defer mt.mu.RUnlock()
    arr := mt.byAccount[accID]
    result := make([]*core.Transaction, 0, len(arr))
    for _, x := range arr {
        tx := *x
        result = append(result, &tx)
    }
    return result, nil
}

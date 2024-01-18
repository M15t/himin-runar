package repo

import (
	"himin-runar/internal/types"

	repoutil "himin-runar/pkg/util/repo"

	"gorm.io/gorm"
)

// Memo represents the client for memo table
type Memo struct {
	*repoutil.Repo[types.Memo]
}

// NewMemo returns a new memo database instance
func NewMemo(gdb *gorm.DB) *Memo {
	return &Memo{repoutil.NewRepo[types.Memo](gdb)}
}

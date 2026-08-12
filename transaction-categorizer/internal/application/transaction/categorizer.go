package transaction

import  dt "github.com/revz79/go-gym/transaction-catgorizer/internal/domain/transaction"


type categorizer struct {

}

func NewCategorizer() *categorizer{
	return  &categorizer{}
}

func (c *categorizer) AssignCategory(txn *dt.Transaction) error {
	panic("not implemented")
}
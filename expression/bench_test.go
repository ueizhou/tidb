package expression

import (
	"testing"
	"github.com/pingcap/tidb/util/types"
	"github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/util/mock"
)

func prepareTransferExpr() Expression{
	col := &Column{Index:0, RetType:types.NewFieldType(mysql.TypeString)}
	con := &Constant{Value:types.NewIntDatum(1), RetType:types.NewFieldType(mysql.TypeLonglong)}
	args := []Expression{col, con}
	expr, _ := NewFunction(mock.NewContext(), "eq", types.NewFieldType(mysql.TypeLonglong), args...)
	return expr
}


func BenchmarkCmpFunc(b *testing.B) {
	b.StopTimer()
	expr := prepareTransferExpr()
	row := []types.Datum{types.NewStringDatum("1.0")}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(row)
	}
}

func prepareNoTransferExpr() Expression{
	col := &Column{Index:0, RetType:types.NewFieldType(mysql.TypeLonglong)}
	con := &Constant{Value:types.NewIntDatum(1), RetType:types.NewFieldType(mysql.TypeLonglong)}
	args := []Expression{col, con}
	expr, _ := NewFunction(mock.NewContext(), "eq", types.NewFieldType(mysql.TypeLonglong), args...)
	return expr

}


func BenchmarkCmpFuncNoTransfer(b *testing.B) {
	b.StopTimer()
	expr := prepareNoTransferExpr()
	row := []types.Datum{types.NewIntDatum(1.0)}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		expr.Eval(row)
	}
}

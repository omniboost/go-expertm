package expertm_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/omniboost/go-expertm"
)

func TestSales(t *testing.T) {
	ss := expertm.Import{
		Sales: expertm.Sales{
			{
				DocType:       10,
				CustomerPrime: 19110,
				DocNumber:     2000344,
				Details: expertm.SalesDetails{
					{
						Account: 9999,
						Amount:  12.0,
						DebCre:  expertm.SideDebet,
						Ref:     "Test debet",
					},
					{
						Account: 1234,
						Amount:  12.0,
						DebCre:  expertm.SideCredit,
						Ref:     "Test credit",
					},
				},
			},
		},
	}

	b, err := xml.MarshalIndent(ss, "", "  ")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}

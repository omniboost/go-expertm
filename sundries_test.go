package expertm_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-expertm"
)

func TestSundries(t *testing.T) {
	ss := expertm.Import{
		Sundries: expertm.Sundries{
			{
				JournalPrime: "11",
				DocNumber:    "1234",
				DocDate:      expertm.Date{Time: time.Now()},
				Status:       1,
				Details: expertm.SundryDetails{{
					Account: "12",
					Amount:  99.0,
					DebCre:  expertm.SideCredit,
					Ventil:  12,
					Ref:     "test",
				}},
			},
		},
	}

	b, err := xml.MarshalIndent(ss, "", "  ")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}

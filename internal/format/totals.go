package format

import (
	"os"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func DisplayGetTotalResponse(resp casheerapi.GetTotalResponse) {
	tw := table.NewWriter()
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter},
		{Number: 2, Align: text.AlignCenter},
		{Number: 3, Align: text.AlignCenter},
		{Number: 4, Align: text.AlignCenter},
	})
	tw.SetOutputMirror(os.Stdout)

	tw.AppendHeader(table.Row{"INCOME (EXPECTED)", "INCOME (ACTUAL)", "SPENT (EXPECTED)", "SPENT (ACTUAL)"})
	tw.AppendRow(table.Row{resp.Data.ExpectedIncome, resp.Data.RunningIncome, resp.Data.ExpectedTotal, resp.Data.RunningTotal})

	tw.Render()
}

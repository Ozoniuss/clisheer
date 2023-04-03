package format

import (
	"fmt"
	"os"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func DisplayListDebtResponse(resp casheerapi.ListDebtResponse, verbose bool) {

	if len(resp.Data) == 0 {
		fmt.Println("No active debts.")
		return
	}

	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)

	columnConfigs := []table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 2, Align: text.AlignRight, AlignHeader: text.AlignCenter},
		{Number: 3, Align: text.AlignLeft, AlignHeader: text.AlignCenter},
	}
	tw.SetColumnConfigs(columnConfigs)

	if verbose {
		columnConfigs = append(columnConfigs, table.ColumnConfig{
			Number: 4, Align: text.AlignCenter, AlignHeader: text.AlignCenter,
		})
	}

	tw.SetColumnConfigs(columnConfigs)
	header := table.Row{"Person", "Amount", "Details"}
	if verbose {
		header = append(header, "ID")
	}
	tw.AppendHeader(header)
	for _, d := range resp.Data {
		row := table.Row{d.Attributes.Person, d.Attributes.Amount, d.Attributes.Details}
		if verbose {
			row = append(row, d.ResourceID.Id)
		}
		tw.AppendRow(row)
	}

	tw.Render()
}

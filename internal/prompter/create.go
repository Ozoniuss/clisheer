package prompter

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ozoniuss/clisheer/internal/models"
)

// CreateDebtPrompter allows the user to insert a debt interactively.
func CreateDebtPrompter(header string) models.Debt {
	if header != "" {
		fmt.Println(header)
	}
	debt := models.Debt{}

	s := bufio.NewScanner(os.Stdin)

	debt.Person = scanStringUntilValid(s, "person: ")
	debt.Amount = scanFloatUntilValid(s, "amount: ")
	debt.Currency = scanStringUntilValid(s, "currency: ")
	debt.Details = scanStringUntilValid(s, "details: ")

	return debt
}

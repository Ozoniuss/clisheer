# Casheer-CLI (clisheer)

Clisheer is a command-line client to my expense manager backend [casheer](https://github.com/Ozoniuss/casheer). It is meant to provide a means for managing personal expense through a terminal.

List of commands
----------------

**Management/General**

- [x] `clisheer ping` checks whether the expenses server is running;
- [x] `clisheer set period --year 2023 --month 11 (or --now)` sets the "current" year and month (defaults to current year and month otherwise);
- [x] `clisheer show lite` shows how much money have been spent this month, and how much money is expected to be spent this month;
- [x] `clisheer show period` shows the current period that is set, or the current date if not set;
- [ ] `clisheer show all` shows the entire sheet of the month;

**Debts**

- [x] `clisheer show debts` shows active debts; 
- [ ] `clisheer add debt` allows adding a debt interactively;
- [x] `clisheer add debt --file debt.json` allows adding a debt from a file (initially only json, but multiple formats may be supported);
- [x] `clisheer delete debt 12345678-1234-5678-1234-123456781234` removes a single debt;
- [x] `clisheer update debt 12345678-1234-5678-1234-123456781234` updates a single debt;

**Others, not yet fully decided**

- [ ] `clisheer show categories` shows just the big categories with their running total;
- [ ] `clisheer show category food` shows the data for a single category;
- [ ] `clisheer add category food` adds a new category if it doesn't exist;
- [ ] `closheer add subcategory "food" "going out"` adds a new subcategory if it doesn't exist;
- [ ] `clisheer delete category` removes a category from this month and all expenses associated with it;
- [ ] `clisheer delete subcategory "food" "going out"` removes a subcategory and all expenses associated with it;
- [ ] `clisheer duplicate period` duplicates all categories for the next period, assuming no entry exists for the new period (otherwise must be forced to overwrite);
- [ ] `clisheer new period interactive` goes through all expenses of the current period, and interactively allows keeping or discarding subcategories;
- [ ] `clisheer scan jsonfile` scans through a json file of expenses from revolut format and interactively asks if you want to add all expenses;

Additional docs
---------------

- [State](./internal/state/) package;

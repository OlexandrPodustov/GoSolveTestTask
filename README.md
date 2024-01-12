Server with one endpoint for search the passed argument


remaining:
- add search
- Add `unit tests` for created components
- In case you’re not able to find index for given value, you can return `index` for any other existing value, assuming that conformation is at `10% level`, (for example, you were looking for `index` for value = `1150`, but in input file you have `1100` and `1200`, so in that case you can return index for `1100` or `1200`).

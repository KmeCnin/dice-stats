# Dice Probability Distribution Software
This software generate probability distribution for a given dice throw.

## Installation
*TODO*
## Use
### Throw query format
The only argument required is a string query following the format`{n}d{m}(k{p})` where:
- `n`: number of dice to throw
- `m`: number of faces of each dice
- `p`: number of best dice results to keep
### Example
Generate probability distribution for throwing 2 dice with 6 faces:
```shell
> ./dice-stats 2d6
```
[Output: 2d6](doc/2d6.png)

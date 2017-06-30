# Dice Probability Distribution Software
This software generate probability distribution for a given dice throw.

## Installation
*TODO*
## Use
### Throw query format
The only **argument** required is a string query following the format`{n}d{m}(k{p})` where:
- `n`: number of dice to throw
- `m`: number of faces of each dice
- `p`: number of best dice results to keep  

**Examples:** `2d6`, `3d10`, `7d6k3`.
### Command Option
#### Iterations
- Syntax: `-i`
- Type: `int`
- Default: `1000000`
- Example: `> ./dice-stats -i=10000 2d6`

Number of iterations to use in order to define probabilities accuracy.
#### Versus
- Syntax:`-vs`
- Type: `string`
- Default: *none*
- Example: `> ./dice-stats -vs=3d6 2d6`

Use another dice throw and compare chances to win against this one.
## Example
#### Generate probability distribution for throwing 2 dice with 6 faces
```shell
> ./dice-stats 2d6
```
![Output: 2d6](doc/output/2d6.png)
#### Generate probability distribution for throwing 3 dice with 10 faces
```shell
> ./dice-stats 3d10
```
![Output: 3d10](doc/output/3d10.png)
#### Generate probability distribution for throwing 7 dice with 6 faces and keeping the 3 best results
```shell
> ./dice-stats 7d6k3
```
![Output: 7d6k3](doc/output/7d6k3.png)

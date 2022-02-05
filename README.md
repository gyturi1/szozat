# Szozat

The [Szózat](https://szozat.miklosdanka.com/) is hungarian version of the [Wordle](https://www.powerlanguage.co.uk/wordle/) game.

This is a CLI to help finding the next guess if some of the letters are already known.

----------
## Word list

There is a [word list](../szozat/pkg/wordmap/words.txt) included in this program. It is generated with this [utility](https://github.com/gyturi1/cleanwords).

----------
## Usage

Download the binary from [here](release)

szozat -g "g u e s s" -l "a ny ly t dzs ű"
- -g flag: the last guess format below
- -l flag: the remainig letters space separated

Guess there must be space separated list of these (5 elements):
- "_" denotes an unknow letter on that position
- "L" denotes a know letter (L) not in the rigth position (orange)
- "*L" denotes a know letter (L) in the rigth position (green)

### Examples

`szozat -g "_ _ *cs ö k" -l "ü t ny ly"`

```
csöcsök
csücsök
köcsök
töcsök
tücsök
5
```
The last line is the nuber of suggestions found.

There will be false suggestions in the list. ("köcsök is not a valid hungarian word"). I still investigating to find a better source of words. Check back for updates. 
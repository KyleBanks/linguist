# Linguist

*This is early days for this project, and it has been throwned together in short order as a proof of concept. Everything is subject to change.*

`linguist` is a command-line tool to practise translating sentences between languages. 

```
$ linguist --lang de 
Hi
(Translate) > Hallo
✓

Run
(Translate) > Laufe
✓

Wow
(Translate) > Plot 
✘ Potzdonner
  Donnerwetter

Fire
(Translate) > Feuer
✓

Help
(Translate) > Hilf
✓
```

## Installation

```
go get -u github.com/KyleBanks/linguist
```

## TODO

- [ ] Difficulty scaling
- [ ] Randomization
- [ ] Mode variations (Translate from Native, Translate to Native)
- [ ] Additional language datasets (currently only supports learning German for native English speakers)
- [ ] Improvements to the game UI (colors, etc.)
- [ ] Improvements to validations (currently uses a simple levenshtein distance comparison. This should at the very least scale based on the sample difficulty)
- [ ] General code/structure improvements

## Acknowledgements

Translation dataset provided by [https://tatoeba.org](https://tatoeba.org).

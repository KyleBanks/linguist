# Linguist

*This is early days for this project, and it has been throwned together in short order as a proof of concept. Everything is subject to change.*

`linguist` is a command-line tool to practise translating sentences between languages. 

```
$ linguist --lang de 
Translate: Hi
> Hallo
✓

Translate: Run
> Lauf
✓

Translate: Wow
> Plotzinger
✘ Potzdonner
  Donnerwetter

Translate: Fire
> Feuer
✓

Translate: Help
> Hilf
✓
```

## Installation

```
go get -u github.com/KyleBanks/linguist
```

## How It Works

A translation dataset is loaded from the `dataset` directory, which contains tab-delimited translation pairs from a native language to the language being learned. For example, the most basic samples from the [English -> German dataset](./dataset/de.txt) look like so: 

```
Hi.	Hallo!
Hi.	Grüß Gott!
Run!	Lauf!
Wow!	Potzdonner!
Wow!	Donnerwetter!
Fire!	Feuer!
Help!	Hilfe!
Help!	Zu Hülf!
```

A sample is picked from the dataset and the user is prompted to enter the translation. A simple [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance) measurement is done on the sanitized input and possible translations, which is used to determine if the user entered the correct translation.

## TODO

There's a lot of work to make this a useful tool, and contributions are greatly appreciated:

- [ ] Difficulty scaling
- [ ] Randomization
- [ ] Non-binary results (when answer is close enough but not perfect, display the correct response)
- [ ] Mode variations (Translate from Native, Translate to Native)
- [ ] Additional language datasets (currently only supports learning German for native English speakers)
- [ ] Improvements to the game UI (colors, etc.)
- [ ] Improvements to validations (currently uses a simple/static levenshtein distance comparison. This should at the very least scale based on the sample difficulty)
- [ ] General code/structure improvements

## Acknowledgements

Translation dataset provided by [https://tatoeba.org](https://tatoeba.org).

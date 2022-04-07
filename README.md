# grop

Toy `grep` implementation in Go for learning purposes.

## Installation

```bash
go install github.com/harveysanders/grop
```

## Usage

Use it just like `grep`:

```shell
$ grop [-i] [--color=when] [term] [file]
```

`grop` accepts input files:

```shell
$ grop go go.mod

go 1.17
```

and `stdin`:

```shell
$ curl example.com | grop h1

% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1256  100  1256    0     0   6939      0 --:--:-- --:--:-- --:--:--  6901
    <h1>Example Domain</h1>
```

It currently only supports the following flags

- `--i` Case insensitive searches
- `--color=["always"|"auto"|"never"]` When to highlight matching patterns in results

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

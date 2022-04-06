# grop

grep implementation in Go

## Installation

```bash
go install github.com/harveysanders/grop
```

## Usage

Use it just like `grep`

`grop` accepts input files:

```shell
$ grop  go go.mod

go 1.17
```

and STDIN

```shell
$ curl example.com | grop h1

% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1256  100  1256    0     0   6939      0 --:--:-- --:--:-- --:--:--  6901
    <h1>Example Domain</h1>
```

It currently only supports the following flags

- `--i`
- `--color`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

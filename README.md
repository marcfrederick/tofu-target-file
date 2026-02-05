# tofu-target-file

A simple shell script that generates target files for `tofu apply`.
This allows you to `tofu apply` a single file instead of a whole directory, which can be useful in some situations.

## Usage

```bash
$ tofu apply -target-file <(tofu-target-file some_file.tf)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

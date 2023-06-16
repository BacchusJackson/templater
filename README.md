# templater

A very simple tool for templating text

## Configuration

must be in the format "KEY=VALUE"

```
FIRST_NAME=Tony
// This is a comment
LAST_NAME=Stark
AGE=47
OCCUPATION=an Avenger
```

Wherever the key is found in the `template.txt` file it will be replaced with the corresponding value

## Usage

```
Usage of templater:
  -config string
        The configuration file (default "config.txt")
  -template string
        The template file (default "template.txt")
```

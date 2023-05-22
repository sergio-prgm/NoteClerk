# NoteClerk | Create pretemplated markdown files

## What it does

NoteClerk is a simple CLI tool that helps avoid writing boilerplate in commonly used files (daily note, doc REAMDE, etc.)

Simply run the binary for the desired OS with the flag corresponding to the file that needs to be created.
Run `NoteClerk -h` to learn more about the different kinds of files that are currently available.

> It is convenient to add the executable to the PATH so that it is accessible everywhere and its location doesn't need to be known at all times.

```sh
./NoteClerk -daily
./NoteClerk -docs
```

## Roadmap

- [ ] Read from a template.md file to create custom files.
- [x] Create the actual file :)
- [ ] Get the necessary files by context (folder name, etc.)
- [ ] Add flag to create recursive docs (e.g. a file per Terraform module)

## Dev Commands

```sh
go build main.go && ./main -daily
```

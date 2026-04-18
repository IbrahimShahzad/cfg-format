# CFG-FORMAT

> [!NOTE] 
> the code in repo is AI assisted

main grammar and treesitter logic is based on `kamaizen`

```sh
KamaiZen/kamailio_cfg/parser.c          → cfg-format/grammar/parser.c
KamaiZen/kamailio_cfg/include.h         → cfg-format/grammar/include.h
KamaiZen/kamailio_cfg/tree_sitter/*.h   → cfg-format/grammar/tree_sitter/
```

Installation

```sh
git clone this repo
cd cfg-format
go install
```
usage

```sh
Usage: cfg-format [flags] [file ...]
  Formats Kamailio .cfg files. Reads stdin when no files are given.

  -check
        exit non-zero if any file is not already formatted
  -dump-tree
        print the parse tree and exit (debug)
  -indent int
        indent width (only used with -spaces) (default 4)
  -spaces
        use spaces instead of tabs for indentation (default true)
  -w    write result back to source file instead of stdout
  -width int
        max line length before if/while conditions are wrapped (default 79)
```

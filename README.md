# Gitfame
Golang is a console utility for calculating the statistics of the authors of the git repository

# Program Capabilities

- `--repository`: Path to the Git repository; defaults to the current directory.
- `--revision`: Commit reference; defaults to `HEAD`.
- `--order-by`: Key for sorting results; one of `lines` (default), `commits`, or `files`.
  - By default, results are sorted in descending order of the key (lines, commits, files).
  - If keys are equal, authors are sorted lexicographically by name.
  - When this flag is used, the corresponding field in the key is moved to the first position.

- `--use-committer`: A boolean flag that switches calculations from the author (default) to the committer.
- `--format`: Output format; one of `tabular` (default), `csv`, `json`, or `json-lines`.

### `tabular`:

```
Name Lines Commits Files
Joe Tsai 64 3 2
Ross Light 2 1 1
ferhat elmas 1 1 1
```

- Human-readable format. Padding is done using spaces (see `text/tabwriter`).

### `csv`:
```
Name,Lines,Commits,Files
Joe Tsai,64,3,2
Ross Light,2,1,1
ferhat elmas,1,1,1
```

### `json`:

```json
[
  {"name": "Joe Tsai", "lines": 64, "commits": 3, "files": 2},
  {"name": "Ross Light", "lines": 2, "commits": 1, "files": 1},
  {"name": "ferhat elmas", "lines": 1, "commits": 1, "files": 1}
]
```

### `json-lines`:
```json
{"name": "Joe Tsai", "lines": 64, "commits": 3, "files": 2}
{"name": "Ross Light", "lines": 2, "commits": 1, "files": 1}
{"name": "ferhat elmas", "lines": 1, "commits": 1, "files": 1}
```

- `--extensions`: List of file extensions that restricts the files considered for analysis. Separate multiple extensions with commas (e.g., .go,.md).
- `--languages`: List of programming or markup languages that restricts files considered for analysis. Separate multiple languages with commas (e.g., go,markdown).
- `--exclude`: Set of Glob patterns that exclude files from the analysis (e.g., foo/*,bar/*). Standard library's path/filepath can be used to work with Globs.
- `--restrict-to`: Set of Glob patterns that exclude all files not matching any pattern in the set.

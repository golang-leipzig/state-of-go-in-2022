# Hacking

Trying to identify active files in Go source tree.

* most touched files, in the past years
* average number of files changed per commit
* average number of commits per day

```
$ mergestat -f ndjson -r . "select commits.author_when, commits.hash, files.path from commits, files('', commits.hash)"  | jq -rc .
```

Or tsv (120kL/s)

```
$ mergestat -f tsv-noheader -r . "select commits.author_when, commits.hash, files.path from commits, files('', commits.hash)"
```

Stats:

```
$ mergestat -f tsv -r . "SELECT commits.hash, stats.* FROM commits, stats('', commits.hash)" > commts.tsv
```

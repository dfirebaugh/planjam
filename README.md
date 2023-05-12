# PlanJAM

PlanJAM is a simple keyboard based project management tool.

PlanJAM creates a `.plan` dir in the current working directory to save data in.  You can choose to ignore this from your repo or commit it.

## Download

#### Download and install with `go install`

```bash
go install github.com/dfirebaugh/planjam
```

For convenience, you could alias `planjam` to `pj`.
```bash
alias pj=planjam
```

> use the `--help` flag 
```bash
pj --help
```

## Example

```bash
# create a planjam board named planjam
pj board planjam 
# create some lanes
pj add lane todo
pj add lane in_progress
pj add lane procrastinate
pj rm lane procrastinate
pj add lane done
# create some features
pj add feature "make a feature"
pj add feature "fix a bug"
pj add feature "make a different feature"
# print the current board   
pj ls

| todo                          | in_progress | done |
|-------------------------------|-------------|------|
|  [0] fix_something            |             |      |
|  [1] make_a_feature           |             |      |
|  [2] fix_a_bug                |             |      |
|  [3] make_a_different_feature |             |      |

# print the board as an asciidoc table
pj ls -a

|===
| todo                     | in_progress | done

| fix_something            |             |
| make_a_feature           |             |
| fix_a_bug                |             |
| make_a_different_feature |             |
|===
```

#### Moving a feature

```bash
pj mv 0

|  [0] todo                |  [1] in_progress |  [2] done |
|--------------------------|------------------|-----------|
| fix_something            |                  |           |
| make_a_feature           |                  |           |
| fix_a_bug                |                  |           |
| make_a_different_feature |                  |           |
# it will prompt you which lane you want to move it to
Which lane should we move [make_a_different_feature] to? 2
Moving feature 'make_a_different_feature' to lane 'done' board:  planjam

| todo               | in_progress | done                          |
|--------------------|-------------|-------------------------------|
|  [0] fix_something |             |  [3] make_a_different_feature |
|  [1] make_a_feature|             |                               |
|  [2] fix_a_bug     |             |                               |
```

#### Show the stats
```bash
pj stat

# todo
- [3]: ██████████████████
# in_progress
- [0]: 
# done
- [1]: ██████
```


#### Adding a field to a feature
You can add notes to a feature by adding fields.

```bash
# e.g. pj add field [feature id] [field label] [field value]
pj add field 0 "url" "http://wikipedia.org"
```

You can look at details of a feature.
```bash
# pj ls [feature id]
pj ls 0

| fix_something                 |
|-------------------------------|
|     url: http://wikipedia.org |
```

# git-issue-ref

_Please read the [documentation](docs/git-issue-ref.md) for more information._

This tool provides automatic issue references in commit messages based on the current branch or a prefix in the commit message. You can also supply manual formats so you don't have to supply anything.

## Go get
`go get github.com/Rukenshia/git-issue-ref`

## Installation from Source
Currently you need to build it yourself using `go 1.7.5`(or newer). Just go into the package and use `go get && go build`.

## Examples

## default format
the following examples assume that you initialized issue-ref using `git issue-ref link`.

```
# on branch master
git commit -am "MYTEAM-123 did something"

--> "[MYTEAM-123] did something"
```

```
# on branch MYTEAM-123-some-awesome-task
git commit -am "did something"

--> "[MYTEAM-123] did something"
```

```
# on branch feature/MYTEAM-123-some-awesome-task
git commit -am "did something"

--> "[MYTEAM-123] did something"
```

## supplying a custom format
```
git issue-ref link --format "[MYTEAM-{ref}] "
git commit -am "123 did something"

--> "[MYTEAM-123] did something"
```

```
git issue-ref link --format "[MYTEAM-123] "
git commit -am "did something"

--> "[MYTEAM-123] did something"
```



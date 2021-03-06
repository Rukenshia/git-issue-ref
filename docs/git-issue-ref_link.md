## git-issue-ref link

Enable git-issue-ref for the current repository

### Synopsis


Use this command to enable git-issue-ref for the current repository.
This command will install a git hook that will, depending on your settings, automatically
prepend an issue reference or ask you to put one at the beginning of your message.

By default, commit messages where no ref could be found will be denied (exit status 1). You can override
this behavior by providing the --non-intrusive flag.

```
git-issue-ref link
```

### Options

```
  -f, --force           force overwriting existing hooks
      --format string   format of the commit message prefix (default "[{ref}] ")
      --non-intrusive   if no reference could be found, omit it and dont fail
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.git-issue-ref.yaml)
```

### SEE ALSO
* [git-issue-ref](git-issue-ref.md)	 - A tool to allow referencing git issues automatically

###### Auto generated by spf13/cobra on 21-May-2017

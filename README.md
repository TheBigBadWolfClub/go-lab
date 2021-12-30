# Go Lab

GoLang Monorepo and Multi-module for project(s), lib(s) && playground.

<hr>

### Howls


>The center of a wolf's universe is its pack, and howling is the glue that keeps the pack together.


Some Go projects …..

<hr>


### Spells

>__spell /spɛl/ spells__
><br>a form of words used as a magical charm or incantation.
><br>*"a spell is laid on the door to prevent entry"*

Set of tools and private libs to group common functionality required on more than one project.

<hr>

### Riddles


>__riddle /ˈrɪd(ə)l/ riddles__
><br>a question or statement intentionally phrased so as to require ingenuity in ascertaining its answer or meaning.
><br>*"they started asking riddles and telling jokes"*

Small unrelated pieces of GO code, to clarify, make proven of  or deepen the knowledge in GO

<hr>

## Module management

#### create
```shell
# create new module
# example create a module named "foo" under spells
> mkdir spells/foo
> cd spells/foo
> go mod init github.com/TheBigBadWolfClub/go-lab/spells/foo
 
```

#### release
```shell
# release version v0.0.1 of spells/foo
> git tag spells/foo/v0.0.1
> git push origin --tags

```


<hr>

# Source maintenance

#### linting
```shell

# install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# help
golangci-lint help linters

# run
golangci-lint run ./...

```


<hr>

# GO Project Layout



## Folders
Much more information (and more standard folders) can be found here: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

* `build` - Packaging and Continuous Integration. (See link above!)
* `cmd` - Main applications for this project, the folder should match the name of the executable.
* `internal` - Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself.
* `pkg` - Library code that's ok to use by external applications. Other projects will import these libraries expecting them to work, so think twice before you put something here.
* `scripts` - Scripts to perform various build, install, analysis, etc operations. These scripts keep the root level Makefile small and simple.



## Versioning
*(This is a TL;DR version of this: https://blog.golang.org/publishing-go-modules#TOC_3.)*



## License information
![WTFPL](license.png)

This projects uses the [WTFPL license](http://www.wtfpl.net/)
(Do **W**hat **T**he **F**uck You Want To **P**ublic **L**icense)
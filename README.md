# GO Project

This is a boilerplate template for new GO projects. 
**If you want to add or improve this boilerplate, create a PR**

## Folders
Much more information (and more standard folders) can be found here: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

* `build` - Packaging and Continuous Integration. (See link above!)
* `cmd` - Main applications for this project, the folder should match the name of the executable.
* `internal` - Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself.
* `pkg` - Library code that's ok to use by external applications. Other projects will import these libraries expecting them to work, so think twice before you put something here.
* `scripts` - Scripts to perform various build, install, analysis, etc operations. These scripts keep the root level Makefile small and simple.
## Install Linters
**These linters are important, as if plan to run on a CI/CD pipeline. If you code wont pass these linters, you wont be able to commit and/or merge it.**

###golangci-lint

Follow this guide:
[Local-Installation Guide (macOS/Linus/Windows)](https://golangci-lint.run/usage/install/#local-installation)

Afterwards you can use `make golint` to execute the linter.



### Git-Hooks
In order to make sure you commit only correctly formatted files please use the `make install-hooks` command to install the githooks.
The following hooks will be installed:
- post-checkout: automatically runs `go mod vendor` after checkout to have the correct dependencies for the branch
- pre-commit: automatically runs `go mod tidy` and `go mod vendor` and stops if that changed your go.mod or go.sum (which most likely means there were unnecessary dependencies in the go.mod before)
- pre-push: automatically runs `make check` before push to make sure you are not pushing bad code

## Versioning
*(This is a TL;DR version of this: https://blog.golang.org/publishing-go-modules#TOC_3.)*

Go modules use a semantic version, thats why our version names look like this vMAJOR.MINOR.PATCH

To create a new version, run do this:
```
// prepare depenendcies by removing local, development stuff
go vendor tidy

// create a new tag with git (with your acutal version)
git tag v1.2.3

// publish the new version
git push origin v1.2.3
```

The version v1.2.3 then will be available to be required via go mod by all projects.


## License information
![WTFPL](license.png)

This projects uses the [WTFPL license](http://www.wtfpl.net/)
(Do **W**hat **T**he **F**uck You Want To **P**ublic **L**icense)
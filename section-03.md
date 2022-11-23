# Your development environment

Few notes as I have a pretty good grasp of this.

This section has a really good introduction to terminals, prompts, the various
approaches mac/linux/unix/windows takes. Great stuff!


## Helpful commands

`go env` should be good for debugging environment variables relevant to go.


## Go modules

Used <https://go.dev/blog/using-go-modules> as reference.

Most of the stuff that I should make sure to remember is around modules.
Workspaces was the way to go previously, and may be something you run in to in
legacy code, but going forwards it is all modules.

Each source project is its own module with a `go.mod` that declares the module
and its dependencies. The list versions that is produced as a result of resolving
dependencies is put into `go.sum`. The resolver preferes the lowest compatible
version. You can upgrade a dependency to the latest version via

```shell
# Latest
go get <dependency>
# To pull a specific version
go get <dependency>@version
# eg
go get rsc.io/sampler@v1.3.1
```

`go list` is used for listing packages and modules. By specifying `-m` we can
list modules, and by adding an `all` we get all modules

```shell
$ go list -m all
example.com/hello
golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4
golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f
golang.org/x/text v0.4.0
golang.org/x/tools v0.1.12
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.1
```

# State of Go in 2022

> What's new (since Go [1.14](https://go.dev/doc/devel/release#go1.14) ~ 02/20)?

Talk at [Leipzig Software Craft
Meetup](https://www.meetup.com/LE-software-craft-community/), [2022-06-30 19:00 CEST](https://www.meetup.com/le-software-craft-community/events/286403431/)

> [Andreas Linz](https://github.com/klingtnet), [Martin Czygan](https://github.com/miku) ([LI](https://www.linkedin.com/in/martin-czygan-58348842))

## Hello 世界!

* hosting [Leipzig Gophers](https://golangleipzig.space/) since [2019](https://golangleipzig.space/posts/meetup-launched/)
* 25+ events hosted, input presentations, code walkthroughs, discussions, startups, quizzes, sponsors, ...
* 400+ members on meetup, the pandemic brought us a truly international audience 🇮🇩🇲🇪🇨🇱🇧🇷🇧🇾🇺🇸🇧🇪🇩🇪🇨🇳🇮🇷 …
* we have a [YouTube channel](https://www.youtube.com/channel/UCFDzViL6Bo0w2AG23Q0_rZQ) 📹

![](static/leipzig-gopher.png)

## Change to the language

```shell
$ git clone https://github.com/golang/website.git
$ grep -l "There are no.*changes to the language." website/_content/doc/go1.1*html
_content/doc/go1.10.html
_content/doc/go1.11.html
_content/doc/go1.12.html
_content/doc/go1.15.html
_content/doc/go1.16.html
```

5 out of the 9 past releases (Go 1.10, 2018-02-16 to 1.18, 2022-03-15) do not
contain *any ... changes* to the language.

The single biggest change to the language is the addition of [generic
types](https://en.wikipedia.org/wiki/Generic_programming) in Go 1.18. All other
changes have been small enhancements (like number literals, changes in
`unsafe`, ...). Some say, we already have [Go 2](https://xeiaso.net/blog/we-have-go-2).

## So, what changed?

We want to highlight a few changes and how they affect everyday Go development:

* move to **Go modules** for dependency management
* introduction of **generic types**
* **standard library** additions:
  [`hash/maphash`](https://pkg.go.dev/hash/maphash) (1.14),
[`testing/Cleanup`](https://pkg.go.dev/testing#T.Cleanup),
[`time/tzdata`](https://pkg.go.dev/time/tzdata) (1.15),
[`embed`](https://pkg.go.dev/embed) (1.16),
[`testing/SetEnv`](https://pkg.go.dev/testing#T.Setenv),
[`debug/buildinfo`](https://pkg.go.dev/debug/buildinfo) (1.18),
[`net/netip`](https://pkg.go.dev/net/netip) (1.18)
* **tools** evolution (e.g. fuzz testing)

As well as use cases, users and **ecosystem** changes.

## Go repo analysis

In fact, Go turned 50 this year.

```
$ git summary

 project  : go
 repo age : 50 years
 active   : 4822 days
 commits  : 52739
 files    : 11670
 authors  :
  7018  Russ Cox                                                    13.3%
  3854  Robert Griesemer                                            7.3%
  2983  Rob Pike                                                    5.7%
  2360  Brad Fitzpatrick                                            4.5%
  2297  Ian Lance Taylor                                            4.4%
  1537  Austin Clements                                             2.9%
  1496  Josh Bleecher Snyder                                        2.8%
  1398  Matthew Dempsky                                             2.7%
  1319  Keith Randall                                               2.5%
  1192  Andrew Gerrand                                              2.3%
  1026  Cherry Zhang                                                1.9%
   935  Bryan C. Mills                                              1.8%
   ...
```

Over 2000 committers to the core project (973 authors with more than single commit, 277 with more than 10 commit):

```
$ git shortlog -s -n | wc -l
2189
```

From a [git-of-theseus](https://github.com/erikbern/git-of-theseus) analysis
(as of [0a1a092c4b56a1d4033372fbd07924dad8cbb50b](github.com/golang/go/commit/0a1a092c4b56a1d4033372fbd07924dad8cbb50b)): authors (top 20 plus others),
code age, extensions (move from [C to Go](https://go.dev/doc/go1.5#c) in 1.5):

![Various plots generated with git-of-theseus](static/theseus/stats_combined.png)

## More stats

Using [mergestat]() for git repo analysis:

```
$ mergestat -f tsv -r . "SELECT commits.hash, stats.* FROM commits, stats('', commits.hash)" > commts.tsv
```

Drop into pydata:

> Commits per year

```
In [37]: df.groupby(df.date.dt.year)["hash"].nunique()
Out[37]:
date
1972       1
1974       1
1988       2
2008    1396
2009    3116
2010    2501
2011    3994
2012    3754
2013    3378
2014    3343
2015    4784
2016    4803
2017    4258
2018    3887
2019    3392
2020    3989
2021    4825
2022    1508
Name: hash, dtype: int64
```

> Top 5 files changed per year (2018-2022)

```python
>>> df.groupby(df.date.dt.year)["dir"].value_counts().sort_values().to_csv("dirfreq.csv")
```

```csv
2018,src/cmd/compile/internal/ssa,854
2018,src/cmd/vendor/golang.org/x/sys/unix,881
2018,src/syscall,914
2018,src/cmd/compile/internal/gc,1174
2018,src/runtime,2258

2019,src/cmd/compile/internal/ssa,753
2019,src/cmd/compile/internal/gc,877
2019,src/cmd/go/testdata/script,1078
2019,src/vendor/golang.org/x/sys/unix,1106
2019,src/runtime,2727

2020,src/cmd/compile/internal/ssa,1475
2020,src/cmd/vendor/golang.org/x/sys/unix,1883
2020,src/cmd/go/testdata/script,2199
2020,src/cmd/compile/internal/gc,2488
2020,src/runtime,2835

2021,src/cmd/compile/internal/types2,2085
2021,src/cmd/go/testdata/script,2094
2021,src/go/types,2159
2021,src/cmd/vendor/golang.org/x/sys/unix,2476
2021,src/runtime,5202

2022,src/cmd/compile/internal/types2,516
2022,src/go/types,530
2022,src/go/types/testdata/fixedbugs,646
2022,test/typeparam,663
2022,src/runtime,1143
```

Most changes (dir):

```
$ cat changes.tsv | tail -20 | column -t
src/cmd/go/internal/modload                  52171   26710   78881
src/cmd/compile/internal/ir                  60114   23182   83296
src/cmd/oldlink/internal/ld                  27955   59698   87653
src/cmd/vendor/golang.org/x/arch/x86/x86asm  44218   48632   92850
src/cmd/compile/internal/typecheck           75263   23292   98555
src/crypto/tls/testdata                      56391   55537   111928
src/cmd/compile/internal/ssa/gen             66771   56927   123698
src/cmd/compile/internal/noder               106658  24330   130988
src/cmd/go/testdata/script                   124204  14609   138813
src/cmd/link/internal/ld                     79214   60428   139642
src/go/types                                 102641  54215   156856
src/cmd/compile/internal/types2              147363  44516   191879
src/syscall                                  120857  83241   204098
doc                                          45680   201068  246748
src/cmd/compile/internal/gc                  100028  284737  384765
src/runtime                                  279826  136985  416811
src/time/tzdata                              229072  213100  442172
src/vendor/golang.org/x/sys/unix             408978  445375  854353
src/cmd/vendor/golang.org/x/sys/unix         411688  540942  952630
src/cmd/compile/internal/ssa                 736274  990519  1726793
```

A few other packages as comparison:

```
$ cat changes.tsv | grep -E "(src/io|src/net|src/os)" | sort -nrk4,4 | head -20 | column -t
src/net/http                 28520  11507  40027
src/net                      20462  10633  31095
src/os                       18546  6748   25294
src/net/netip                9808   318    10126
src/io/fs                    6054   507    6561
src/os/exec                  3966   2352   6318
src/os/signal                3603   1204   4807
src/net/http/httputil        3231   549    3780
src/net/url                  2609   919    3528
src/io/ioutil                1020   1247   2267
src/os/user                  1273   660    1933
src/io                       1206   386    1592
src/net/http/cgi             1045   546    1591
src/net/http/pprof           1121   250    1371
src/net/mail                 986    118    1104
src/net/http/httptest        862    200    1062
src/net/http/fcgi            695    189    884
src/net/smtp                 752    71     823
src/net/http/internal/ascii  780    0      780
src/net/textproto            509    268    777
```

## Generics

Or: [parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)

Relevant proposal:

* [43651](https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md) (2021, accepted)

Previous discussion:

> Go should support some form of generic programming.
Generic programming enables the representation of algorithms and data
structures in a generic form, with concrete elements of the code (such as
types) factored out. It means the ability to express algorithms with minimal
assumptions about data structures, and vice-versa.

[...]

> Generics permit type-safe polymorphic containers. Go currently has a very
> limited set of such containers: slices, and maps of most but not all types.
> Not every program can be written using a slice or map.

> Look at the functions SortInts, SortFloats, SortStrings in the sort package.
> Or SearchInts, SearchFloats, SearchStrings. Or the Len, Less, and Swap
> methods of byName in package io/ioutil. **Pure boilerplate copying**.

Who wrote that? [...] [Proposal: Go should have generics](https://github.com/golang/proposal/blob/master/design/15292-generics.md) (2011)

A few more:

> As Russ pointed out, **generics are a trade off between programmer time, compilation time, and execution time.**

### Key Ideas

* Functions and types can have **type parameters**, which are defined using **constraints**, which are **interface types**.
* Constraints describe the **methods required** and the **types permitted** for a type argument.
* Constraints describe the methods and operations permitted for a type parameter.
* **Type inference** will often permit omitting type arguments when calling functions with type parameters.

> This design is completely backward compatible.

There are many things Go Generics do not support: No metaprogramming (macros);
no operator methods (e.g. no syntax like `c[i]` for custom types); no
[covariance or contravariance](https://en.wikipedia.org/wiki/Covariance_and_contravariance_(computer_science)) of function parameters; ...


### Notable changes

* `any` (alias for `interface{}`, [operations
permitted](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#operations-permitted-for-any-type))
and `comparable` as [predeclared
identifiers](https://go.dev/ref/spec#Predeclared_identifiers)
* a new [contraints](https://pkg.go.dev/golang.org/x/exp/constraints) package,
not in the standard library yet (but in
[exp](https://pkg.go.dev/golang.org/x/exp#section-readme), *In short, code in
this subrepository is not subject to the Go 1 compatibility promise.*)

### Syntax

```
func G[T Constraint](p T) { }
```

* contraint may be `any` ~ `interface{}` or `comparable` or an interface with methods and

### Other Bits

* interfaces allow to capture common behaviour

> In other words, interface types in Go are a form of generic programming. They
> let us capture the common aspects of different types and express them as
> methods. -- [Go generic programming today](https://go.dev/blog/why-generics)

There is a container package in the standard library, not widely used:

* [https://pkg.go.dev/container/list@go1.18.3](https://pkg.go.dev/container/list@go1.18.3)

> the Go standard library package container is mostly unused. Everything in the
> container package deals with interface{}/any values, which is Go for
> "literally anything". -- [We have Go 2](https://xeiaso.net/blog/we-have-go-2)

### Examples


#### Reverse

Finally, we can write a generic reverse for slices of **any** type.

* [https://go.dev/play/p/LDdnZMiQN6b](https://go.dev/play/p/LDdnZMiQN6b)

```go
package main

import "fmt"

// Reverse returns a new slice, with elements in reverse order.
func Reverse[T any](vs []T) []T {
    var (
        n      = len(vs)
        result = make([]T, n)
    )
    for i, v := range vs {
        result[n-1-i] = v
    }
    return result
}

func main() {
    var (
        s  = []string{"a", "b", "c"}
        i  = []int{1, 2, 3}
        f  = []float64{1, 2, 3}
        rs = Reverse(s)
        ri = Reverse(i)
        rf = Reverse(f)
    )
    fmt.Printf("%v %T\n", rs, rs)
    fmt.Printf("%v %T\n", ri, ri)
    fmt.Printf("%v %T\n", rf, rf)
    // [c b a] []string
    // [3 2 1] []int
    // [3 2 1] []float64
}
```


#### Interface Constraint

* [https://go.dev/play/p/MM38gRuTRKB](https://go.dev/play/p/MM38gRuTRKB)

```go
// Dummy example, showing contraints as interface type and type parameters.

type Number interface {
    int | int64 | float64
}

func Min[Number T](u, v T) T {
    if u < v {
        return u
    }
    return v
}
```

#### Constraints and Methods

* [https://go.dev/play/p/HsJtEJnzJdl](https://go.dev/play/p/HsJtEJnzJdl)

```go
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
	String() string
}

func Min[T Number](u, v T) T {
	if u < v {
		return u
	}
	return v
}

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("[%d]", i)
}

func main() {
	r := Min(MyInt(3), MyInt(4))
	fmt.Printf("%v %T", r, r)
	// 3 int
}
```

#### Set

```go
// Package sets implements sets of any comparable type.
package sets

// Set is a set of values.
type Set[T comparable] map[T]struct{}

// Make returns a set of some element type.
func Make[T comparable]() Set[T] {
    return make(Set[T])
}

// Add adds v to the set s.
// If v is already in s this has no effect.
func (s Set[T]) Add(v T) {
    s[v] = struct{}{}
}

// Delete removes v from the set s.
// If v is not in s this has no effect.
func (s Set[T]) Delete(v T) {
    delete(s, v)
}

// Contains reports whether v is in s.
func (s Set[T]) Contains(v T) bool {
    _, ok := s[v]
    return ok
}
```

#### Dot Product

```go
// Numeric is a constraint that matches any numeric type.
// It would likely be in a constraints package in the standard library.
type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~complex64 | ~complex128
}

// DotProduct returns the dot product of two slices.
// This panics if the two slices are not the same length.
func DotProduct[T Numeric](s1, s2 []T) T {
    if len(s1) != len(s2) {
        panic("DotProduct: slices of unequal length")
    }
    var r T
    for i := range s1 {
        r += s1[i] * s2[i]
    }
    return r
}
```


### Summary

* basic generic programming support
* libraries already on the way: [https://github.com/samber/lo](https://github.com/samber/lo), [https://github.com/elliotchance/pie](https://github.com/elliotchance/pie)

Examples:

```go
names := lo.Uniq([]string{"Samuel", "Marc", "Samuel"})
```

Calculating an average or finding the mode:

* [https://pkg.go.dev/github.com/elliotchance/pie/v2#Average](https://pkg.go.dev/github.com/elliotchance/pie/v2#Average)
* [https://pkg.go.dev/github.com/elliotchance/pie/v2#Mode](https://pkg.go.dev/github.com/elliotchance/pie/v2#Mode)

And more. Functional programming gets more realistic in Go:

* [https://pkg.go.dev/github.com/samber/mo#Either](https://pkg.go.dev/github.com/samber/mo#Either)
* [https://github.com/nikgalushko/fx](https://github.com/nikgalushko/fx)

Generic data structures:

* [https://github.com/zyedidia/generic](https://github.com/zyedidia/generic)

Generic concurrency patterns:

* [https://github.com/lobocv/simpleflow](https://github.com/lobocv/simpleflow)

We'll see more of these this year; standard library may get some generic
support for slices and maps (like map keys, etc).


### Performance

Following is a list of benchmark results where a large PNG image was resized to simulate a CPU heavy task.

The benchmark output format is:

```
<benchmark name>-<#cpu cores> <#iterations>
```

```sh
Benchmarking Go 1.9
BenchmarkImageResizing-4              10         162645758 ns/op
Benchmarking Go 1.10
BenchmarkImageResizing-4              10         161542379 ns/op
Benchmarking Go 1.11
BenchmarkImageResizing-4              10         151867950 ns/op
Benchmarking Go 1.12
BenchmarkImageResizing-4              10         145375775 ns/op
Benchmarking Go 1.13
BenchmarkImageResizing-4               7         149602369 ns/op
Benchmarking Go 1.14
BenchmarkImageResizing-4               7         144319423 ns/op
Benchmarking Go 1.15
BenchmarkImageResizing-4               7         143604268 ns/op
Benchmarking Go 1.16
BenchmarkImageResizing-4               8         140225932 ns/op
Benchmarking Go 1.17
BenchmarkImageResizing-4               8         141211938 ns/op
Benchmarking Go 1.18
BenchmarkImageResizing-4               8         129669589 ns/op
```

CPU intensive tasks are consistently getting faster with each Go release.
Overall we see a 25% speedup with the latest Go release compared to Go 1.9.

- Go 1.11 included [big improvements for `arm64` targets](https://go.dev/doc/go1.11#performance).
- Go 1.18 applied [register based calling convention](https://go.googlesource.com/proposal/+/refs/changes/78/248178/1/design/40724-register-calling.md) for ARM targets.

### Testing

Testing hasn't changed much, except the addition of [fuzz testing](https://go.dev/doc/tutorial/fuzz).

But, some utility methods were added to the already great `testing` library.
Below is an example that uses `t.Cleanup` as well as `t.SetEnv`, which were added with Go 1.14 and 1.17:

```go
package main

import (
	"os"
	"testing"
)

func TestDemonstrateAdditions(t *testing.T) {
	err := os.WriteFile("./test.json", []byte(`{"foo": "bar"}`), 0o600)
	if err != nil {
		t.Fatal(err.Error())
	}
	// Cleanup will be called when test finishes, i.e. if it has failed or succeeded.
	defer t.Cleanup(func() {
		t.Log("called")
		os.Remove("./test.json")
	})

	// This will be unset automatically after test finishes.
	t.Setenv("USERNAME", "andreas")

	// ...
}

```

## Popular libraries and Frameworks

Following is a list of popular frameworks and libraries, which you can take as a reference.  Of course, the most popular package is not always the best choice, but often a reliable one.

Our method for determining what the most popular frameworks and libraries are, is to simply sort them by GitHub stars.  Go's [package search](https://pkg.go.dev/search) does not allow to sort by number of imports.

Most popular:

- web framework is [`gin`](https://github.com/gin-gonic/gin)
- HTTP client is [`resty`](https://github.com/go-resty/resty)
- HTTP router is [`gorilla/mux`](https://github.com/gorilla/mux)
- microservice framework is [`kit`](https://github.com/go-kit/kit)
- testing library is [`testify`](https://github.com/stretchr/testify)
- object relational mapper is [`gorm`](https://github.com/go-gorm/gorm)
- logging library is [`logrus`](https://github.com/sirupsen/logrus)

This list is probably no surprise for most longer term Go developers, since those packages exist since quite some time and their popularity hasn't changed much.

But, please be aware that you can come a long way without needing any third party dependencies!

Also, if you need to choose a dependency make sure that it's somehow compatible with standard libraries interface's (e.g. the ubiquitous `http.Handler` or `io.Reader` and `io.Writer`).

Onboarding new developers will be much easier then, since they don't have to learn new paradigms just to do the same thing.
Also, it will help when needing to replace a dependency, since a bit of compatibility is guaranteed.

## Packaging

Go started with a very opinionated design choice on how to organize your code, the `$GOPATH`.
A typical `GOPATH` looked like this:

```
~/go
   /bin # executables installed via go get (now go install)
   /pkg # pre-compiled libraries
   /src
      /company.com/a/project # a personal project
         /.git/
         /main.go
         /vendor/ # contains vendored/copied dependencies
         /...
      /github.com/dependency/x # dependencies
      /github.com/dependency/y
```

Developers could not use their typical `~/code` folder (or similar) to setup a Go project.

Instead they needed to work inside `$GOPATH/src/<import-path>`, which is pretty inconvenient.  There were some workarounds around this limitation, e.g. setting up repo specific GOPATH's, but this often broke tooling.

The biggest limitation of this code organization was, that you could only use one version of any dependency at a time, even across projects.

Except, dependencies were vendored.  But, this resulted often in huge git repositories.

At Google this wasn't a problem since everything was stored in a big monorepo anyways, but most companies didn't work with this code organization setup.

A bunch of tools were developed by the community to workaround the dependency problem, with `dep` being the most popular one.

However, the Go team decided to work on their own concept which would then become _Go Modules_.  Modules can be used with the `go` tool and don't require any third-party tool but in the beginning they were not enabled by default.

Still, adoption of Go Modules was already very high in 2019 as the following chart from [2019's developer survey results](https://go.dev/blog/survey2019-results) shows:

![Adoption of Go Modules in 2019](https://go.dev/blog/survey2019/fig22.svg)

With Go 1.14, released on 25th of February 2020, Modules became the default.
Ultimately, this shows how much need there was for a fast and easy to use packaging solution.

Since then the `GOPATH` is nothing to care about anymore, instead a new Go project can be setup in any folder using the following commands:

```sh
$ mkdir myproject && cd myproject
$ go mod init myproject
```

There's also a official tutorial on [how to write Go code](https://go.dev/doc/code) that explains a project setup in more detail.

Not only did Modules made development more convenient, they also helped to ensure reproducibility and authenticity of dependencies.

The `go` command verifies a cryptographic hash of each (public) direct and indirect dependency against a global hash sum database to ensure their authenticity.  Those checksums are what's stored inside a Modules [`go.sum` file](https://go.dev/ref/mod#go-sum-files).

## Embed

Deployment and distribution was always one of the strong points of Go since it can be easily (cross-) compiled into a statically linked binary that is then shipped.

With the [`embed` package](https://pkg.go.dev/embed) there is now a standard way of bundling any number of files or directories into a binary.

A webserver for example can be shipped as a single binary including all its assets, templates, Javascript files and so on.

Another common usecase is to bundle SQL migration files.
Embedding something is very easy, all it takes to embed a directory is this:

```go
//go:embed assets
var assets embed.FS

// Here's how to serve this directory from an HTTP server:
router.Handle("/assets/", http.FileServer(http.FS(assets)))
```

Individual files can be directly embedded as `[]byte` or `string`.


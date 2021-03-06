# make-the-journey-map

:warning: This project is currently a prototype. It will be rewritten with decent standards once these are proven:

- [x] It can store a tech ~~tree~~ graph.
- [x] The graph is validated for syntax.
- [x] The graph is validated for references.
- [x] The graph can be printed in [DOT format][dot].
- [ ] The graph visibly contains the citation URLs.
- [ ] A decent view can be generated without external tooling.
  I think an octilinear routing layout and an isometric view would look nice.

## Current view

![View](docs/map.svg)

## Running

### Short version

```
bin/check-and-generate.sh
```

will validate and generate the current version of `docs/map.png`.

### Detailed version

To validate the entire map:

```
go run .
```

To display a `.dot` format output:

```
go run . -graph
```

To render the graph (assumes MacOS and `graphviz` installed):

```
go run . -graph | dot -Tpng -Gdpi=200 -odocs/map.png && open docs/map.png
```

All commands with exit non-zero status if the map is invalid, but only `go run .` will tell why.

[dot]: https://en.wikipedia.org/wiki/DOT_(graph_description_language)

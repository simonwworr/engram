# Engram

> A fast, local-first memory layer for AI coding assistants.

Engram is a fork of [Gentleman-Programming/engram](https://github.com/Gentleman-Programming/engram) — a lightweight tool that indexes your codebase into compressed, semantic chunks so AI assistants can retrieve relevant context without sending your entire project to a remote server.

## Features

- **Local-first**: All indexing and chunk storage happens on your machine.
- **Fast retrieval**: Compressed `.jsonl.gz` chunks enable quick semantic lookups.
- **CLI-driven**: Simple commands to index, query, and manage your memory store.
- **Plugin marketplace**: Extend engram with community plugins via `.claude-plugin/marketplace.json`.

## Installation

```bash
go install github.com/your-org/engram@latest
```

Or build from source:

```bash
git clone https://github.com/your-org/engram.git
cd engram
go build -o engram ./cmd/engram
```

## Usage

### Index your project

```bash
engram index .
```

This scans the current directory, splits files into semantic chunks, and writes compressed chunks to `.engram/chunks/`.

### Query the memory store

```bash
engram query "how does authentication work"
```

Returns the most relevant chunks from your indexed codebase. By default, returns the top 5 results — use `--limit` to adjust:

```bash
engram query "how does authentication work" --limit 20
```

### Show manifest

```bash
engram manifest
```

Prints the current `.engram/manifest.json` — a summary of all indexed files and chunk metadata.

### Clear the store

```bash
engram clear
```

Removing all chunks and resets the manifest.

## Configuration

Engram reads configuration from `.engram/manifest.json`. You can adjust chunk size, ignore patterns, and embedding settings there.

## Plugin Marketplace

Plugins are listed in `.claude-plugin/marketplace.json`. To install a plugin:

```bash
engram plugin install <plugin-name>
```

## Contributing

We welcome contributions! Please open an issue using one of the templates in `.github/ISSUE_TEMPLATE/` before submitting a pull request.

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/my-feature`
3. Commit your changes following [Conventional Commits](https://www.conventionalcommits.org/)
4. Open a pull request

## License

MIT — see [LICENSE](LICENSE) for details.

# forwarding

Fork of [noble-assets/forwarding](https://github.com/noble-assets/forwarding). IBC middleware module that manages forwarding accounts for automatic token routing through predefined channels.

## Spec

See [`spec/`](./spec/) for module specification:

- [State](./spec/01-state.md) - Forwarding account structure and storage
- [Messages](./spec/02-messages.md) - Tx messages
- [Events](./spec/03-events.md) - Emitted events
- [Queries](./spec/04-queries.md) - gRPC/REST queries
- [CLI](./spec/05-cli.md) - Command-line interface

## Build

```sh
go build ./...
```

## E2E Tests

```sh
cd e2e && go test -v ./...
```

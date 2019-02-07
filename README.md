# Go Starter kit

Setting up my own Go starter kit featuring:

- [x] Docker containers to compile, run and expose the application(s)
- [x] Docker Compose to start all apps and setup networking between them as needed
- [x] `realize` to auto-recompile (and install dependencies as introduced!) on code changes


Plan to add:

- [ ] ~~Go Modules support~~ - holding this for a while. Realize is not compatible with Go modules yet.
- [ ] ELK stack
- [ ] Filebeat sidecar to ship logs to the ELK stack
- [ ] Testing (they should run on each save/launch)


Examples:

- `hello-world` HTTP Server hello world
- `todos-api` TODOs API
- `logger` Just a dumming project that's not a HTTP server

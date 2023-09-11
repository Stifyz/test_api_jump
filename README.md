# Commands

To launch api:
`docker-compose up`

To launch docker in dev mod:  
`docker-compose up db`  
`docker build -f Dockerfile.dev . -t api-dev:latest && docker run --rm -it -v $(pwd):/api -p 8080:8080 --network api-network api-dev:latest /bin/sh`

# Arch

- `cmd/`                          is the main package
- `internal/db`                   is the db initilization
- `internal/handlers`             contains a set of fonctions that returns requestHandlers
- `internal/mocks`                fakes structs used for tests
- `internal/mocks/data`           fakes data used for tests
- `internal/models`               structs which defines objects passing through the api
- `internal/repositories`         used by services and performs DB calls
- `internal/repositories/errors`  defines errors which can be returned by the repositories
- `internal/routes`               setups handlers on route's paths
- `internal/services`             makes links between handlers and repositories

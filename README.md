# Grafana "Kitchen Sink" Example

This project is a simple demo of how to use grafana plugins with a simple
golang based http server backend, deployed and run using docker-compose

### Go-based http server

We will simply follow the [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin) 
tutorial from the official golang docs site

Accessible at http://localhost:8080/albums (see the tutorial for more endpoints)

### Grafana

Grafana is run in a docker container on port 3000

The plugins need to be built before they can be used in grafana, run `yarn dev` to compile, or `yarn watch` to watch 
for changes and automatically recompile. Remember to restart grafana if adding a new plugin.

## Running the example

Use docker-compose to start the golang api server on :8080 and Grafana on :3000

Grafana will load the plugins as they currently exist on the filesystem, use `docker-compose restart grafana` to load
any new plugins. The container is bound to the filesystem so you don't need to restart the container when recompiling an
existing plugin (only when adding a new one).
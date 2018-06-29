# Cartola REST API

The `cartola-rest-api` service is responsible for providing a way to access our MongoDB database in a single, unified API.

To run it, you need to install [Golang](https://golang.org/doc/install) and [MongoDB](https://docs.mongodb.com/manual/installation/). Follow the links for installation instructions.

After that, run the following commands inside your project:

``` bash
dep ensure
go run main.go
```

Start MongoDB if you haven't already (`brew services start mongodb` on macOS, for example), in order to make the API execute operations in the database.

## Endpoints

### GET /players
Returns a JSON array of all players in the championship.
  
### POST /players
The body must contain a JSON array of players, similar to the one in the CartolaFC [endpoint](https://api.cartolafc.globo.com/atletas/mercado).
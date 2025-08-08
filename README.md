# starzz-gin

This is a REST API backend created using the Gin framework of Go.

### The Dataset

This project uses a database of fictional galaxies, constellations and stars.  

Here is a diagram to describe the tables and their relationships:

<img src="assets/schema.png" width="500" height="200"/>

Stars are located in constellations, which are in turn located in galaxies.

The `galaxies`, `constellations` and `stars` tables contain the additional
fields `added_by` and `verified_by` to indicate the id of the users who made
the finding and verified it, respectively.

The database was created using SQLite.  The scripts to create the tables and
load the dummy data are included in `assets` for reference.  Note that the primary
keys of each table should actually increment automatically but are simply defined
as `INTEGER` and `PRIMARY KEY`, like so:

    CREATE TABLE users (
        user_id INTEGER,
        ...
        PRIMARY KEY (user_id)
    );
    
    CREATE TABLE galaxies (
        galaxy_id INTEGER,
        ...
        PRIMARY KEY (galaxy_id),
        ...
    );
    
    CREATE TABLE constellations (
        constellation_id INTEGER,
        ...
        PRIMARY KEY (constellation_id),
        ...
    );
    
    CREATE TABLE stars (
        star_id INTEGER,
        ...
        PRIMARY KEY (star_id),
        ...
    );

because in SQLite, if a column is defined as `INTEGER`
and `PRIMARY KEY`, there is no need to
define it as `AUTO_INCREMENT`.

### The Application

We setup our Go workspace first with `go mod init starzz-gin`.

All code committed at each chapter is available with the commit message of the chapter name.

<img src="assets/commits.jpg" width="300" height="300"/>

#### Chapter 1: Setting up the routes

Go libraries added:

    gin

We set up the project structure as follows:

    assets           -> contains the application's assets
    main.go          -> the main module for the application
    |_ controllers   -> modules to handle application logic
    |_ routers       -> modules to handle application requests

The module `main.go` contains code to dispatch the requests to the application, to the modules
in `routers`:

    ...
    func main() {
        router := gin.Default()

        {
            constellations := router.Group("/constellations")
            constellations.GET("/", routers.HandleListConstellations)
            constellations.POST("/", routers.HandleRegisterConstellation)
            constellations.GET("/:id", routers.HandleGetConstellationByID)
            constellations.PUT("/:id", routers.HandleUpdateConstellationByID)
            constellations.DELETE("/:id", routers.HandleDeleteConstellationByID)
        }

        ...

        router.Run("localhost:8080")
    }
    ...

Notice that we assign the routes to groups, to emphasize that they are closely related.

A sample module in the `routers` package is `constellations.go`.  It contains the functions to
handle the requests to the `/constellations` endpoints:

    ...
    func HandleListConstellations(c *gin.Context) {
        statusCode, message := controllers.ListConstellations()
        c.JSON(statusCode, message)
    }

    func HandleRegisterConstellation(c *gin.Context) {
        var newData constellationData

        if err := c.BindJSON(&newData); err != nil {
            // if the conversion fails, this will automatically return HTTP 400
            // so there is no need to explicitly handle it
            return
        }

        statusCode, message := controllers.RegisterConstellation(&newData)
        c.JSON(statusCode, message)
    }
    ...

Note the `c.BindJSON(&newData)` statement.  It transforms the JSON data sent by the client as part of the request,
to a Go `struct` that is defined as:

    type constellationData struct {
        ConstellationID   int    `json:"constellation_id"`
        ConstellationName string `json:"constellation_name"`
        GalaxyID          int    `json:"galaxy_id"`
        AddedBy           int    `json:"added_by"`
        VerifiedBy        int    `json:"verified_by"`
    }

For example, the JSON data:

    {
        "constellation_id": 17,
        "constellation_name": "Orion",
        "galaxy_id": 54,
        "added_by": 100,
        "verified_by": 2
    }

is converted to:

    {
        ConstellationID: 17,
        ConstellationName: "Orion",
        GalaxyID: 54,
        AddedBy: 100,
        VerifiedBy: 2
    }

The different requests are then forwarded to different functions in `constellations.go` in the `controllers` package.

    ...
    func ListConstellations() (int, map[string]string) {
        return http.StatusOK, map[string]string{"message": "Successfully listed constellations."}
    }

    func RegisterConstellation(newData any) (int, map[string]any) {
        return http.StatusCreated, map[string]any{"message": "Successfully added constellation.", "data": newData}
    }
    ...

The other modules in the `routers` package follow a similar logic.

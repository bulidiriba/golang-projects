## Modules Used 

To install use

    go get <module-location>

If there is no go.mod in the current directory, don't for get to create/inialize it with the ff cmd

    go mod init "github.com/bulidiriba"  or
    go mod init "example.com/m"

And after installing/downloading them we need to import in the file we need them

    import ( "module-name" )

1. `gorilla-mux`

        go get "github.com/gorilla/mux"

    but since we are outside of the module, we will use `go install` with the version, mainly `latest` 

        go install "github.com/gorilla/mux@latest"

    To import

        import ( "github.com/gorilla/mux" )


## Library Used

They are internal library, so we can import them
They are simply imported with `import` keyword

    import ( "library" )

1. `encoding/json` library is used to encode data into json format

2. `math/rand` -- used to create random ID for created movie

3. `net/http` -- used to create server in golang

4. `strconv` -- used to convert string into integer

5. 



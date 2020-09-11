# Learnings

- Use `go mod init <module-name>` to initialize a module, allowing you to import packages within the module
- Use `go mod vendor` to download dependencies listed in `go.mod`. This downloads the packages in to a `vendor` directory, which has precedence over `$GOPATH/src/`
- make() is used to initialize objects like arrays and maps. Need to pass in size when making arrays, but not maps
- To pass objects into a http handler function, write a wrapper:

    ```
    func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
        return func (w http.ResponseWriter, r *http.Request) {
            url, ok := pathsToUrls[r.URL.Path]
            if ok {
                http.Redirect(w, r, url, 301)
            } else {
                fallback.ServeHTTP(w, r)
            }
        }
    }

    http.ListenAndServe(":8080", MapHandler)

    ```

    HTTP handler now has access to parameters like pathsToUrls
- Parsing YAML requires structure to be defined
    ```
    yamlStr := `
    - path: /urlshort
    url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
    url: https://github.com/gophercises/urlshort/tree/solution
    `

    type yamlStruct struct {
        Entries []map[string]string  `yaml:"entries"`
    }

    var parsedYAML []map[string]string
    err := yaml.Unmarshal(yamlStr, &parsedYAML)
    ```

    OR

    ```
    type entryStruct struct {
        Path string
	    URL string`
    }

    type yamlStruct struct {
        Entries []entryStruct
    }

    var parsedYAML []entryStruct
    err := yaml.Unmarshal(yamlStr, &parsedYAML)
    ```

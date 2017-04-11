# geb - An ALI example

- Project layout
  - The `geb.yaml` project file
- Starting an API
- Adding types
- Adding databases
- Adding routes and resources
- Building, running, and using
- Adding functionality to the server

## Starting the API

### The design layout

The typical layout for an API project is defined by the design layout.
It is more complex that the typical CLI.
Eventually, many of these sub-designs will be generated
from the API design itself. Stay tuned!

API:

```
my-project/

    geb.yaml
    Makefile   (for simple command sequences)

    design/

        server/
            api.yaml
            routes/
                <route-name>.yaml
                ...
                <route-name>/               (when it has subroutes)
                    <sub-route-name>.yaml
                    ...
            resources/
                <resource-name>.yaml
                ...
                <resource-name>/            (when it has subresources)
                    <sub-resource-name>.yaml
                    ...

        lib/
            app-specific/
                pkg.yaml
                ...
                subdir/type.yaml
                ...
            types/
                type.yaml
                ...
                subdir/type.yaml
```


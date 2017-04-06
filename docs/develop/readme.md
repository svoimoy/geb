# Developing

The docs are pretty much m.i.a. but you should still dive in!

Run: `geb sys dev` to enter development mode.

### Overriding generators

To override a generator, in any project directory:

Setup:

1. make a dir `dsl`
1. add sub-dirs to `dsl`
  1. match the directory structure of the library dsls
  1. Be sure to include the `geb-dsl.yaml` and `geb-gen.yaml` files
1. make sure the `dsl-config.paths` has the correct entries.

To __override__, add the file
to the expected location.

To __extend__ the templates,
add any new files you wish to.
Repeated templates also need
to be associated in the `geb-gen.yaml` file.

### Extending dsls

DSLs are also easily extended.
Since both dsl and design are data
both are quite flexible together,
and combine for simple extensibility
and customization
(with the ability to override generators as well).

To extend a dsl:

1. Add any additional data to the design file, in the dsl specs.
1. Reference the new fields in your overriden generators.

Eventually DSLs will become more formal
and the documentation will also
become more detailed on the
dsl specifics (types too)
 






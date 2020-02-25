# import-submodule

Testing how to use [multiple modules of a single repository](https://github.com/vikyd/submodule01).

# Run

```sh
git clone git@github.com:vikyd/import-submodule.git

cd import-submodule

go run main.go
```

prints:

```
F01
F01 v2
F02 in module: github.com/vikyd/submodule01/module01
F02 in module: github.com/vikyd/submodule01/module01/v2 , yeah v2
```

# blipblop

## Requirements
1. Make sure Go is installed (it's super easy)
2. done

If you haven't used Go before, welcome to the simple life...

> Mixing languages is better than writing everything in one, if and only if using only that one is likely to overcomplicate the program - Eric Steven Raymond

## Go Modules

This is a monorepo.  The **core** module is where all the 'core' Go data is stored, this should not depend on any other package.  **app** stores our frontends.

## Contributing
Make your change in a feature branch i.e. `scraper/feature/apify-concurrent`

Make sure you're code is readable and you have `go fmt` your files before pushing.

Then submit a PR which must get approval from at least one other dev.


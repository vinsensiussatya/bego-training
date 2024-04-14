![technology:go](https://img.shields.io/badge/technology-go-blue.svg)
# Bego Training

Just training on Go Backend

## How do I get set up? ###

1. Install go 1.21 or later
2. Prepare dependencies on your local
3. Copy .env.example to .env; then set the value as needed especially for dependencies.
   * Or you can set up vault on your local, and follow the keys from .env.example and set the VAULT_ENABLED = true
4. Run it with `go build && ./bego-training [ARGS]`
   * or alternatively `make run ARGS=[ARGS]`
   * see more command with `go build && ./bego-training --help`

### How to unit test
- Install mockery `brew install mockery`
- Generate mocks `make mocks`
- Run test: `make test`

## Contribution guidelines
- Install [staticcheck](https://staticcheck.io/docs/getting-started/) for linter.
- Using Trunk Based Development. We only have master as our main branch. Always sync with it.

  - Create branch from master, do the work, then PR with master as target.
  - Build the unit test is mandatory.
  - run `make prepare` before create PR.
  - Ask your teammates to do the code review, if at least 2 approved, merge the PR.
  - Good practice for commit message: use [conventional commit](https://www.conventionalcommits.org)

### Architecture / Project Structure
We use clean architecture in this project, so we implement a strict separation of concerns and dependency rule.
Read: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

We have three layers in this project:
#### Presentation
Please reduce any logical code especially the complex one in this layer.

You can put your:

- Rest API (Handler)
- Consumer
- Scheduler
- GraphQL
- etc

#### [TO BE ADDED] Data Layer (Repository)
This is a layer that will contains the code that communicate with the outside of our application/business logic. Generally we don't write much code in this layer. Please reduce any logical code especially the complex one in this layer.

You can put your:

- Database Interaction
- 3rd Party Communication
- Metric

#### Domain / Business Layer (Service)
Where you can put your complex logic that related to our business logic. Since we don't want this layer to have any dependencies to the outside, we put the repository interfaces in this layer and make the data layer dependent to them instead.


### Structs
Each layer have their own struct, in this case:

- Data Layer -> Model/Table struct
- Presentation Layer -> Payload/Request/Response/Message struct
- Domain Layer -> You can create the struct according to the needs

If a functions have more than three parameters you can consider to create a struct as the parameter. So it will increase our code readability.

## Docs
see documentation [here](docs/)


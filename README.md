# Graphql Federation Example

Complete example of Graphql Federation using NodeJS, Java and Go.

## Author project

Subgraph for author using Go lang using gqlgen

To run this project use the command

```
/author> go run server.go
```

This will run the graphql subgraph on `http://localhost:8081/query`

## Book project

Subgraph for author using Java lang, Spring boot using Netflix DGS

To run this project use the command

```
/book> ./gradlew build clean
/book> ./gradlew bootRun
```

This will run the graphql subgraph on `http://localhost:8080/graphql`

## Gateway project

Supergraph Gateway build on Apollo Server using Introspect and Compose in Node lang

To run this project use the command

```
/gateway> npm run dev
```

This will run the studio on `http://localhost:4000/`

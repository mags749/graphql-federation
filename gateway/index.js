import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: 'books', url: 'http://localhost:8080/graphql'},
      { name: 'authors', url: 'http://localhost:8081/query'}
    ]
  }),
});

const server = new ApolloServer({ gateway });

const { url } = await startStandaloneServer(server);

console.log(`Server ready at ${url}`);
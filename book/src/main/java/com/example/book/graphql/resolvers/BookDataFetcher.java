package com.example.book.graphql.resolvers;

import java.util.List;

import com.example.book.graphql.models.Book;
import com.netflix.graphql.dgs.DgsComponent;
import com.netflix.graphql.dgs.DgsQuery;

@DgsComponent
public class BookDataFetcher {
  private final List<Book> bookList = List.of(
    new Book(1, "First book", 1995),
    new Book(2, "Second book", 1998),
    new Book(3, "Third book", 1999)
  );

  @DgsQuery
  private List<Book> books() {
    return this.bookList;
  }
}

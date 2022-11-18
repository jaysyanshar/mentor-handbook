import { useParams } from "react-router-dom";

import { BookDetailCard } from "../cards/bookDetail";

import booksData from "../data/books.json"

export function BookDetailRoute() {
    const { isbn } = useParams();

    const bookMatched = (book) => {
        return book.isbn === isbn;
    }

    const book = booksData.find(bookMatched)

    return (
        <BookDetailCard book={book}></BookDetailCard>
    )
}

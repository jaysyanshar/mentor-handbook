import { BookCard } from "../cards/book"

import { Link, Outlet } from "react-router-dom"

import booksData from "../data/books.json"

export function BookRoute() {
    return (
        <>
            <nav>
                <Link to="/">Home</Link> | Books List
            </nav>
            <h1>
                List of Books
            </h1>
            <div>
                {booksData.map((book, index) => (
                    <BookCard position={index + 1} bookData={book} key={book.isbn}></BookCard>  
                ))}
            </div>

            <hr />
            <Outlet data={booksData} />
        </>
    )
}

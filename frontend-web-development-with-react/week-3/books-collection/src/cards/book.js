import { Link } from "react-router-dom"

export function BookCard(props) {
    return (
        <div>
            <h2>
                {props.position}. <Link to={props.bookData.isbn}>{props.bookData.title}</Link>
            </h2>
        </div>
    )
}

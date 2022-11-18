import { Link } from "react-router-dom"

export function HomeRoute() {
    return (
        <>
            <h1>
                Welcome to my Books Collection
            </h1>
            <h2>
                <Link to="/book">Enter</Link>
            </h2>
        </>
    )
}

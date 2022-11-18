export function BookDetailCard(props) {
    return (
        <div style={{"text-align":"left",padding:"1em"}}>
            <h3>
                Title: {props.book.title}
            </h3>
            <h4>
                Description:<br/>
                {props.book.description}
            </h4>
        </div>
    )
}

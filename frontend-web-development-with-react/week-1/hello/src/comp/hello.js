import { useState } from "react";

function Hello(props) {
    const [counter, setCounter] = useState("a little");

    return <div>
        <h1>I am {props.name}</h1>
        <p>I like Javascript {counter}</p>
        <button onClick={() => setCounter("a lot")}>Change</button>
    </div>
}

export default Hello
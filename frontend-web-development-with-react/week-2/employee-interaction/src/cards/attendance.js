function AttendanceButton(props) {
    return (
        <button onClick={(e) => props.changePresent(e.target.textContent)}>
            {props.present}
        </button> 
    )
}

export default AttendanceButton
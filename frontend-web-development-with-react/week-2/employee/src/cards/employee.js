function EmployeeCard(props) {
    let isOdd = true;
    if (props.arrIndex % 2 === 0) {
        isOdd = false;
    }

    return isOdd ? 
    (
        <div style={{'background-color':'lightgray'}}>
            {props.name}
        </div> 
     ) : (
        <div style={{'background-color':'white'}}>
            {props.name}
        </div>
     )
}

export default EmployeeCard
import { useState } from "react";

import AttendanceButton from "./attendance";

function EmployeeCard(props) {
    const [present, setPresent] = useState("Absent")

    const changePresent = (present) => {
        if (present === "Present") {
            setPresent("Absent")
        }
        else {
            setPresent("Present")
        }
    }

    let isOdd = true;
    if (props.arrIndex % 2 === 0) {
        isOdd = false;
    }

    return isOdd ? 
    (
        <div style={{'background-color':'lightgray'}}>
            {props.name}
            <AttendanceButton present={present} changePresent={changePresent}></AttendanceButton>
        </div> 
     ) : (
        <div style={{'background-color':'white'}}>
            {props.name}
            <AttendanceButton present={present} changePresent={changePresent}></AttendanceButton>
        </div>
     )
}

export default EmployeeCard
import logo from './logo.svg';
import './App.css';

import EmployeeCard from './cards/employee.js';


let employeeData = [
  {
    "id": 1,
    "name": "Jason"
  },
  {
    "id": 2,
    "name":"Edwin"
  },
  {
    "id": 3,
    "name": "Jack"
  },
  {
    "id": 4,
    "name": "Jane"
  },
  {
    "id": 5,
    "name":"Jill"
  }]

function App() {
  return (
    <div>
      {employeeData.map((employee, index) => (
        <EmployeeCard name={employee.name} key={employee.id} arrIndex={index}></EmployeeCard>  
      ))}
    </div>
  );
}

export default App;

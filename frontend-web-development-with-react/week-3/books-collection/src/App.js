import logo from './logo.svg';
import { Routes, Route } from "react-router-dom";
import './App.css';
import { HomeRoute } from './routes/home';
import { BookRoute } from './routes/book';
import { BookDetailRoute } from './routes/bookDetail';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<HomeRoute />} />
        <Route path="book" element={<BookRoute />} >
          <Route path=":isbn" element={<BookDetailRoute />} />
        </Route>
      </Routes>
    </div>
  );
}

export default App;

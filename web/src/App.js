import React from 'react';
import './App.css';
import {AddTodo, TodoList} from './components/todo';

function App() {
  return (
      <div className="App">
        <header className="App-header">
          <AddTodo/>
          <TodoList/>
        </header>
      </div>
  );
}

export default App;

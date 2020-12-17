import React from 'react';
import './app.css';
import {connect} from 'react-redux';
import {AddTodo} from '../addTask';
import {ListTodo} from '../listTasks';

class App extends React.Component {
  render() {
    return (
        <div className="App">
          <AddTodo/>
          <ListTodo/>
        </div>
    );
  }
}

export default App;

function mapStateToProps(state) {
  return {};
}

const actionCreators = {};

const connectedApp = connect(mapStateToProps, actionCreators)(App);
export {connectedApp as App};

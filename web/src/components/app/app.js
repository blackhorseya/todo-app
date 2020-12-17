import React from 'react';
import './app.css';
import {connect} from 'react-redux';
import {AddTodo} from '../addTask';
import {ListTodo} from '../listTasks';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';

class App extends React.Component {
  render() {
    return (
        <React.Fragment>
          <CssBaseline/>
          <Container maxWidth="sm">
            <AddTodo/>
            <ListTodo/>
          </Container>
        </React.Fragment>
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

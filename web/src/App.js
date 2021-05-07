import React from 'react';
import './App.css';
import {AddTodo, TodoList} from './components/todo';
import {Container, Grid} from '@material-ui/core';

function App() {
  return (
      <Container maxWidth={'xs'} style={{padding: 20}}>
        <Grid container justify="center" alignItems="center" spacing={3}>
          <Grid item xs={12}>
            <AddTodo/>
          </Grid>
          <Grid item xs={12}>
            <TodoList/>
          </Grid>
        </Grid>
      </Container>
  );
}

export default App;

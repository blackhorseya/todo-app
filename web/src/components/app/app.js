import React from 'react';
import './app.css';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';

class App extends React.Component {
  componentDidMount() {
    this.props.list();
  }

  render() {
    const {tasks} = this.props;

    return (
        <div className="App">
          <ul>
            {
              tasks.item && tasks.item.map((task, index) =>
                  <li key={index}>
                    {task.title}
                  </li>,
              )
            }
          </ul>
        </div>
    );
  }
}

export default App;

function mapStateToProps(state) {
  const {tasks} = state;
  return {tasks};
}

const actionCreators = {
  list: taskActions.list,
};

const connectedApp = connect(mapStateToProps, actionCreators)(App);
export {connectedApp as App};

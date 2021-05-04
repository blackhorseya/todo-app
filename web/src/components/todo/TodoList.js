import React from 'react';
import {todoActions} from '../../_actions';
import {connect} from 'react-redux';

class TodoList extends React.Component {
  componentDidMount() {
    this.props.list(0, 10);
  }

  render() {
    const {todo} = this.props;

    return (
        <div>
          {todo.loading ? <h1>Loading...</h1> : <h1>Todo List</h1>}
          {todo.loading === false && <ul>
            {todo.data.map((item, i) =>
                <li key={item.id}>{item.title}</li>,
            )}
          </ul>}
        </div>
    );
  }

}

function mapStateToProps(state) {
  const {todo} = state;
  return {todo};
}

const actionCreators = {
  list: todoActions.list,
};

const connectedTodoList = connect(mapStateToProps, actionCreators)(TodoList);
export {connectedTodoList as TodoList};
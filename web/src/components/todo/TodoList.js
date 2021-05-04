import React from 'react';
import {todoActions} from '../../_actions';
import {connect} from 'react-redux';

class TodoList extends React.Component {
  constructor(props) {
    super(props);

    this.handleRemove = this.handleRemove.bind(this);
  }

  componentDidMount() {
    this.props.list(0, 10);
  }

  handleRemove(id) {
    if (id) {
      this.props.remove(id);
    }
  }

  render() {
    const {todo} = this.props;

    return (
        <div>
          {todo.loading ? <h1>Loading...</h1> : <h1>Todo List</h1>}
          {todo.loading === false && <ul>
            {todo.data.map((item, _) =>
                <li key={item.id}>{item.title}
                  <button onClick={() => this.handleRemove(item.id)}>remove
                  </button>
                </li>,
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
  remove: todoActions.remove,
};

const connectedTodoList = connect(mapStateToProps, actionCreators)(TodoList);
export {connectedTodoList as TodoList};
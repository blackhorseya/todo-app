import React from 'react';
import {todoActions} from '../../_actions';
import {connect} from 'react-redux';

class TodoList extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      start: 0,
      end: 4,
      size: 5,
    };

    this.handleRemove = this.handleRemove.bind(this);
    this.handleChangeStatus = this.handleChangeStatus.bind(this);
    this.handleChangeSize = this.handleChangeSize.bind(this);
    this.handlePrevious = this.handlePrevious.bind(this);
    this.handleNext = this.handleNext.bind(this);
  }

  componentDidMount() {
    const {start, end} = this.state;
    this.props.list(start, end);
  }

  handleChangeStatus(id, status) {
    if (id) {
      this.props.changeStatus(id, !status);
    }
  }

  handleRemove(id) {
    if (id) {
      this.props.remove(id);
    }
  }

  handleChangeSize(e) {
    const {value} = e.target;

    const {start, size} = this.state;
    const newEnd = start + size;
    const newSize = parseInt(value, 10);
    this.setState({end: newEnd, size: newSize});

    this.props.list(start, newEnd);
  }

  handlePrevious(e) {
    const {start, end, size} = this.state;
    const newStart = start - size;
    const newEnd = end - size;

    this.setState({start: newStart, end: newEnd});

    this.props.list(newStart, newEnd);
  }

  handleNext(e) {
    const {start, end, size} = this.state;
    const newStart = start + size;
    const newEnd = end + size;

    this.setState({start: newStart, end: newEnd});

    this.props.list(newStart, newEnd);
  }

  render() {
    const {size} = this.state;
    const {todo} = this.props;

    return (
        <div>
          {todo.loading ? <h1>Loading...</h1> : <div>
            <h1>Todo List</h1>
            <button onClick={this.handlePrevious}>{`<`}</button>
            <input type="number" name="size" value={size}
                   onChange={this.handleChangeSize}/>
            <button onClick={this.handleNext}>{`>`}</button>
          </div>}
          {todo.loading === false && todo.data && <ul>
            {todo.data.map((item, _) =>
                <li key={item.id}>
                  <button onClick={() => this.handleChangeStatus(
                      item.id, item.completed)}>{item.completed ?
                      'X' :
                      'V'}</button>
                  {item.title}
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
  changeStatus: todoActions.changeStatus,
};

const connectedTodoList = connect(mapStateToProps, actionCreators)(TodoList);
export {connectedTodoList as TodoList};
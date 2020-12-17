import React from 'react';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';

class AddTask extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      title: '',
      submitted: false,
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(e) {
    const {name, value} = e.target;
    this.setState({[name]: value});
  }

  handleSubmit(e) {
    e.preventDefault();

    this.setState({submitted: true});
    const {title} = this.state;
    this.props.add(title);
  }

  render() {
    const {title} = this.state;

    return (
        <div>
          <form onSubmit={this.handleSubmit}>
            <input name="title" value={title} onChange={this.handleChange}
                   placeholder="Input title..."/>
            <button type="submit">Add Task</button>
          </form>
        </div>
    );
  }
}

function mapStateToProps(state) {
  const {tasks} = state;
  return {tasks};
}

const actionCreators = {
  add: taskActions.add,
};

const connectedAddTodo = connect(mapStateToProps, actionCreators)(AddTask);
export {connectedAddTodo as AddTodo};

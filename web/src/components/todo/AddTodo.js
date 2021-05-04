import React from 'react';
import {connect} from 'react-redux';
import {todoActions} from '../../_actions';

class AddTodo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      title: '',
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

    const {title} = this.state;
    if (title) {
      this.props.add({title: title});
    }
    this.setState({'title': ''});
  }

  render() {
    const {title} = this.state;

    return (
        <div>
          <h3>Add Todo</h3>
          <input type="text" name="title" value={title}
                 onChange={this.handleChange}/>
          <button onClick={this.handleSubmit}>add</button>
        </div>
    );
  }
}

function mapStateToProps(state) {
  const {todo} = state;
  return {todo};
}

const actionCreators = {
  add: todoActions.add,
};

const connectedAddTodo = connect(mapStateToProps, actionCreators)(AddTodo);
export {connectedAddTodo as AddTodo};
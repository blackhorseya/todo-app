import React from 'react';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';
import TextField from '@material-ui/core/TextField';
import {Button} from '@material-ui/core';

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
    if (title !== '') {
      this.props.add(title);
      this.setState({'title': ''});
    }
  }

  render() {
    const {title} = this.state;

    return (
        <form onSubmit={this.handleSubmit}>
          <div>
            <TextField name="title" value={title} onChange={this.handleChange}
                       variant="outlined" placeholder="Title..."/>
            <Button type="submit" variant="contained"
                    color="primary">Add</Button>
          </div>
        </form>
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

const connectedAddTask = connect(mapStateToProps, actionCreators)(AddTask);
export {connectedAddTask as AddTask};

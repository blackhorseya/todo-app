import React from 'react';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';
import Title from '../_shared/title';
import {Table} from '@material-ui/core';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import TableBody from '@material-ui/core/TableBody';
import Checkbox from '@material-ui/core/Checkbox';

class ListTasks extends React.Component {
  componentDidMount() {
    this.props.list();
  }

  render() {
    const {tasks} = this.props;

    return (
        <React.Fragment>
          <Title>Tasks</Title>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell>Completed</TableCell>
                <TableCell>Title</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {tasks.item && tasks.item.map((row, index) => (
                  <TableRow key={index}>
                    <TableCell>{row.id}</TableCell>
                    <TableCell>
                      <Checkbox checked={row.completed}
                                inputProps={{'aria-label': 'primary checkbox'}}/>
                    </TableCell>
                    <TableCell>{row.title}</TableCell>
                  </TableRow>
              ))}
            </TableBody>
          </Table>
        </React.Fragment>
    );
  }
}

function mapStateToProps(state) {
  const {tasks} = state;
  return {tasks};
}

const actionCreators = {
  list: taskActions.list,
};

const connectedListTodo = connect(mapStateToProps, actionCreators)(ListTasks);
export {connectedListTodo as ListTodo};

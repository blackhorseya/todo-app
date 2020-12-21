import React from 'react';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';
import {
  IconButton,
  Paper,
  Table,
  TableContainer,
} from '@material-ui/core';
import DeleteIcon from '@material-ui/icons/Delete';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import TableBody from '@material-ui/core/TableBody';
import TableFooter from '@material-ui/core/TableFooter';
import TablePagination from '@material-ui/core/TablePagination';

class ListTasks extends React.Component {
  componentDidMount() {
    this.props.list();

    this.handleFirstPage = this.handleFirstPage.bind(this);
    this.handleLastPage = this.handleLastPage.bind(this);
    this.handleNextPage = this.handleNextPage.bind(this);
    this.handleBackPage = this.handleBackPage.bind(this);
    this.handleRemoveTask = this.handleRemoveTask.bind(this);
    this.handleChangeStatus = this.handleChangeStatus.bind(this);
  }

  handleRemoveTask(id) {
    this.props.remove(id);
  }

  handleChangeStatus(id, completed) {
    if (completed === undefined || !completed) {
      this.props.changeStatus(id, 2)
    } else {
      this.props.changeStatus(id, 1)
    }
  }

  handleFirstPage(e) {
    console.log('first page');
  }

  handleLastPage(e) {
    console.log('last page');
  }

  handleNextPage(e) {
    console.log('next page');
  }

  handleBackPage(e) {
    console.log('back page');
  }

  render() {
    const {tasks} = this.props;

    return (
        <TableContainer component={Paper}>
          <Table size="small">
            <TableBody>
              {tasks.item && tasks.item.map((row) => (
                  <TableRow key={row.id}>
                    <TableCell>
                      <IconButton
                          onClick={() => this.handleChangeStatus(row.id,
                              row.completed)}>
                        <span className="material-icons">
                          {row.completed
                              ? 'task_alt'
                              : 'radio_button_unchecked'}
                        </span>
                      </IconButton>
                    </TableCell>
                    <TableCell>{row.title}</TableCell>
                    <TableCell>
                      <IconButton onClick={() => this.handleRemoveTask(row.id)}
                                  edge="end"
                                  aria-label="delete">
                        <DeleteIcon/>
                      </IconButton>
                    </TableCell>
                  </TableRow>
              ))}
            </TableBody>
            {/*<TableFooter>*/}
            {/*  <TableRow>*/}
            {/*    <TablePagination rowsPerPageOptions={[10, 50]}/>*/}
            {/*  </TableRow>*/}
            {/*</TableFooter>*/}
          </Table>
        </TableContainer>
    );
  }
}

function mapStateToProps(state) {
  const {tasks} = state;
  return {tasks};
}

const actionCreators = {
  list: taskActions.list,
  remove: taskActions.remove,
  changeStatus: taskActions.changeStatus,
};

const connectedListTasks = connect(mapStateToProps, actionCreators)(ListTasks);
export {connectedListTasks as ListTasks};

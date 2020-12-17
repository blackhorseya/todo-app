import React from 'react';
import {taskActions} from '../../_actions';
import {connect} from 'react-redux';
import {Paper, Table, TableContainer} from '@material-ui/core';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import TableBody from '@material-ui/core/TableBody';
import Checkbox from '@material-ui/core/Checkbox';
import TableFooter from '@material-ui/core/TableFooter';
import TablePagination from '@material-ui/core/TablePagination';

class ListTasks extends React.Component {
  componentDidMount() {
    this.props.list();

    this.handleFirstPage = this.handleFirstPage.bind(this);
    this.handleLastPage = this.handleLastPage.bind(this);
    this.handleNextPage = this.handleNextPage.bind(this);
    this.handleBackPage = this.handleBackPage.bind(this);
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
                    <TableCell component="th"
                               scope="row">{row.title}</TableCell>
                  </TableRow>
              ))}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TablePagination rowsPerPageOptions={[10, 50]}/>
              </TableRow>
            </TableFooter>
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
};

const connectedListTodo = connect(mapStateToProps, actionCreators)(ListTasks);
export {connectedListTodo as ListTodo};

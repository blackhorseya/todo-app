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
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import TableBody from '@material-ui/core/TableBody';
import TableFooter from '@material-ui/core/TableFooter';
import TablePagination from '@material-ui/core/TablePagination';

class ListTasks extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      page: 0,
      size: 5,
    };

    this.handleRemoveTask = this.handleRemoveTask.bind(this);
    this.handleChangeStatus = this.handleChangeStatus.bind(this);

    this.handleFirstPage = this.handleFirstPage.bind(this);
    this.handleLastPage = this.handleLastPage.bind(this);
    this.handleNextPage = this.handleNextPage.bind(this);
    this.handleBackPage = this.handleBackPage.bind(this);

    this.handleChangePage = this.handleChangePage.bind(this);
    this.handleChangeSize = this.handleChangeSize.bind(this);
  }

  componentDidMount() {
    const {page, size} = this.state;
    this.props.list(page + 1, size);
  }

  handleRemoveTask(id) {
    this.props.remove(id);
  }

  handleChangeStatus(id, completed) {
    if (completed === undefined || !completed) {
      this.props.changeStatus(id, 2);
    } else {
      this.props.changeStatus(id, 1);
    }
  }

  handleChangePage(e, page) {
    this.setState({'page': page});
    const {size} = this.state;

    this.props.list(page + 1, size);
  }

  handleChangeSize(e) {
    this.setState({'size': parseInt(e.target.value, 10)});
    this.setState({'page': 0});

    this.props.list(1, parseInt(e.target.value, 10));
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
    const {page, size} = this.state;
    const {tasks} = this.props;

    return (
        <TableContainer component={Paper}>
          <Table size="small">
            <TableBody>
              {tasks.item && tasks.item.data.map((row) => (
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
            <TableFooter>
              <TableRow>
                <TablePagination
                    rowsPerPageOptions={[5, 10, 25, {label: 'All', value: -1}]}
                    colSpan={3}
                    count={tasks.item !== undefined
                        ? tasks.item.total
                        : 0}
                    rowsPerPage={size}
                    page={tasks.item !== undefined ? page : 0}
                    SelectProps={{
                      inputProps: {'aria-label': 'rows per page'},
                      native: true,
                    }}
                    onChangePage={this.handleChangePage}
                    onChangeRowsPerPage={this.handleChangeSize}
                />
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
  remove: taskActions.remove,
  changeStatus: taskActions.changeStatus,
};

const connectedListTasks = connect(mapStateToProps, actionCreators)(ListTasks);
export {connectedListTasks as ListTasks};

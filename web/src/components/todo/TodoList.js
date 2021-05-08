import React from 'react';
import {todoActions} from '../../_actions';
import {connect} from 'react-redux';
import {
  Grid,
  IconButton,
  List,
  ListItem,
  ListItemIcon,
  ListItemSecondaryAction,
  ListItemText,
  Paper,
  TablePagination,
} from '@material-ui/core';
import {Check, Close, Delete} from '@material-ui/icons';

class TodoList extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      page: 0,
      size: 5,
    };

    this.handleRemove = this.handleRemove.bind(this);
    this.handleChangeStatus = this.handleChangeStatus.bind(this);
    this.handleChangeRowsPerPage = this.handleChangeRowsPerPage.bind(this);
    this.handleChangePage = this.handleChangePage.bind(this);
  }

  componentDidMount() {
    const {page, size} = this.state;
    const start = page * size;
    const end = (page + 1) * size - 1;

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

  handleChangePage(e, newPage) {
    const {size} = this.state;
    this.setState({page: newPage});
    const start = newPage * size;
    const end = (newPage + 1) * size - 1;

    this.props.list(start, end);
  }

  handleChangeRowsPerPage(e) {
    const newSize = parseInt(e.target.value, 10);
    this.setState({size: newSize});
    const start = 0;
    const end = newSize - 1;

    this.props.list(start, end);
  }

  render() {
    const {size, page} = this.state;
    const {todo} = this.props;

    return (
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Paper style={{padding: 20}}>
              <Grid container spacing={2}>
                <Grid item xs={12}>
                  <TablePagination
                      component="div"
                      count={parseInt(todo.total, 10)}
                      page={page}
                      onChangePage={this.handleChangePage}
                      rowsPerPage={size}
                      rowsPerPageOptions={[5, 10, 20, 50]}
                      onChangeRowsPerPage={this.handleChangeRowsPerPage}
                  />
                </Grid>
                <Grid item xs={12}>
                  {todo.loading ? <h1>Loading...</h1> : todo.data && <List>
                    {todo.data.map((item, _) =>
                        <ListItem key={item.id} role={undefined} dense button
                                  onClick={() => this.handleChangeStatus(
                                      item.id, item.completed)}>
                          <ListItemIcon>
                            {item.completed ? <Close/> : <Check/>}
                          </ListItemIcon>
                          <ListItemText id={item.id} primary={item.title}/>
                          <ListItemSecondaryAction>
                            <IconButton edge="end"
                                        onClick={() => this.handleRemove(
                                            item.id)}>
                              <Delete/>
                            </IconButton>
                          </ListItemSecondaryAction>
                        </ListItem>,
                    )}
                  </List>}
                </Grid>
              </Grid>
            </Paper>
          </Grid>
        </Grid>
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
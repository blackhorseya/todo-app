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
} from '@material-ui/core';
import {Check, Close, Delete} from '@material-ui/icons';

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
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Paper style={{padding: 20}}>
              <Grid container spacing={2}>
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

          {/*{todo.loading ? <h1>Loading...</h1> : <div>*/}
          {/*  <h1>Todo List</h1>*/}
          {/*  <button onClick={this.handlePrevious}>{`<`}</button>*/}
          {/*  <input type="number" name="size" value={size}*/}
          {/*         onChange={this.handleChangeSize}/>*/}
          {/*  <button onClick={this.handleNext}>{`>`}</button>*/}
          {/*  <FormLabel>total: {todo.total}</FormLabel>*/}
          {/*</div>}*/}
          {/*{todo.loading === false && todo.data && <ul>*/}
          {/*  {todo.data.map((item, _) =>*/}
          {/*      <li key={item.id}>*/}
          {/*        <button onClick={() => this.handleChangeStatus(*/}
          {/*            item.id, item.completed)}>{item.completed ?*/}
          {/*            'X' :*/}
          {/*            'V'}</button>*/}
          {/*        {item.title}*/}
          {/*        <button onClick={() => this.handleRemove(item.id)}>remove*/}
          {/*        </button>*/}
          {/*      </li>,*/}
          {/*  )}*/}
          {/*</ul>}*/}
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
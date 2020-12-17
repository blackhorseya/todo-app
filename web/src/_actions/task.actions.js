import {taskConstants} from '../_constants';
import {taskServices} from '../_services';

function list() {
  // action creators
  function request() {return {type: taskConstants.LIST_REQUEST};}

  function success(data) {
    return {
      type: taskConstants.LIST_SUCCESS,
      data,
    };
  }

  function failure(error) {return {type: taskConstants.LIST_FAILURE, error};}

  // actions
  return (dispatch) => {
    dispatch(request());

    // list tasks
    taskServices.list().then(
        (data) => dispatch(success(data)),
        (error) => dispatch(failure(error.toString())),
    );
  };
}

function add(title) {
  // action creators
  function request(task) {return {type: taskConstants.ADD_REQUEST, task};}

  function success(task) {return {type: taskConstants.ADD_SUCCESS, task};}

  function failure(error) {return {type: taskConstants.ADD_FAILURE, error};}

  // actions
  return (dispatch) => {
    dispatch(request({title}));

    // add task
    taskServices.add(title).then(
        (data) => dispatch(success(data)),
        (error) => dispatch(failure(error.toString())),
    );
  };
}

export const taskActions = {
  list,
  add,
};

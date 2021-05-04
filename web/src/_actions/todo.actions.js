import {todoConstants} from '../_constants';
import {todoService} from '../_services';

export const todoActions = {
  list,
  add,
};

function list(start, end) {
  return dispatch => {
    dispatch(request());

    todoService.list(start, end).then(
        tasks => {
          dispatch(success(tasks));
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: todoConstants.LIST_REQUEST};
  }

  function success(tasks) {
    return {type: todoConstants.LIST_SUCCESS, tasks};
  }

  function failure(error) {
    return {type: todoConstants.LIST_FAILURE, error};
  }
}

function add(task) {
  return dispatch => {
    dispatch(request());

    todoService.add(task).then(
        task => {
          dispatch(success(task));
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: todoConstants.ADD_REQUEST};
  }

  function success(task) {
    return {type: todoConstants.ADD_SUCCESS, task};
  }

  function failure(error) {
    return {type: todoConstants.ADD_FAILURE, error};
  }
}
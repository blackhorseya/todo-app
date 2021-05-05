import {todoConstants} from '../_constants';
import {todoService} from '../_services';

export const todoActions = {
  list,
  add,
  remove,
  changeStatus,
};

function list(start, end) {
  return dispatch => {
    dispatch(request());

    todoService.list(start, end).then(
        resp => {
          dispatch(success(resp.data, resp.total));
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: todoConstants.LIST_REQUEST};
  }

  function success(tasks, total) {
    return {type: todoConstants.LIST_SUCCESS, tasks, total};
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

function remove(id) {
  return dispatch => {
    dispatch(request());

    todoService.remove(id).then(
        id => {
          dispatch(success(id));
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: todoConstants.REMOVE_REQUEST};
  }

  function success(id) {
    return {type: todoConstants.REMOVE_SUCCESS, id};
  }

  function failure(error) {
    return {type: todoConstants.REMOVE_FAILURE, error};
  }
}

function changeStatus(id, status) {
  return dispatch => {
    dispatch(request());

    todoService.changeStatus(id, status).then(
        task => {
          dispatch(success(task));
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: todoConstants.CHANGE_STATUS_REQUEST};
  }

  function success(task) {
    return {type: todoConstants.CHANGE_STATUS_SUCCESS, task};
  }

  function failure(error) {
    return {type: todoConstants.CHANGE_STATUS_FAILURE, error};
  }
}

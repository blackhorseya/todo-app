import {taskConstants} from '../_constants';
import {taskServices} from '../_services';
import {func} from 'prop-types';

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

function remove(id) {
  // action creators
  function request(id) {return {type: taskConstants.REMOVE_REQUEST, id};}

  function success(id) {return {type: taskConstants.REMOVE_SUCCESS, id};}

  function failure(error) {return {type: taskConstants.REMOVE_FAILURE, error};}

  // actions
  return (dispatch) => {
    dispatch(request({id}));

    // remove task
    taskServices.remove(id).then(
        (data) => dispatch(success(data)),
        (error) => dispatch(failure(error.toString())),
    );
  };
}

function changeStatus(id, completed) {
  // action creators
  function request(
      id, completed) {
    return {
      type: taskConstants.CHANGE_STATUS_REQUEST,
      id,
      completed,
    };
  }

  function success(
      id, completed) {
    return {
      type: taskConstants.CHANGE_STATUS_SUCCESS,
      id,
      completed,
    };
  }

  function failure(error) {
    return {
      type: taskConstants.CHANGE_STATUS_FAILURE,
      error,
    };
  }

  // actions
  return (dispatch) => {
    dispatch(request(id, completed));

    // remove task
    taskServices.changeStatus(id, completed).then(
        (data) => dispatch(success(id, completed)),
        (error) => dispatch(failure(error.toString())),
    );
  };
}

export const taskActions = {
  list,
  add,
  remove,
  changeStatus,
};

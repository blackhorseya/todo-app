import {taskConstants} from '../_constants';
import {taskServices} from '../_services';

export const taskActions = {
  list,
};

function list() {
  // actions
  return dispatch => {
    dispatch(request());

    // execute task service
    taskServices.list().then(
        data => dispatch(success(data)),
        error => dispatch(failure(error.toString())),
    );
  };

  // action creators
  function request() {return {type: taskConstants.LIST_REQUEST};}

  function success(data) {return {type: taskConstants.LIST_SUCCESS, data: data};}

  function failure(error) {return {type: taskConstants.LIST_FAILURE, error};}
}
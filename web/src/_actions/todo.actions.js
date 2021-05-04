import {todoConstants} from '../_constants';
import {todoService} from '../_services';

export const todoActions = {
  list,
};

function list(start, end) {
  return dispatch => {
    dispatch(request());

    // todo: 2021-05-04|20:51|doggy|implement caller api via services
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
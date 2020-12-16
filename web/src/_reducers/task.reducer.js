import {taskConstants} from '../_constants';

export function tasks(state = {}, action) {
  switch (action.type) {
    case taskConstants.LIST_REQUEST:
      return {
        loading: true,
      };
    case taskConstants.LIST_SUCCESS:
      return {
        item: action.data,
      };
    case taskConstants.LIST_FAILURE:
      return {
        error: action.error,
      };
    default:
      return state;
  }
}
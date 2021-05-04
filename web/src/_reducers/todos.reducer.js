import {todoConstants} from '../_constants';

export function todos(state = {}, action) {
  switch (action.type) {
    case todoConstants.LIST_REQUEST:
      return {
        ...state,
        loading: true,
      };
    case todoConstants.LIST_SUCCESS:
      return {
        loading: false,
        data: action.tasks,
        error: '',
      };
    case todoConstants.LIST_FAILURE:
      return {
        loading: false,
        data: [],
        error: action.error,
      };
    default:
      return state;
  }
}
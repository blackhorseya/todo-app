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
    case taskConstants.ADD_REQUEST:
      return {
        loading: true,
        item: state.item,
      };
    case taskConstants.ADD_SUCCESS:
      return {
        item: {
          ...state.item,
          data: state.item.data.concat(action.task),
          total: state.item.total + 1,
        },
      };
    case taskConstants.ADD_FAILURE:
      return {
        error: action.error,
      };
    case taskConstants.REMOVE_REQUEST:
      return {
        loading: true,
        item: state.item,
      };
    case taskConstants.REMOVE_SUCCESS:
      return {
        item: {
          ...state.item,
          data: state.item.data.filter((task) => task.id !== action.id),
          total: state.item.total - 1,
        },
      };
    case taskConstants.REMOVE_FAILURE:
      return {
        error: action.error,
      };
    case taskConstants.CHANGE_STATUS_REQUEST:
      return {
        item: state.item,
      };
    case taskConstants.CHANGE_STATUS_SUCCESS:
      return {
        item: {
          ...state.item,
          data: state.item.data.map((task) => {
            if (task.id === action.id) {
              return {...task, completed: action.completed === 2};
            } else {
              return task;
            }
          }),
        },
      };
    case taskConstants.CHANGE_STATUS_FAILURE:
      return {
        error: action.error,
      };
    default:
      return state;
  }
}
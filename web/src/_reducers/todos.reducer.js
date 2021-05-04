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

    case todoConstants.ADD_REQUEST:
      return {
        ...state,
        loading: true,
      };
    case todoConstants.ADD_SUCCESS:
      return {
        loading: false,
        data: [...state.data, action.task],
        error: '',
      };
    case todoConstants.ADD_FAILURE:
      return {
        loading: false,
        data: [...state.data],
        error: action.error,
      };

    case todoConstants.REMOVE_REQUEST:
      return {
        ...state,
        loading: true,
      };
    case todoConstants.REMOVE_SUCCESS:
      return {
        loading: false,
        data: [...state.data.filter(x => x.id !== action.id)],
        error: '',
      };
    case todoConstants.REMOVE_FAILURE:
      return {
        loading: false,
        data: [...state.data],
        error: action.error,
      };

    default:
      return state;
  }
}
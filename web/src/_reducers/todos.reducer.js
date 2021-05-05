import {todoConstants} from '../_constants';

const initState = {
  loading: false,
  data: [],
  error: '',
};

export function todos(state = initState, action) {
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

    case todoConstants.CHANGE_STATUS_REQUEST:
      return {
        ...state,
        loading: true,
      };
    case todoConstants.CHANGE_STATUS_SUCCESS:
      const i = state.data.findIndex((o => o.id === action.task.id));

      return {
        loading: false,
        data: [
          ...state.data.slice(0, i),
          action.task,
          ...state.data.slice(i + 1)],
        error: '',
      };
    case todoConstants.CHANGE_STATUS_FAILURE:
      return {
        loading: false,
        data: [...state.data],
        error: action.error,
      };

    default:
      return state;
  }
}
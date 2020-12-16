import {combineReducers} from 'redux';
import {tasks} from './task.reducer';

const rootReducer = combineReducers({
  tasks,
});

export default rootReducer;

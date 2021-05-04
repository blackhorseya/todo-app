import {configureStore} from '@reduxjs/toolkit';
import counterReducer from '../features/counter/counterSlice';
import {todos} from '../_reducers';

export const store = configureStore({
  reducer: {
    todo: todos,
    counter: counterReducer,
  },
});

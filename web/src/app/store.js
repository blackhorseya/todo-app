import {configureStore} from '@reduxjs/toolkit';
import {todos} from '../_reducers';

export const store = configureStore({
  reducer: {
    todo: todos,
  },
});

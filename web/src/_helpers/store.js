import {applyMiddleware, createStore} from 'redux';
import rootReducer from '../_reducers';
import thunkMiddleware from 'redux-thunk';

export const store = createStore(
    rootReducer,
    applyMiddleware(
        thunkMiddleware,
    ),
);
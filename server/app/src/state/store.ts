import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query/react";
import { enableMapSet } from "immer";
import throttle from "lodash/throttle";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { eKonyvApi } from "./api";
import historyReducer from "./historySlice";
import notificationReducer from "./notificationSlice";
import { websocketConn } from "./socketMiddleware";
import stateReducer from "./stateSlice";
import { getWebsocketURL } from "./util";

enableMapSet();

export const store = configureStore({
  reducer: {
    state: stateReducer,
    history: historyReducer,
    notifications: notificationReducer,
    [eKonyvApi.reducerPath]: eKonyvApi.reducer
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
      websocketConn(getWebsocketURL().href),
      eKonyvApi.middleware
    ),
});

setupListeners(store.dispatch);

const saveState = (key: string, state: any): void => {
  try {
    const serialized = JSON.stringify(state);
    localStorage.setItem(key, serialized);
  } catch (err) {}
};

store.subscribe(
  throttle(() => {
    saveState("history", store.getState().history.items);
    saveState("active", store.getState().state.activeItem);
  }, 1000)
);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export const useAppDispatch: () => AppDispatch = useDispatch;
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

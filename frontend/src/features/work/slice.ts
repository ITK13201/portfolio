import { createSlice, Draft, PayloadAction } from '@reduxjs/toolkit';
import { Work } from 'models/work';

type State = Readonly<{
  work?: Work;
}>;

const initialState: State = {};

export const workSlice = createSlice({
  name: 'work',
  initialState,
  reducers: {
    workUpdated: (state: Draft<State>, action: PayloadAction<Work>) => {
      state.work = action.payload;
    },
  },
});

export const workReducer = workSlice.reducer;
export const { workUpdated } = workSlice.actions;

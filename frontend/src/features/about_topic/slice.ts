import { createSlice, Draft, PayloadAction } from '@reduxjs/toolkit';
import { AboutTopic } from 'models/about_topic';

type State = Readonly<{
  aboutTopic?: AboutTopic;
}>;

const initialState: State = {};

export const aboutTopicSlice = createSlice({
  name: 'about_topic',
  initialState,
  reducers: {
    aboutTopicUpdated: (
      state: Draft<State>,
      action: PayloadAction<AboutTopic>
    ) => {
      state.aboutTopic = action.payload;
    },
  },
});

export const aboutTopicReducer = aboutTopicSlice.reducer;
export const { aboutTopicUpdated } = aboutTopicSlice.actions;

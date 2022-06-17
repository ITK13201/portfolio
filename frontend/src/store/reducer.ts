import { combineReducers } from '@reduxjs/toolkit';
import { aboutTopicReducer } from 'features/about_topic/slice';
import { skillReducer } from 'features/skill/slice';
import { workReducer } from 'features/work/slice';

const rootReducer = combineReducers({
  aboutTopic: aboutTopicReducer,
  skill: skillReducer,
  work: workReducer,
});

export default rootReducer;

import { createSlice, Draft, PayloadAction } from '@reduxjs/toolkit';
import { LanguageSkill, TechnologySkill, DatabaseSkill } from 'models/skill';

type State = Readonly<{
  languageSkill?: LanguageSkill;
  technologySkill?: TechnologySkill;
  databaseSkill?: DatabaseSkill;
}>;

const initialState: State = {};

export const skillSlice = createSlice({
  name: 'skill',
  initialState,
  reducers: {
    languageSkillUpdated: (
      state: Draft<State>,
      action: PayloadAction<LanguageSkill>
    ) => {
      state.languageSkill = action.payload;
    },
    technologySkillUpdated: (
      state: Draft<State>,
      action: PayloadAction<TechnologySkill>
    ) => {
      state.technologySkill = action.payload;
    },
    databaseSkillUpdated: (
      state: Draft<State>,
      action: PayloadAction<DatabaseSkill>
    ) => {
      state.databaseSkill = action.payload;
    },
  },
});

export const skillReducer = skillSlice.reducer;
export const {
  languageSkillUpdated,
  technologySkillUpdated,
  databaseSkillUpdated,
} = skillSlice.actions;

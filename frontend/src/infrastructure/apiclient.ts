import axios, { AxiosRequestConfig } from 'axios';
import { AboutTopic } from '../models/about_topic';

export const createApiClient = (config?: AxiosRequestConfig) => {
  const http = axios.create(config);
  return Object.freeze({
    async login(): Promise<void> {
      await http.post('api/v1/auth/login');
    },
    async refreshToken(): Promise<void> {
      await http.post('api/v1/auth/refresh_token');
    },
    async fetchAllAboutTopics(): Promise<AboutTopic[]> {
      return (await http.get('api/v1/about_topic')).data;
    },
  });
};

export type ApiClient = ReturnType<typeof createApiClient>;

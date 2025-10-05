import { CreateClientConfig } from './api/client/types.gen';

export const createClientConfig: CreateClientConfig = (config) => ({
    ...config,
    baseUrl: 'http://localhost:8080/api',
    headers: {
        'X-user-ID': '1'
    }
});

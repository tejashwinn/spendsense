import { CreateClientConfig } from './api/client/types.gen';

export const createClientConfig: CreateClientConfig = (config) => ({
    ...config,
    baseUrl: 'http://localhost:8080/api',
    // Hardcoding since user functionlities are not implemented yet
    headers: {
        'X-user-ID': '1'
    }
});

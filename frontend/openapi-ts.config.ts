import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
 input: '../backend/docs/swagger.json', // sign up at app.heyapi.dev
  output: 'lib/api',
});

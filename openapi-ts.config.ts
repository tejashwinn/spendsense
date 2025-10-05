import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
  input: '../backend/docs/swagger.json', // sign up at app.heyapi.dev
  output: 'lib/api',
  plugins: [

    {
      name: '@hey-api/client-next',
      runtimeConfigPath: '../../lib/client-api',

    }


  ],

});

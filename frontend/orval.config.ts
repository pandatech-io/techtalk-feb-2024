import { OutputOptions } from '@orval/core';
import { defineConfig } from 'orval';

const BASE_SERVICE_CONFIG_OUTPUT: OutputOptions = {
  client: 'react-query',
  mock: false,
  mode: 'tags-split',
  prettier: true,
};

export default defineConfig({
  doc: {
    input: { target: '../docs/doc.yaml' },
    output: {
      ...BASE_SERVICE_CONFIG_OUTPUT,
      override: {
        mutator: {
          path: './src/api/client.ts',
          name: 'request',
        },
      },
      schemas: 'src/api/generated/models/doc',
      target: 'src/api/generated/services/doc',
    },
  },
});
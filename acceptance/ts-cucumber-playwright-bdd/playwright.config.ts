import { defineConfig } from '@playwright/test';
import { defineBddConfig } from 'playwright-bdd';

const testDir = defineBddConfig({
  paths: ['features/*.feature'],
  require: ['src/steps/*.ts'],
});

export default defineConfig({
  testDir,
  timeout: 30 * 1000,
  expect: {
    timeout: 5000,
  },
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: 'html',
  use: {
    baseURL: 'http://localhost:8080',
    trace: 'on-first-retry',
    actionTimeout: 0,
    headless: true,
  },
  projects: [
    {
      name: 'domain',
      testMatch: '**/*domain*.spec.ts',
    },
    {
      name: 'http-inprocess',
      testMatch: '**/*http-inprocess*.spec.ts',
    },
    {
      name: 'http-executable',
      testMatch: '**/*http-executable*.spec.ts',
    },
    {
      name: 'http-docker',
      testMatch: '**/*http-docker*.spec.ts',
    },
    {
      name: 'ui',
      testMatch: '**/*ui*.spec.ts',
      use: {
        baseURL: 'http://localhost:3000',
      },
    },
  ],
  webServer: {
    command: 'cd ../../ && make server',
    port: 8080,
    reuseExistingServer: !process.env.CI,
  },
});
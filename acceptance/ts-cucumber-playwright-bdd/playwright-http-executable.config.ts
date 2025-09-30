import { defineConfig } from '@playwright/test';
import { defineBddConfig } from 'playwright-bdd';

const testDir = defineBddConfig({
  features: 'features/*.feature',
  steps: 'src/steps/http-steps.ts',
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
  workers: 1,
  reporter: 'line',
  use: {
    baseURL: 'http://localhost:8080',
    actionTimeout: 0,
    headless: true,
  },
  projects: [
    {
      name: 'http-executable-tests',
    },
  ],
  webServer: {
    command: 'cd ../../back-end && make build && ./bin/server',
    port: 8080,
    reuseExistingServer: !process.env.CI,
    timeout: 120 * 1000, // 2 minutes for build + startup
  },
});
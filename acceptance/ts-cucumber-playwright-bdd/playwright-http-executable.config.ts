import { defineConfig } from '@playwright/test';
import { defineBddConfig } from 'playwright-bdd';

const testDir = defineBddConfig({
  paths: ['features/*.feature'],
  require: ['src/steps/http-steps.ts'],
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
  // Use external server (requires manual startup)
});
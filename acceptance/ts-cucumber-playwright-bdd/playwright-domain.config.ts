import { defineConfig } from '@playwright/test';
import { defineBddConfig, cucumberReporter } from 'playwright-bdd';

const testDir = defineBddConfig({
  features: 'features/*.feature',
  steps: 'src/steps/domain-steps.ts',
});

export default defineConfig({
  testDir,
  timeout: 10 * 1000,
  expect: {
    timeout: 2000,
  },
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: 'line',
  use: {
    actionTimeout: 0,
    headless: true,
  },
  projects: [
    {
      name: 'domain-tests',
    },
  ],
});
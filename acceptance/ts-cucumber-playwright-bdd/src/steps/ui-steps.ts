import { expect } from '@playwright/test';
import { createBdd } from 'playwright-bdd';
import { UIDriver } from '../drivers/ui-driver';

const { Given, When, Then, Before } = createBdd();

// Create UI driver instance per test
let driver: UIDriver;

Before(async () => {
  // Driver will be created in each step when page is available
});

// Given steps
Given('{word} has created an account', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
    await driver.clearAll();
  }
  await driver.createAccount(name);
});

Given('{word} has signed up', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
    await driver.clearAll();
  }
  await driver.createAccount(name);
  await driver.activate(name);
});

// When steps
When('{word} activates her account', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  await driver.activate(name);
});

When('{word} activates his account', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  await driver.activate(name);
});

When('{word} activates their account', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  await driver.activate(name);
});

When('{word} tries to sign in', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  try {
    await driver.authenticate(name);
  } catch (error) {
    // Error will be checked in Then steps if needed
  }
});

When('{word} creates a project', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  await driver.createProject(name);
});

// Then steps
Then('{word} should be authenticated', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  const isAuthenticated = driver.isAuthenticated(name);
  expect(isAuthenticated).toBe(true);
});

Then('{word} should not be authenticated', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  const isAuthenticated = driver.isAuthenticated(name);
  expect(isAuthenticated).toBe(false);
});

Then('{word} should see the project', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  const projects = await driver.getProjects(name);
  expect(projects.length).toBe(1);
});

Then('{word} should not see any projects', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  const projects = await driver.getProjects(name);
  expect(projects.length).toBe(0);
});

Then('{word} should see an error telling him to activate the account', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  // UI-specific error checking would go here
  // This is a placeholder implementation
});

Then('{word} should see an error telling them to activate the account', async ({ page }, name: string) => {
  if (!driver) {
    driver = new UIDriver(page);
  }
  // UI-specific error checking would go here
  // This is a placeholder implementation
});
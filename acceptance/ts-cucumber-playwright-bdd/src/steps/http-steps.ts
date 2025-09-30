import { expect } from '@playwright/test';
import { createBdd } from 'playwright-bdd';
import { HttpDriver } from '../drivers/http-driver';

const { Given, When, Then, Before } = createBdd();

// Create HTTP driver instance per test
let driver: HttpDriver;
const lastErrors = new Map<string, Error | null>();

function setLastError(name: string, error: Error | null) {
  lastErrors.set(name, error);
}

function getLastError(name: string): Error | null {
  return lastErrors.get(name) || null;
}

Before(async () => {
  driver = new HttpDriver();
  lastErrors.clear();
  await driver.clearAll(); // Clear backend state before each test scenario
});

// Given steps
Given('{word} has created an account', async ({ }, name: string) => {
  await driver.createAccount(name);
});

Given('{word} has signed up', async ({ }, name: string) => {
  await driver.createAccount(name);
  await driver.activate(name);
  // Note: activate automatically authenticates the user in the backend
});

// When steps
When('{word} activates her account', async ({ }, name: string) => {
  await driver.activate(name);
  // Note: activate automatically authenticates the user in the backend
});

When('{word} activates his account', async ({ }, name: string) => {
  await driver.activate(name);
  // Note: activate automatically authenticates the user in the backend
});

When('{word} activates their account', async ({ }, name: string) => {
  await driver.activate(name);
  // Note: activate automatically authenticates the user in the backend
});

When('{word} tries to sign in', async ({ }, name: string) => {
  try {
    await driver.authenticate(name);
    setLastError(name, null);
  } catch (error) {
    setLastError(name, error as Error);
  }
});

When('{word} creates a project', async ({ }, name: string) => {
  await driver.createProject(name);
});

// Then steps
Then('{word} should be authenticated', async ({ }, name: string) => {
  const isAuthenticated = driver.isAuthenticated(name);
  expect(isAuthenticated).toBe(true);
});

Then('{word} should not be authenticated', async ({ }, name: string) => {
  const isAuthenticated = driver.isAuthenticated(name);
  expect(isAuthenticated).toBe(false);
});

Then('{word} should see the project', async ({ }, name: string) => {
  const projects = await driver.getProjects(name);
  expect(projects.length).toBe(1);
});

Then('{word} should not see any projects', async ({ }, name: string) => {
  const projects = await driver.getProjects(name);
  expect(projects.length).toBe(0);
});

Then('{word} should see an error telling him to activate the account', async ({ }, name: string) => {
  const lastError = getLastError(name);
  expect(lastError).not.toBeNull();
  expect(lastError!.message).toContain('activate');
});

Then('{word} should see an error telling them to activate the account', async ({ }, name: string) => {
  const lastError = getLastError(name);
  expect(lastError).not.toBeNull();
  expect(lastError!.message).toContain('activate');
});
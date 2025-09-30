import { Page } from '@playwright/test';
import { Account, Project } from '../client/models';
import { TestDriver } from './driver';

export class UIDriver implements TestDriver {
  private page: Page;
  private authenticatedAccounts = new Set<string>();

  constructor(page: Page) {
    this.page = page;
  }

  async createAccount(name: string): Promise<void> {
    await this.page.goto('/signup');

    // Wait for the form to load
    await this.page.waitForSelector('input[name="name"]', { timeout: 5000 });

    // Fill in the name field
    await this.page.fill('input[name="name"]', name);

    // Click the submit button
    await this.page.click('button[type="submit"]');

    // Wait for success message
    await this.page.waitForSelector('.success.message', { timeout: 5000 });
  }

  async clearAll(): Promise<void> {
    this.authenticatedAccounts.clear();
    // Use the React frontend's Clear page to clear backend data
    await this.page.goto('/admin/clear');
    await this.page.click('button', { timeout: 5000 });
    await this.page.waitForSelector('.success', { timeout: 5000 });
  }

  async getAccount(name: string): Promise<Account> {
    await this.page.goto(`/account/${name}`);

    try {
      // Wait for account info to load
      await this.page.waitForSelector('.account-info', { timeout: 5000 });

      // Extract account data from the page
      const activated = await this.page.$('.account-info') !== null;
      const authenticated = this.authenticatedAccounts.has(name);

      return {
        name,
        activated,
        authenticated
      } as Account;
    } catch {
      throw new Error(`Account not found: ${name}`);
    }
  }

  async authenticate(name: string): Promise<void> {
    await this.page.goto('/login');

    try {
      // Wait for login form
      await this.page.waitForSelector('input[name="name"]', { timeout: 5000 });

      // Fill and submit login form
      await this.page.fill('input[name="name"]', name);
      await this.page.click('button[type="submit"]');

      // Check if authentication was successful by looking for success redirect or error
      try {
        await this.page.waitForSelector('.success', { timeout: 2000 });
        this.authenticatedAccounts.add(name);
      } catch {
        // Check for error message
        const errorElement = await this.page.$('.error');
        if (errorElement) {
          const errorMessage = await errorElement.textContent();
          throw new Error(errorMessage || 'Authentication failed');
        }
        throw new Error('Authentication failed');
      }
    } catch (error) {
      // Let the error propagate for step definitions to catch
      throw error;
    }
  }

  isAuthenticated(name: string): boolean {
    return this.authenticatedAccounts.has(name);
  }

  async activate(name: string): Promise<void> {
    await this.page.goto(`/activate/${name}`);

    // Wait for activate button and click it
    await this.page.waitForSelector('button.activate', { timeout: 5000 });
    await this.page.click('button.activate');

    // Wait for success message
    await this.page.waitForSelector('.success', { timeout: 5000 });

    // Track as authenticated since activation automatically authenticates
    this.authenticatedAccounts.add(name);
  }

  async createProject(name: string): Promise<void> {
    await this.page.goto(`/account/${name}/projects`);

    // Wait for and click create project button
    await this.page.waitForSelector('button.create-project', { timeout: 5000 });
    await this.page.click('button.create-project');

    // Wait for project created confirmation
    await this.page.waitForSelector('.project-created', { timeout: 5000 });
  }

  async getProjects(name: string): Promise<Project[]> {
    await this.page.goto(`/account/${name}/projects`);

    try {
      // Wait for projects list to load
      await this.page.waitForSelector('.projects-list', { timeout: 5000 });

      // Extract project data from the page
      const projects = await this.page.$$eval('.project-item', (elements) =>
        elements.map((el, index) => ({
          id: `project-${index}`,
          name: el.textContent || '',
        }))
      );

      return projects as Project[];
    } catch {
      // No projects found
      return [];
    }
  }
}
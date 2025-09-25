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
    await this.page.fill('[data-testid="account-name"]', name);
    await this.page.click('[data-testid="signup-button"]');
    await this.page.waitForSelector('[data-testid="signup-success"]');
  }

  clearAll(): void {
    this.authenticatedAccounts.clear();
  }

  async getAccount(name: string): Promise<Account> {
    // Navigate to account details page or API endpoint
    // This is a simplified implementation
    return { name } as Account;
  }

  async authenticate(name: string): Promise<void> {
    await this.page.goto('/signin');
    await this.page.fill('[data-testid="account-name"]', name);
    await this.page.click('[data-testid="signin-button"]');

    try {
      // Check if authenticated successfully
      await this.page.waitForSelector('[data-testid="dashboard"]', { timeout: 2000 });
      this.authenticatedAccounts.add(name);
    } catch {
      // Authentication failed - might need activation
      const errorMessage = await this.page.textContent('[data-testid="error-message"]');
      if (errorMessage?.includes('activate')) {
        // Account needs activation, this is expected for some scenarios
        return;
      }
      throw new Error(`Authentication failed: ${errorMessage}`);
    }
  }

  isAuthenticated(name: string): boolean {
    return this.authenticatedAccounts.has(name);
  }

  async activate(name: string): Promise<void> {
    // Simulate email activation
    // In a real app, this would involve email verification
    await this.page.goto(`/activate?account=${name}`);
    await this.page.click('[data-testid="activate-button"]');
    await this.page.waitForSelector('[data-testid="activation-success"]');
  }

  async createProject(name: string): Promise<void> {
    await this.page.goto('/projects');
    await this.page.click('[data-testid="new-project-button"]');
    await this.page.fill('[data-testid="project-name"]', `${name}-project`);
    await this.page.click('[data-testid="create-project-button"]');
    await this.page.waitForSelector('[data-testid="project-created"]');
  }

  async getProjects(name: string): Promise<Project[]> {
    await this.page.goto('/projects');
    const projects = await this.page.$$eval('[data-testid="project-item"]', (elements) =>
      elements.map((el) => ({
        name: el.textContent || '',
        id: el.getAttribute('data-project-id') || '',
      }))
    );
    return projects as Project[];
  }
}
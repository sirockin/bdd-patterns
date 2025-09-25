import { Account, Project } from '../client/models';
import { TestDriver } from './driver';

export class DomainDriver implements TestDriver {
  private accounts = new Map<string, { account: Account; activated: boolean }>();
  private authenticatedAccounts = new Set<string>();
  private projects = new Map<string, Project[]>();

  async createAccount(name: string): Promise<void> {
    if (this.accounts.has(name)) {
      throw new Error(`Account ${name} already exists`);
    }
    this.accounts.set(name, {
      account: { name },
      activated: false,
    });
  }

  clearAll(): void {
    this.accounts.clear();
    this.authenticatedAccounts.clear();
    this.projects.clear();
  }

  async getAccount(name: string): Promise<Account> {
    const accountData = this.accounts.get(name);
    if (!accountData) {
      throw new Error(`Account ${name} not found`);
    }
    return accountData.account;
  }

  async authenticate(name: string): Promise<void> {
    const accountData = this.accounts.get(name);
    if (!accountData) {
      throw new Error(`Account ${name} not found`);
    }

    if (!accountData.activated) {
      throw new Error(`Account ${name} needs to be activated`);
    }

    this.authenticatedAccounts.add(name);
  }

  isAuthenticated(name: string): boolean {
    return this.authenticatedAccounts.has(name);
  }

  async activate(name: string): Promise<void> {
    const accountData = this.accounts.get(name);
    if (!accountData) {
      throw new Error(`Account ${name} not found`);
    }
    accountData.activated = true;
  }

  async createProject(name: string): Promise<void> {
    if (!this.authenticatedAccounts.has(name)) {
      throw new Error(`Account ${name} is not authenticated`);
    }

    const existingProjects = this.projects.get(name) || [];
    const newProject: Project = {
      name: `${name}-project`,
      id: `project-${name}-${Date.now()}`,
      owner: name,
    };

    this.projects.set(name, [...existingProjects, newProject]);
  }

  async getProjects(name: string): Promise<Project[]> {
    return this.projects.get(name) || [];
  }
}
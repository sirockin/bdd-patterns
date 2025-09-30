import { DefaultApi, Account, Project, Configuration } from '../client';
import { TestDriver } from './driver';

export class HttpDriver implements TestDriver {
  private api: DefaultApi;
  private authenticatedAccounts = new Set<string>();

  constructor(baseURL: string = 'http://localhost:8080') {
    const configuration = new Configuration({
      basePath: baseURL,
    });
    this.api = new DefaultApi(configuration);
  }

  async createAccount(name: string): Promise<void> {
    try {
      await this.api.createAccount({ name });
    } catch (error: any) {
      if (error.response?.status !== 201) {
        throw new Error(`Failed to create account: ${error.message}`);
      }
    }
  }

  async clearAll(): Promise<void> {
    this.authenticatedAccounts.clear();
    try {
      await this.api.clearAll();
    } catch (error: any) {
      throw new Error(`Failed to clear backend state: ${error.message}`);
    }
  }

  async getAccount(name: string): Promise<Account> {
    try {
      const response = await this.api.getAccount(name);
      return response.data;
    } catch (error: any) {
      throw new Error(`Failed to get account: ${error.message}`);
    }
  }

  async authenticate(name: string): Promise<void> {
    try {
      const response = await this.api.authenticateAccount(name);
      if (response.status === 200) {
        this.authenticatedAccounts.add(name);
      }
    } catch (error: any) {
      if (error.response?.status === 400) {
        // Account needs activation - preserve the error message from backend
        const backendMessage = error.response?.data?.error || error.response?.data?.message || error.message || '';
        throw new Error(backendMessage);
      }
      if (error.response?.status === 404) {
        throw new Error(`Account not found: ${name}`);
      }
      throw new Error(`Failed to authenticate: ${error.message}`);
    }
  }

  isAuthenticated(name: string): boolean {
    return this.authenticatedAccounts.has(name);
  }

  async activate(name: string): Promise<void> {
    try {
      await this.api.activateAccount(name);
      // Backend automatically authenticates user upon activation
      this.authenticatedAccounts.add(name);
    } catch (error: any) {
      throw new Error(`Failed to activate account: ${error.message}`);
    }
  }

  async createProject(name: string): Promise<void> {
    try {
      await this.api.createProject(name);
    } catch (error: any) {
      throw new Error(`Failed to create project: ${error.message}`);
    }
  }

  async getProjects(name: string): Promise<Project[]> {
    try {
      const response = await this.api.getProjects(name);
      return response.data || [];
    } catch (error: any) {
      throw new Error(`Failed to get projects: ${error.message}`);
    }
  }
}
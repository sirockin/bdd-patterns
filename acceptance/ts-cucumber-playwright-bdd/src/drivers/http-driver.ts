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

  clearAll(): void {
    this.authenticatedAccounts.clear();
    // Note: In a real implementation, this would clear the backend state
    // For now, we just clear the local authentication state
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
      if (error.response?.status === 403) {
        // Account needs activation
        return;
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
    } catch (error: any) {
      throw new Error(`Failed to activate account: ${error.message}`);
    }
  }

  async createProject(name: string): Promise<void> {
    try {
      await this.api.createProject(name, { name: `${name}-project` });
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
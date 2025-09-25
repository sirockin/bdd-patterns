import { Account, Project } from '../client/models';

export interface TestDriver {
  createAccount(name: string): Promise<void>;
  clearAll(): void;
  getAccount(name: string): Promise<Account>;
  authenticate(name: string): Promise<void>;
  isAuthenticated(name: string): boolean;
  activate(name: string): Promise<void>;
  createProject(name: string): Promise<void>;
  getProjects(name: string): Promise<Project[]>;
}
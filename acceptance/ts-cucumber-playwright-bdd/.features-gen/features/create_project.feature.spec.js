/** Generated from: features/create_project.feature */
import { test } from "playwright-bdd";

test.describe("Create project", () => {

  test("Create one project", async ({ Given, When, Then }) => {
    await Given("Sue has signed up");
    await When("Sue creates a project");
    await Then("Sue should see the project");
  });

  test("Try to see someone else's project", async ({ Given, And, When, Then }) => {
    await Given("Sue has signed up");
    await And("Bob has signed up");
    await When("Sue creates a project");
    await Then("Bob should not see any projects");
  });

});

// == technical section ==

test.use({
  $test: ({}, use) => use(test),
  $uri: ({}, use) => use("features/create_project.feature"),
  $bddFileMeta: ({}, use) => use(bddFileMeta),
});

const bddFileMeta = {
  "Create one project": {"pickleLocation":"5:3"},
  "Try to see someone else's project": {"pickleLocation":"10:3"},
};
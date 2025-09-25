/** Generated from: features/sign_up.feature */
import { test } from "playwright-bdd";

test.describe("Sign up", () => {

  test("Successful sign-up", async ({ Given, When, Then }) => {
    await Given("Tanya has created an account");
    await When("Tanya activates her account");
    await Then("Tanya should be authenticated");
  });

  test("Try to sign in without activating account", async ({ Given, When, Then, And }) => {
    await Given("Bob has created an account");
    await When("Bob tries to sign in");
    await Then("Bob should not be authenticated");
    await And("Bob should see an error telling him to activate the account");
  });

});

// == technical section ==

test.use({
  $test: ({}, use) => use(test),
  $uri: ({}, use) => use("features/sign_up.feature"),
  $bddFileMeta: ({}, use) => use(bddFileMeta),
});

const bddFileMeta = {
  "Successful sign-up": {"pickleLocation":"6:3"},
  "Try to sign in without activating account": {"pickleLocation":"11:3"},
};
## About Screenplay

The [Screenplay Pattern](https://serenity-js.org/handbook/design/screenplay-pattern/) is an evolution of the Page Object Pattern for UI test automation. It uses Actors with Abilities to perform Actions which can be grouped into Tasks. Actors can also ask Questions to verify outcomes.

While it has its origins in UI test automation, the pattern is applicable to acceptance testing in general, including API testing as demonstrated here.

It is useful where there are multiple user roles (Actors) with different capabilities (Abilities) performing different operations (Tasks) to achieve different goals (Questions).

It carries with it an extra layer of complexity and is probably best suited to projects in organizations where BDD is already well understood and widely adopted.

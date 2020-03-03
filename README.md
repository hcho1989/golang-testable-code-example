# Good Practices to Write Unit-Testable Golang Code

## TLDR:

To produce readable and testable code, remember: Dependency Injection!


## What this is about:

1. showing examples to transform untestable code into testable code
2. mindset that we should have to write testable code

## What this is not about:

1. picking on codes written untestably
2. setting rules of thumb on coding


## Golang:
- strong typed
- no runtime change for a struct, cannot stub function call (unlike javascript you can always change stub a function call)
- hard to mock on the fly (still possible by some monkey-patching approach, but it will make your codes messy and hairy)



To show untestable code, I have written an app that sends a slack message, both in testable and untestable code.


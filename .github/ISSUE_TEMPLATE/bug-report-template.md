name: Bug Report 🐞
description: Report a bug and help Search.gov improve.
title: "Search.gov - Bug: [YOUR TITLE]"
labels: ["bug"]
projects: ["GSA-TTS/60"]
body:
  - type: textarea
    id: problem
    attributes:
      label: Describe the bug
      description: Add a clear and concise description of the bug. Let us know if it impacts major or minor functionality and if you have workaround.
    validations:
      required: true
  - type: textarea
    id: reproduce
    attributes:
      label: Steps to reproduce the bug
      description: Describe how to reproduce this issue.
      placeholder: |
        1. Go to '...'
        2. Click on '....'
        3. Scroll down to '....'
        4. See error...
    validations:
      required: true
  - type: textarea
    id: expectation
    attributes:
      label: Expected Behavior
      description: Add a clear and concise description of what you expected to happen.
    validations:
      required: true
  - type: textarea
    id: screenshots
    attributes:
      label: Screenshots
      description: Add screenshots to help provide context, if applicable.
    validations:
      required: false
  - type: textarea
    id: system
    attributes:
      label: System setup
      description: Provide your system details. Be sure to include your device, operating system, and browser (with version).
      placeholder: |
        - Device:
        - Operating system:
        - Browser and version:
    validations:
      required: false
  - type: textarea
    id: context
    attributes:
      label: Additional context
      description: Add any other context about the problem.
    validations:
      required: false
  - type: checkboxes
    id: terms
    attributes:
      label: Code of Conduct
      description: Please confirm the following
      options:
        - label:
            I agree to abide by the [Digital.gov Community Guidelines](https://digital.gov/communities/community-guidelines/) and the [TTS Code of Conduct](https://handbook.tts.gsa.gov/code-of-conduct/). Respect your peers, use plain language, and be patient. 
          required: true
        - label:
            I checked the [current
              issues](https://github.com/GSA-TTS/FAC/issues?q=is%3Aissue+is%3Aopen+label%3A%22Type%3A+Bug%22) for
              duplicate bug reports.
          required: true